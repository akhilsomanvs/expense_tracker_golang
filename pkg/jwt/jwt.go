package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`

	jwtlib.RegisteredClaims
}
type TokenPayload struct {
	UserID string
	Email  string
}

type Manager struct {
	secret []byte
	issuer string
	ttl    time.Duration
}

func New(secret, issuer string, ttl time.Duration) *Manager {
	return &Manager{
		secret: []byte(secret),
		issuer: issuer,
		ttl:    ttl,
	}
}

func (s *Manager) GenerateToken(tokenPayload *TokenPayload) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: tokenPayload.UserID,
		Email:  tokenPayload.Email,
		RegisteredClaims: jwtlib.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   tokenPayload.UserID,
			IssuedAt:  jwtlib.NewNumericDate(now),
			ExpiresAt: jwtlib.NewNumericDate(now.Add(s.ttl)),
		},
	}

	token := jwtlib.NewWithClaims(
		jwtlib.SigningMethodES256,
		claims,
	)

	return token.SignedString(s.secret)
}

func (s *Manager) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwtlib.Token) (any, error) {
			return s.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwtlib.ErrTokenInvalidClaims
	}

	return claims, nil

}

func (s *Manager) TTL() time.Duration {
	return s.ttl
}
