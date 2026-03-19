package base

import (
	"time"
	// "github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
	// "gorm.io/gorm"
)

type BaseModel struct {
	// ID        string    `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
}

// func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	if b.ID == "" {
// 		b.ID = idgen.NewID()
// 	}
// 	return
// }