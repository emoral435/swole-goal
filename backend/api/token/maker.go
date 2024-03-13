package token

import "time"

// Manages the interfaces for managing tokens
type Maker interface {
	// creates a new token for a specific email and duration
	CreateToken(email string, duration time.Duration) (string, error)

	// checks if the input token is valid or not
	VerifyToken(token string) (*Payload, error)
}
