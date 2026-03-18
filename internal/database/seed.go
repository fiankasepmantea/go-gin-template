package database

import (
	user "github.com/fiankasepman/go-gin-template/internal/modules/user"
	"github.com/fiankasepman/go-gin-template/internal/auth"
)

func SeedAll() {
	SeedGroup()
	SeedEndpoint()
	SeedAdmin()
	SeedGroupEndpoint()
}

func SeedGroup() {

	var count int64

	DB.Table("groups").
		Where("group_name = ?", "Super Admin").
		Count(&count)

	if count > 0 {
		return
	}

	DB.Table("groups").Create(map[string]interface{}{
		"group_name": "Super Admin",
	})
}

func SeedEndpoint() {

	endpoints := []map[string]interface{}{
		{"value": "/users", "method": "GET", "type": "backend", "bypass": 0},
		{"value": "/me", "method": "GET", "type": "backend", "bypass": 0},
		{"value": "/login", "method": "POST", "type": "backend", "bypass": 1},
	}

	for _, e := range endpoints {

		var count int64

		DB.Table("endpoint").
			Where("value = ? AND method = ?", e["value"], e["method"]).
			Count(&count)

		if count == 0 {
			DB.Table("endpoint").Create(&e)
		}
	}
}

func SeedAdmin() {

	var count int64

	DB.Table("users").
		Where("username = ?", "admin").
		Count(&count)

	if count > 0 {
		return
	}

	pass, _ := auth.HashPassword("admin123")
	isAdmin := int16(1)

	var groupID int
	DB.Table("groups").
		Select("group_id").
		Where("group_name = ?", "Super Admin").
		Scan(&groupID)

	admin := user.User{
		UserID:  "admin-1",
		Name:    "Admin",
		Username: "admin",
		Password: pass,
		GroupID: &groupID,
		IsAdmin: &isAdmin,
	}

	DB.Table("users").Create(&admin)
}

func SeedGroupEndpoint() {

	var count int64
	DB.Table("group_endpoint").Count(&count)
	if count > 0 {
		return
	}

	var groupID int
	DB.Table("groups").
		Select("group_id").
		Where("group_name = ?", "Super Admin").
		Scan(&groupID)

	var endpoints []struct {
		EndpointID int `gorm:"column:endpoint_id"`
	}

	DB.Table("endpoint").Select("endpoint_id").Find(&endpoints)

	for _, e := range endpoints {
		DB.Table("group_endpoint").Create(map[string]interface{}{
			"group_id":   groupID,
			"endpoint_id": e.EndpointID,
		})
	}
}