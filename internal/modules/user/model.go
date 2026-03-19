package user

import (
	"time"

	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
)

type User struct {
	UserID    string     `gorm:"column:user_id;primaryKey"`
	UserUnique int       `gorm:"column:user_unique;autoIncrement;unique"`
	GroupID   *string    `gorm:"column:group_id"`
	DomainID  int        `gorm:"column:domain_id"`
	Name      string     `gorm:"column:name"`
	Email     *string    `gorm:"column:email"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Avatar    *string    `gorm:"column:avatar"`
	Status    *int16     `gorm:"column:status"`
	Token     *string    `gorm:"column:token"`
	CreatedAt time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	LoginDate time.Time  `gorm:"column:login_date;default:CURRENT_TIMESTAMP"`
	AccessAt  time.Time  `gorm:"column:access_at;default:CURRENT_TIMESTAMP"`
	IsAdmin   *int16     `gorm:"column:is_admin;default:0"`
	NoWA      *string    `gorm:"column:no_wa"`
	JoinDate  *time.Time `gorm:"column:join_date"`
	Gender    *string    `gorm:"column:gender"`
	NIK       *string    `gorm:"column:nik"`
	Device    *string    `gorm:"column:device"`
}

// NewUserID returns a new UUID string for user_id
func NewUserID() string {
	return idgen.NewUserID()
}

func (User) TableName() string {
	return "users"
}