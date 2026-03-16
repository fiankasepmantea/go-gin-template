package group

import "time"

type Group struct {
	GroupID   int        `gorm:"column:group_id;primaryKey"`
	DomainID  *int       `gorm:"column:domain_id"`
	GroupName *string    `gorm:"column:group_name"`
	Status    *int       `gorm:"column:status"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	Avatar    *string    `gorm:"column:avatar"`
}

func (Group) TableName() string {
	return "group"
}