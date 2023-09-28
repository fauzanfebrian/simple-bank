package token

import "fmt"

var (
	ErrInvalidSecretKeySize    = fmt.Errorf("invalid key size: must be at least %d characters", KeySize)
	ErrInvalidSymmetricKeySize = fmt.Errorf("invalid key size: must be exactly %d characters", KeySize)
	ErrInvalidToken            = fmt.Errorf("invalid token")
	ErrTokenExpired            = fmt.Errorf("token expired")
)
