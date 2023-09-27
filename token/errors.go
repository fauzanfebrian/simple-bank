package token

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

var (
	ErrInvalidSecretKeySize    = fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	ErrInvalidSymmetricKeySize = fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	ErrInvalidToken            = fmt.Errorf("invalid token")
	ErrTokenExpired            = fmt.Errorf("token expired")
)
