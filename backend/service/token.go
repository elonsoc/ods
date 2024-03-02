package service

import (
	"context"
	"log"
	"time"

	"aidanwoods.dev/go-paseto"
)

type TokenIFace interface {
	NewToken(string) (string, error)
	RefreshAccessToken(string) (string, error)
	ValidateToken(string) (bool, error)
	GetUidFromToken(string) (string, error)
	GenerateRefreshToken(string) (string, error)
	// InvalidateToken(string) error
}

type Token struct {
	key *paseto.V4AsymmetricSecretKey
	Redis RedisIFace
}

func NewTokenServicer(redis RedisIFace) *Token {
	key := paseto.NewV4AsymmetricSecretKey()

	return &Token{
		&key,
		redis,
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

func (t *Token) GenerateRefreshToken(uid string) (string, error) {
    token := paseto.NewToken()
    
    token.SetIssuedAt(time.Now())
    token.SetExpiration(time.Now().Add(30 * 24 * time.Hour))
    
    token.SetSubject(uid)
    
    refreshToken := token.V4Sign(*t.key, nil)
    
    ctx := context.Background()
    expiration := 30 * 24 * time.Hour
    err := t.Redis.Set(ctx, refreshToken, uid, expiration)
    if err != nil {
        log.Println("Error storing refresh token in Redis: ", err)
        return "", err
    }
    
    return refreshToken, nil
}


func (t *Token) RefreshAccessToken(refreshToken string) (string, error) {
    ctx := context.Background()
    
    odsId, err := t.Redis.Get(ctx, refreshToken)
    if err == nil {
        log.Fatal(err);
	}

    if odsId == "" {
        log.Fatal("odsId not found for refreshToken: " + refreshToken);
    }

    token, err := t.NewToken(odsId)
    if err != nil {
		log.Fatal(err);
    }

    return token, nil
}

