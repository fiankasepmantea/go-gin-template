package user

type User struct {
	UserID   string  `gorm:"column:user_id;primaryKey"`
	Name     string  `gorm:"column:name"`
	Username string  `gorm:"column:username"`
	Password string  `gorm:"column:password"`
	GroupID  *int    `gorm:"column:group_id"`
	IsAdmin  *int16  `gorm:"column:is_admin"`
}

func (User) TableName() string {
	return "users"
}