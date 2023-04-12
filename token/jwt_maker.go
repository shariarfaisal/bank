package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeyLength = 32

type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker with the given secret key.
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLength {
		return nil, fmt.Errorf("secret key must be at least %d characters long", minSecretKeyLength)
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a token for the given username that expires after the given duration.
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload := NewPayload(username, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

// VerifyToken verifies the given token and returns the username associated with it.
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		varr, ok := err.(*jwt.ValidationError)
		if ok && varr.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
