package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("token is invalid")

// this contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	numericDate := &jwt.NumericDate{Time: p.ExpiredAt}
	return numericDate, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	numericDate := &jwt.NumericDate{Time: p.IssuedAt}
	return numericDate, nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	numericDate := &jwt.NumericDate{Time: p.IssuedAt}
	return numericDate, nil
}

func (p *Payload) GetIssuer() (string, error) {
	return "server", nil
}

func (p *Payload) GetSubject() (string, error) {
	return p.ID.String(), nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}

func NewPayload(email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return payload, nil
}
