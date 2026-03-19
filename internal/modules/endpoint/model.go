package endpoint

import (
	"time"

	"github.com/fiankasepman/go-gin-template/internal/pkg/idgen"
)

type Endpoint struct {
	EndpointID string     `gorm:"column:endpoint_id;primaryKey"`
	Value      string     `gorm:"column:value"`
	Description *string   `gorm:"column:description"`
	CreatedBy  *string    `gorm:"column:created_by"`
	CreatedAt  time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	Method     string     `gorm:"column:method"`
	Type       string     `gorm:"column:type"`
	Bypass     int        `gorm:"column:bypass"`
	Pagination *string    `gorm:"column:pagination"`
}

func NewEndpointID() string {
	return idgen.NewEndpointID()
}

func (Endpoint) TableName() string {
	return "endpoint"
}