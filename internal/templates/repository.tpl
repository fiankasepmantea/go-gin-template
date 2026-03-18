package {{.Package}}

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[{{.StructName}}]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[{{.StructName}}]{DB: db},
	}
}