package service

import (
	"context"
	"log"
	"time"

	"aidanwoods.dev/go-paseto"
	errors "github.com/pkg/errors"
)

type TokenIFace interface {
	NewToken(string) (string, error)
	ValidateToken(string) (bool, error)
	GetUidFromToken(string) (string, error)
	GenerateAccessToken(string) (string, error)
	GenerateRefreshToken(string) (string, error)
	RefreshAccessToken(string) (string, string, error)
	InvalidateToken(string) error
}

const (
	AccessTokenLife  = 5 * time.Minute
	RefreshTokenLife = 7 * 24 * time.Hour
)

type Token struct {
	key *paseto.V4AsymmetricSecretKey
	Db InMemoryDbIFace
}

func NewTokenServicer(db InMemoryDbIFace) *Token {
	key := paseto.NewV4AsymmetricSecretKey()

	return &Token{
		&key,
		db,
	}
}

func (t *Token) NewToken(uid string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	err := token.Set("odsId", uid)
	log.Println("uid set: " + uid)
	if err != nil {
		return "", nil
	}

	return token.V4Sign(*t.key, []byte("public")), nil
}

func (t *Token) parseToken(jwt string) (*paseto.Token, error) {
	parser := paseto.NewParser()

	return parser.ParseV4Public(t.key.Public(), jwt, []byte("public"))

}

func (t *Token) ValidateToken(token string) (bool, error) {
	_, err := t.parseToken(token)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (t *Token) GetUidFromToken(jwt string) (string, error) {
	token, err := t.parseToken(jwt)
	if err != nil {
		return "", err
	}

	var id string
	err = token.Get("odsId", &id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (t *Token) GenerateAccessToken(uid string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(AccessTokenLife))
	err := token.Set("odsId", uid)
	if err != nil {
		return "", errors.Wrap(err, "failed to set UID in access token")
	}

	accessToken := token.V4Sign(*t.key, []byte("public"))

	err = t.Db.Set(context.Background(), "access_token:"+uid, accessToken, AccessTokenLife)
	if err != nil {
		return "", errors.Wrap(err, "failed to store access token in cache")
	}

	return accessToken, nil
}

func (t *Token) GenerateRefreshToken(uid string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(RefreshTokenLife))
	err := token.Set("odsId", uid)
	if err != nil {
		return "", errors.Wrap(err, "failed to set UID in refresh token")
	}

	refreshToken := token.V4Sign(*t.key, []byte("public"))

	err = t.Db.Set(context.Background(), "refresh_token:"+uid, refreshToken, RefreshTokenLife)
	if err != nil {
		return "", errors.Wrap(err, "failed to store refresh token in cache")
	}

	return refreshToken, nil
}

func (t *Token) RefreshAccessToken(refreshToken string) (string, string, error) {
	uid, err := t.GetUidFromToken(refreshToken)
	if err != nil {
		return "", "", errors.Wrap(err, "invalid refresh token")
	}

	if err := t.Db.Del(context.Background(), "access_token:"+uid); err != nil {
		log.Printf("Warning: Failed to remove old access token from cache for user %s: %v\n", uid, err)
	}

	if err := t.Db.Del(context.Background(), "refresh_token:"+uid); err != nil {
		log.Printf("Warning: Failed to remove old refresh token from cache for user %s: %v\n", uid, err)
	}


	newAccessToken, err := t.GenerateAccessToken(uid)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate new access token")
	}

	newRefreshToken, err := t.GenerateRefreshToken(uid)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to generate new refresh token")
	}
	return newAccessToken, newRefreshToken, nil
}


func (t *Token) InvalidateToken(tokenKey string) error {
	if err := t.Db.Del(context.Background(), tokenKey); err != nil {
		return errors.Wrap(err, "failed to remove token from cache")
	}

	return nil
}