package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/Minhajxdd/Ephemr/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type TokenManager interface {
	Generate(payload Claims) (string, error)
	Validate(tokenStr string) (*Claims, error)
}

type jwtManager struct {
	secret string
	ttl    time.Duration
}

func NewJWTManager(secret string, ttl time.Duration) TokenManager {
	return &jwtManager{secret: secret, ttl: ttl}
}

func (j *jwtManager) Generate(payload Claims) (string, error) {
	payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(config.Cfg.JwtTTl))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(j.secret))
}

func (j *jwtManager) Validate(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
