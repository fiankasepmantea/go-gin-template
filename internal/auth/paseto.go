package auth

import (
	"errors"
	"time"

	"github.com/o1egl/paseto"
	"github.com/fiankasepman/go-gin-template/configs"
)

var pasetoV2 = paseto.NewV2()

type TokenPayload struct {
	UserID   string    `json:"user_id"`
	TokenID  string    `json:"token_id"`
	DomainID int       `json:"domain_id"`
	Exp      time.Time `json:"exp"`
}

func GenerateToken(userID, tokenID string, domainID int) (string, error) {

	payload := TokenPayload{
		UserID:   userID,
		TokenID:  tokenID,
		DomainID: domainID,
		Exp:      time.Now().Add(configs.AccessTokenDuration),
	}

	return pasetoV2.Encrypt([]byte(configs.PasetoKey), payload, nil)
}

func ValidateToken(token string) (*TokenPayload, error) {

	var payload TokenPayload

	err := pasetoV2.Decrypt(token, []byte(configs.PasetoKey), &payload, nil)
	if err != nil {
		return nil, err
	}

	if time.Now().After(payload.Exp) {
		return nil, errors.New("token expired")
	}

	return &payload, nil
}