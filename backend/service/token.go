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
	InvalidateTokensForUid(string) error
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

func (t *Token) accessTokenKey(uid string) string {
    return "access_token:" + uid;
}

func (t *Token) refreshTokenKey(uid string) string {
    return "refresh_token:" + uid;
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

	accessTokenKey := t.accessTokenKey(uid)
	err = t.Db.Set(context.Background(), accessTokenKey, accessToken, AccessTokenLife)
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

	refreshTokenKey := t.refreshTokenKey(uid)
	err = t.Db.Set(context.Background(), refreshTokenKey, refreshToken, RefreshTokenLife)
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
  
	err = t.InvalidateTokensForUid(uid);
	if err != nil {
		return "", "", errors.Wrap(err, "failed to invalidate tokens for uid")
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

func (t *Token) InvalidateTokensForUid(uid string) error {
	if err := t.Db.Del(context.Background(), t.accessTokenKey(uid)); err != nil {
		return errors.Wrap(err, "failed to remove access token from cache")
	}

	if err := t.Db.Del(context.Background(), t.refreshTokenKey(uid)); err != nil {
		return errors.Wrap(err, "failed to remove refresh token from cache")
	}

	return nil
}


func (t *Token) InvalidateToken(tokenKey string) error {
	if err := t.Db.Del(context.Background(), tokenKey); err != nil {
		return errors.Wrap(err, "failed to remove token from cache")
	}

	return nil
}