package service

import (
	"time"

	"aidanwoods.dev/go-paseto"
)

type TokenIFace interface {
	NewToken(string) (string, error)
	ValidateToken(string) (bool, error)
	GetUidFromToken(string) (string, error)
	// InvalidateToken(string) error
}

type Token struct {
	key *paseto.V4AsymmetricSecretKey
}

func NewTokenServicer() *Token {
	key := paseto.NewV4AsymmetricSecretKey()

	return &Token{
		&key,
	}
}

func (t *Token) NewToken(uid string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	err := token.Set("odsId", uid)
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
