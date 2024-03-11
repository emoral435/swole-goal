package token

import (
	"errors"
	"time"

	"aidanwoods.dev/go-paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	paseto       *paseto.V2SymmetricKey
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, errors.New("invalid key size for pasetoMaker")
	}

	newSymKey := paseto.NewV2SymmetricKey()
	maker := &PasetoMaker{
		paseto:       &newSymKey,
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// creates a new token for a specific email and duration
func (p *PasetoMaker) CreateToken(email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(email, duration)

	if err != nil {
		return "", err
	}

	return paseto.NewToken().V2Encrypt(), nil
}

// checks if the input token is valid or not
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error)
