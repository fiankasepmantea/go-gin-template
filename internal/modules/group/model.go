package group

import (

	// "github.com/fiankasepman/go-gin-template/internal/base"
	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
)

type Group struct {
	// base.BaseModel
	GroupID   string     `gorm:"column:group_id;primaryKey"`
	DomainID  int        `gorm:"column:domain_id"`
	GroupName string     `gorm:"column:group_name"`
	Status    *int       `gorm:"column:status"`
	Avatar    *string    `gorm:"column:avatar"`
}

func NewGroupID() string {
	return idgen.NewGroupID()
}

func (Group) TableName() string {
	return "groups"
}
