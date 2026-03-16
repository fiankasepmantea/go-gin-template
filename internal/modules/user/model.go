package user

import "time"

type User struct {
	UserID     string     `gorm:"column:user_id;primaryKey"`
	UserUnique int        `gorm:"column:user_unique"`
	GroupID    *int       `gorm:"column:group_id"`
	DomainID   *int       `gorm:"column:domain_id"`

	Name     string  `gorm:"column:name"`
	Email    *string `gorm:"column:email"`
	Username string  `gorm:"column:username"`
	Password string  `gorm:"column:password"`

	Avatar *string `gorm:"column:avatar"`
	Status *int16  `gorm:"column:status"`
	Token  *string `gorm:"column:token"`

	CreatedAt *time.Time `gorm:"column:created_at"`
	LoginDate *time.Time `gorm:"column:login_date"`
	AccessAt  *time.Time `gorm:"column:access_at"`

	IsAdmin *int16  `gorm:"column:is_admin"`
	NoWA    *string `gorm:"column:no_wa"`
}

func (User) TableName() string {
	return "user"
}