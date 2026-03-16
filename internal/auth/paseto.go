package auth

import (
	"time"

	"github.com/o1egl/paseto"
)

var pasetoMaker = paseto.NewV2()

var symmetricKey = []byte("12345678901234567890123456789012")

type Payload struct {
	UserID string
	Exp    time.Time
}

func GenerateToken(userID string) (string, error) {

	payload := Payload{
		UserID: userID,
		Exp:    time.Now().Add(time.Hour * 24),
	}

	return pasetoMaker.Encrypt(symmetricKey, payload, nil)
}

func VerifyToken(token string) (*Payload, error) {

	var payload Payload

	err := pasetoMaker.Decrypt(token, symmetricKey, &payload, nil)

	return &payload, err
}