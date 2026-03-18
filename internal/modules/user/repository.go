package user

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) FindAll(users *[]User) error {
	return r.DB.Find(users).Error
}

func (r *Repository) FindByUsername(username string, user *User) error {
	return r.DB.Where("username = ?", username).First(user).Error
}

func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}