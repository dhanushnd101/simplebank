package token

import (
	"time"
	"fmt"
	"errors"

	// "github.com/google/uuid"
	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32
// JWTMaker is a JSON web token maker 
type JWTMaker struct{
	secretKey string
}

// NewJWTMaker creates a new JWTMaker 
func NewJWTMaker(secretKey string) (Maker, error){
	if(len(secretKey) < minSecretKeySize){
		return nil, fmt.Errorf("Invalid key size: The scretKey is less than %d charachers", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil	
}

	// CreateTOken creates a new token for a specific username and duration
func (maker *JWTMaker)	CreateToken(username string, duration time.Duration)(string, error){
	payload,err := NewPayload(duration, username)
	if(err != nil){
		return "",err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

	// VerifyToken takes the token and verifyies if its a valied token or not
func (maker *JWTMaker)	VerifyToken(token string) (*Payload, error){
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
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