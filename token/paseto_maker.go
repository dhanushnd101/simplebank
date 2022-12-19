package token

import (
	"time"
	"fmt"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)
type PaseetoMaker struct{
	paseto *paseto.V2
	symmetricKey []byte
	
}

// NewPasetoMaker creates a new PasetoMaker
func NewPasetoMaker(symmetricKey string)(Maker, error){
	if len(symmetricKey) != chacha20poly1305.KeySize{
		return nil, fmt.Errorf("invalid key size: musthave exactly %d characters", chacha20poly1305.KeySize)
	}
	maker := &PaseetoMaker{
		paseto: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

	// CreateTOken creates a new token for a specific username and duration
func (maker *PaseetoMaker)	CreateToken(username string, duration time.Duration)(string, *Payload, error){
	payload, err := NewPayload(duration, username)
	if(err != nil){
		return "", payload, err
	}
	
	token, err :=  maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

	// VerifyToken takes the token and verifyies if its a valied token or not
func (maker *PaseetoMaker)	VerifyToken(token string) (*Payload, error){
	payload := &Payload{}

	err:= maker.paseto.Decrypt(token,maker.symmetricKey, payload,nil)
	if(err!= nil){
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if(err!= nil){
		return nil, err
	}
	return payload, nil
}