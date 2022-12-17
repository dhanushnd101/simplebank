package token

import (
	"time"
)

// Maker is a interface for managing tokesn
type Maker interface{
	// CreateTOken creates a new token for a specific username and duration
	CreateToken(username string, duration time.Duration)(string, error)

	// VerifyToken takes the token and verifyies if its a valied token or not
	VerifyToken(token string) (*Payload, error)
}