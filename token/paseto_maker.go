package token

import (
	"encoding/json"
	"time"

	"aidanwoods.dev/go-paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, ErrInvalidSymmetricKeySize
	}

	v4SymmetricKey, err := paseto.V4SymmetricKeyFromBytes([]byte(symmetricKey))
	if err != nil {
		return nil, err
	}

	pasetoMaker := &PasetoMaker{
		symmetricKey: v4SymmetricKey,
	}
	return pasetoMaker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	claims, err := payloadToPasetoClaim(payload)
	if err != nil {
		return "", err
	}

	token, err := paseto.MakeToken(claims, nil)
	if err != nil {
		return "", err
	}

	return token.V4Encrypt(maker.symmetricKey, nil), nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	parser := paseto.MakeParser([]paseto.Rule{})

	tokenPaseto, err := parser.ParseV4Local(maker.symmetricKey, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	claimsJson := tokenPaseto.ClaimsJSON()
	payload, err := pasetoClaimToPayload(claimsJson)
	if err != nil {
		return nil, err
	}

	if time.Now().After(payload.ExpiredAt) {
		return payload, ErrTokenExpired
	}

	return payload, nil
}

func payloadToPasetoClaim(payload *Payload) (map[string]any, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// Convert JSON data to a map[string]interface{}
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func pasetoClaimToPayload(claimsJson []byte) (*Payload, error) {
	payload := &Payload{}

	if err := json.Unmarshal(claimsJson, payload); err != nil {
		return nil, err
	}

	return payload, nil
}
