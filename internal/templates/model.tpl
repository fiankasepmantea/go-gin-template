package {{.Package}}

import "time"

type {{.StructName}} struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ({{.StructName}}) TableName() string {
	return "{{.TableName}}"
}