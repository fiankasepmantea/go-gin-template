package groupendpoint

import (
	"github.com/fiankasepman/go-gin-template/internal/base"
	"gorm.io/gorm"
)

type Repository struct {
	base.BaseRepository[GroupEndpoint]
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BaseRepository: base.BaseRepository[GroupEndpoint]{DB: db},
	}
}

func (r *Repository) DeleteByGroupAndEndpoint(groupID, endpointID string) error {
	return r.DB.
		Where("group_id = ? AND endpoint_id = ?", groupID, endpointID).
		Delete(&GroupEndpoint{}).Error
}

func (r *Repository) FindByGroup(groupID string, out *[]GroupEndpoint) error {
	return r.DB.
		Where("group_id = ?", groupID).
		Find(out).Error
}