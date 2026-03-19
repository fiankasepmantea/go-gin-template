package database

import (
	user "github.com/fiankasepman/go-gin-template/internal/modules/user"
	group "github.com/fiankasepman/go-gin-template/internal/modules/group"
	endpoint "github.com/fiankasepman/go-gin-template/internal/modules/endpoint"
	groupendpoint "github.com/fiankasepman/go-gin-template/internal/modules/groupendpoint"
	"github.com/fiankasepman/go-gin-template/internal/auth"
)

func SeedAll() {
	SeedGroup()
	SeedEndpoint()
	SeedAdmin()
	SeedGroupEndpoint()
}

// -------------------- GROUP --------------------
func SeedGroup() {
	var count int64
	DB.Table("groups").Where("group_name = ?", "Super Admin").Count(&count)
	if count > 0 {
		return
	}

	DB.Table("groups").Create(map[string]interface{}{
		"group_id":   group.NewGroupID(),
		"group_name": "Super Admin",
		"domain_id":  1,
	})
}

// -------------------- ENDPOINT --------------------
func SeedEndpoint() {
	endpoints := []map[string]interface{}{
		{"endpoint_id": endpoint.NewEndpointID(), "value": "/users", "method": "GET", "type": "backend", "bypass": 0},
		{"endpoint_id": endpoint.NewEndpointID(), "value": "/me", "method": "GET", "type": "backend", "bypass": 0},
		{"endpoint_id": endpoint.NewEndpointID(), "value": "/login", "method": "POST", "type": "backend", "bypass": 1},
	}

	for _, e := range endpoints {
		var count int64
		DB.Table("endpoint").Where("value = ? AND method = ?", e["value"], e["method"]).Count(&count)
		if count == 0 {
			DB.Table("endpoint").Create(&e)
		}
	}
}

// -------------------- ADMIN USER --------------------
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

	var groupID string
	DB.Table("groups").
		Select("group_id").
		Where("group_name = ?", "Super Admin").
		Scan(&groupID)

	admin := user.User{
		UserID:   user.NewUserID(),
		Name:     "Admin",
		Username: "admin",
		Password: pass,
		GroupID:  &groupID,
		IsAdmin:  &isAdmin,
		DomainID: 1,
	}

	DB.Table("users").Create(&admin)
}

// -------------------- GROUP_ENDPOINT --------------------
func SeedGroupEndpoint() {
	var count int64
	DB.Table("group_endpoint").Count(&count)
	if count > 0 {
		return
	}

	var groupID string
	DB.Table("groups").
		Select("group_id").
		Where("group_name = ?", "Super Admin").
		Scan(&groupID)

	var endpoints []struct {
		EndpointID string `gorm:"column:endpoint_id"`
	}

	DB.Table("endpoint").Select("endpoint_id").Find(&endpoints)

	for _, e := range endpoints {
		DB.Table("group_endpoint").Create(map[string]interface{}{
			"id":          groupendpoint.NewGroupEndpointID(),
			"group_id":    groupID,
			"endpoint_id": e.EndpointID,
		})
	}
}