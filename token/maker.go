package token

import "time"

type Maker interface {
	// CreateToken creates a token for the given username that expires after the given duration.
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken verifies the given token and returns the username associated with it.
	VerifyToken(token string) (*Payload, error)
}
