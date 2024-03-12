package pkg

type TokenIFace interface {
	NewToken(string) (string, error)
	ValidateToken(string) (bool, error)
	GetUidFromToken(string) (string, error)
	GenerateAccessToken(string) (string, error)
	GenerateRefreshToken(string) (string, error)
	RefreshAccessToken(string) (string, string, error)
	InvalidateToken(string) error
}

type Token struct{}

func NewTokenService(authURL string) TokenIFace {

	return &Token{}
}
func (t *Token) NewToken(s string) (string, error) {
	return "", nil
}

func (t *Token) ValidateToken(s string) (bool, error) {
	return true, nil
}

func (t *Token) GetUidFromToken(string) (string, error) {
	return "", nil
}

func (t *Token) GenerateAccessToken(string) (string, error) {
	return "", nil
}

func (t *Token) GenerateRefreshToken(string) (string, error) {
	return "", nil
}

func (t *Token) RefreshAccessToken(string) (string, string, error) {
	return "", "", nil
}

func (t *Token) InvalidateToken(string) error {
	return nil
}
