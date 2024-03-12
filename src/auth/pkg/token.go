package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TokenIFace interface {
	NewToken(string) (string, error)
	ValidateToken(string) (bool, error)
	GetUidFromToken(string) (string, error)
	RefreshAccessToken(string) (string, string, error)
	InvalidateToken(string) error
}

type Token struct {
	url string
}

type TokenHttp struct {
	T string `json:"Token"`
}

func NewTokenService(authURL string) TokenIFace {

	return &Token{url: authURL}
}
func (t *Token) NewToken(s string) (string, error) {
	c, _ := json.Marshal(TokenHttp{T: s})
	res, err := http.DefaultClient.Post(fmt.Sprintf("http://%s/token/", t.url), "encoding/json", bytes.NewBuffer(c))
	if err != nil {
		return "", err
	}

	a := TokenHttp{}

	json.NewDecoder(res.Body).Decode(&a)
	return a.T, nil
}

func (t *Token) ValidateToken(s string) (bool, error) {
	c, _ := json.Marshal(TokenHttp{T: s})
	res, err := http.DefaultClient.Post(fmt.Sprintf("http://%s/token/validate", t.url), "encoding/json", bytes.NewBuffer(c))
	if err != nil {
		return false, err
	}

	a := struct {
		Verdict bool `json:"verdict"`
	}{}

	json.NewDecoder(res.Body).Decode(&a)
	return a.Verdict, nil
}

func (t *Token) GetUidFromToken(s string) (string, error) {
	c, _ := json.Marshal(TokenHttp{T: s})
	res, err := http.DefaultClient.Post(fmt.Sprintf("http://%s/token/uid", t.url), "encoding/json", bytes.NewBuffer(c))
	if err != nil {
		return "", err
	}

	a := struct {
		Uid string `json:"uid"`
	}{}

	json.NewDecoder(res.Body).Decode(&a)
	return a.Uid, nil
}

func (t *Token) RefreshAccessToken(s string) (string, string, error) {
	c, _ := json.Marshal(TokenHttp{T: s})

	res, err := http.DefaultClient.Post(fmt.Sprintf("http://%s/token/refresh", t.url), "encoding/json", bytes.NewBuffer(c))
	if err != nil {
		return "", "", err
	}

	a := struct {
		Access  string
		Refresh string
	}{}

	json.NewDecoder(res.Body).Decode(&a)
	return a.Access, a.Refresh, nil
}

func (t *Token) InvalidateToken(s string) error {
	c, _ := json.Marshal(TokenHttp{T: s})

	_, err := http.DefaultClient.Post(fmt.Sprintf("http://%s/token/invalidate", t.url), "encoding/json", bytes.NewBuffer(c))
	if err != nil {
		// we could log the res, if necessary.
		return err
	}
	return nil
}
