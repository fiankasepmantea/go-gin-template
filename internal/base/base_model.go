package base

import (
	"time"
	// "github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
	// "gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
}

// func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	if b.ID == "" {
// 		b.ID = idgen.NewID()
// 	}
// 	return
// }