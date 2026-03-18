package endpoint

import "time"

type Endpoint struct {
	EndpointID int64      `gorm:"column:endpoint_id;primaryKey"`
	Value      string     `gorm:"column:value"`
	Method     string     `gorm:"column:method"`
	Type       string     `gorm:"column:type"`
	Bypass     int        `gorm:"column:bypass"`

	Description *string   `gorm:"column:description"`
	CreatedBy   *string   `gorm:"column:created_by"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	Pagination  *string   `gorm:"column:pagination"`
}

func (Endpoint) TableName() string {
	return "endpoint"
}