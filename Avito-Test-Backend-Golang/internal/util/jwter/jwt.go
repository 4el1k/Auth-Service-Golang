package jwter

import (
	"Avito-Test-Backend-Golang/internal/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"time"
)

type JWTer interface {
	EncodeToken(string) (string, error)
	DecodeToken(string) (string, error)
}

type JwtManager struct {
	ttl    time.Duration
	secret string
	issuer string
}

func NewJwtManager(ttl time.Duration, secret string, issuer string) *JwtManager {
	return &JwtManager{ttl: ttl, secret: secret, issuer: issuer}
}

type AuthClaims struct {
	Principal uuid.UUID
	IsAdmin   bool
	TagId     int
	jwt.RegisteredClaims
}

func (m *JwtManager) EncodeToken(user *model.User) (string, error) {
	exp := time.Now()
	authClaims := &AuthClaims{
		Principal: user.Id,
		IsAdmin:   user.IsAdmin,
		TagId:     user.TagId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			ExpiresAt: jwt.NewNumericDate(exp.Add(m.ttl)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims)
	t, err := token.SignedString([]byte(m.secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (m *JwtManager) DecodeToken(jwtToken string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &AuthClaims{}, m.getKeyFunc())
	if err != nil {
		return &AuthClaims{}, err
	}
	claims, ok := token.Claims.(*AuthClaims)
	if !ok {
		return &AuthClaims{}, errors.New("invalid token")
	}
	return claims, nil
}

func (m *JwtManager) getKeyFunc() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.secret), nil
	}
}
