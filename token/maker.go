package token

import (
	"time"

	"golang.org/x/crypto/chacha20poly1305"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

const KeySize int = chacha20poly1305.KeySize
