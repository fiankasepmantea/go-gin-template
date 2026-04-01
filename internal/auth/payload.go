package auth

type Payload struct {
	UserID  string `json:"user_id"`
	TokenID string `json:"token_id"`
	DomainID int `json:"domain_id"`
}