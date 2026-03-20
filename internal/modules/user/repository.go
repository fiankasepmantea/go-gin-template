package user

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[User]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[User]{DB: db},
	}
}

func (r *Repository) FindByUsername(username string, out *User) error {
	return r.DB.Where("username = ?", username).First(out).Error
}