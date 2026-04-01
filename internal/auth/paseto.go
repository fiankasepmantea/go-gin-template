package auth

import (
	"time"

	"github.com/o1egl/paseto"
)

var pasetoV2 = paseto.NewV2()

var symmetricKey = []byte("12345678901234567890123456789012")

type TokenPayload struct {
	UserID  string    `json:"user_id"`
	TokenID string    `json:"token_id"`
	Exp     time.Time `json:"exp"`
}

func GenerateToken(userID, tokenID string) (string, error) {

	payload := TokenPayload{
		UserID:  userID,
		TokenID: tokenID,
		Exp:     time.Now().Add(24 * time.Hour),
	}

	return pasetoV2.Encrypt(symmetricKey, payload, nil)
}

func ValidateToken(token string) (*TokenPayload, error) {

	var payload TokenPayload

	err := pasetoV2.Decrypt(token, symmetricKey, &payload, nil)
	if err != nil {
		return nil, err
	}

	if time.Now().After(payload.Exp) {
		return nil, err
	}

	return &payload, nil
}