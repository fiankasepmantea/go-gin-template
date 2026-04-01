package usertoken

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[UserToken]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[UserToken]{DB: db},
	}
}

func (r *Repository) FindByToken(token string, out *UserToken) error {
	return r.DB.Where("refresh_token = ?", token).First(out).Error
}

func (r *Repository) DeleteByToken(token string) error {
	return r.DB.Where("refresh_token = ?", token).Delete(&UserToken{}).Error
}

func (r *Repository) DeleteByUser(userID string) error {
	return r.DB.Where("user_id = ?", userID).Delete(&UserToken{}).Error
}

func (r *Repository) DeleteByID(id string) error {
	return r.DB.Delete(&UserToken{}, "id = ?", id).Error
}

func (r *Repository) FindByUser(userID string, out *[]UserToken) error {
	return r.DB.Where("user_id = ?", userID).Find(out).Error
}
func (r *Repository) DeleteExpired() error {
	return r.DB.
		Where("expires_at < NOW()").
		Delete(&UserToken{}).Error
}