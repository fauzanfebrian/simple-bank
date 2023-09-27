package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, ErrInvalidSecretKeySize
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	var keyFunc jwt.Keyfunc = func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	payload := &Payload{}
	jwtToken, err := jwt.ParseWithClaims(token, payload, keyFunc)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			// if token expired we still return payload in case the payload used
			return payload, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
