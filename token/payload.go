package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	*jwt.RegisteredClaims
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:       uuid,
		Username: username,
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(duration),
			},
		},
	}
	return payload, nil
}
