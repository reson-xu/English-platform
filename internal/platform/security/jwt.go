package security

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/reson-xu/english-platform/internal/platform/constants"
)

var (
	ErrInvalidToken = errors.New("invalid access token")
	ErrExpiredToken = errors.New("access token expired")
)

type TokenService struct {
	secret []byte
	ttl    time.Duration
}

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"uid"`
	Email  string `json:"eml"`
	Role   string `json:"rol"`
}

func NewTokenService(secret string, ttl time.Duration) *TokenService {
	return &TokenService{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

func (s *TokenService) GenerateAccessToken(userID string, email string, role string) (string, error) {
	now := time.Now().UTC()
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.ttl)),
		},
		UserID: userID,
		Email:  email,
		Role:   role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.secret)
	if err != nil {
		return "", fmt.Errorf("sign jwt: %w", err)
	}

	return signed, nil
}

func (s *TokenService) ParseAccessToken(token string) (Claims, error) {
	token = strings.TrimSpace(token)
	if strings.HasPrefix(token, constants.BearerPrefix) {
		token = strings.TrimPrefix(token, constants.BearerPrefix)
	}

	var claims Claims
	parsed, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w: unexpected signing method %v", ErrInvalidToken, t.Header["alg"])
			}
			return s.secret, nil
		},
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return Claims{}, ErrExpiredToken
		}
		return Claims{}, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	if !parsed.Valid {
		return Claims{}, ErrInvalidToken
	}

	return claims, nil
}
