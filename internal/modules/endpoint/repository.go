package endpoint

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindAll(out *[]Endpoint) error {
	return r.DB.Find(out).Error
}

func (r *Repository) FindByPath(path, method string, out *Endpoint) error {
	return r.DB.Where("value = ? AND method = ?", path, method).First(out).Error
}