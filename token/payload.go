package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"jti"`
	Username  string    `json:"username"`
	ExpiredAt time.Time `json:"exp"`
	IssuedAt  time.Time `json:"iat"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        uuid,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// GetExpirationTime implements the Claims interface.
func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	expiredAt := &jwt.NumericDate{
		Time: payload.ExpiredAt,
	}
	return expiredAt, nil
}

// GetNotBefore implements the Claims interface.
func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{}, nil
}

// GetIssuedAt implements the Claims interface.
func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	issuedAt := &jwt.NumericDate{
		Time: payload.IssuedAt,
	}
	return issuedAt, nil
}

// GetAudience implements the Claims interface.
func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

// GetIssuer implements the Claims interface.
func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

// GetSubject implements the Claims interface.
func (payload *Payload) GetSubject() (string, error) {
	return "", nil
}
