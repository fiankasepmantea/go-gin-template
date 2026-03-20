package group

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[Group]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[Group]{DB: db},
	}
}