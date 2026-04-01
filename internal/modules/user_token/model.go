package usertoken

import "time"

type UserToken struct {
	ID           string     `gorm:"column:id;primaryKey"`
	UserID       string     `gorm:"column:user_id"`
	RefreshToken string     `gorm:"column:refresh_token"`
	Device       *string    `gorm:"column:device"`
	UserAgent    *string    `gorm:"column:user_agent"`
	IPAddress    *string    `gorm:"column:ip_address"`
	ExpiresAt    time.Time  `gorm:"column:expires_at"`
	CreatedAt    time.Time  `gorm:"column:created_at"`
}

func (UserToken) TableName() string {
	return "user_tokens"
}