package database

import (
	"github.com/fiankasepman/go-gin-template/internal/auth"
	"github.com/fiankasepman/go-gin-template/internal/modules/user"
)

func SeedSuperAdmin() {

	var count int64

	DB.Table("user").
		Where("username = ?", "admin").
		Count(&count)

	if count > 0 {
		return
	}

	pass, _ := auth.HashPassword("admin123")

	isAdmin := int16(1)

	admin := user.User{
		UserID:  "admin-1",
		Name:    "Super Admin",
		Username:"admin",
		Password: pass,
		IsAdmin: &isAdmin,
	}

	DB.Create(&admin)
}