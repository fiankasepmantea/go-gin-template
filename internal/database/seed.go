package database

import (
	"github.com/fiankasepman/go-gin-template/internal/auth"
	endpoint "github.com/fiankasepman/go-gin-template/internal/modules/endpoint"
	group "github.com/fiankasepman/go-gin-template/internal/modules/group"
	groupendpoint "github.com/fiankasepman/go-gin-template/internal/modules/groupendpoint"
	user "github.com/fiankasepman/go-gin-template/internal/modules/user"
)

const (
	GroupNameSuperAdmin = "Super Admin"
	WhereGroupByDomain  = "group_name = ? AND domain_id = ?"
	DefaultDomainID     = 1
	DefaultUsername     = "admin"
	DefaultPassword     = "admin123"
)

func SeedAll() {
	SeedGroup()
	SeedEndpoint()
	SeedAdmin()
	SeedGroupEndpoint()
}

// ================= GROUP =================
func SeedGroup() {
	var count int64

	DB.Table("groups").
		Where(WhereGroupByDomain, GroupNameSuperAdmin, DefaultDomainID).
		Count(&count)

	if count > 0 {
		return
	}

	DB.Table("groups").Create(map[string]interface{}{
		"group_id":   group.NewGroupID(),
		"group_name": GroupNameSuperAdmin,
		"domain_id":  DefaultDomainID,
	})
}

// ================= ENDPOINT =================
func SeedEndpoint() {

	endpoints := []map[string]interface{}{
		{"value": "/login", "method": "POST", "bypass": 1},
		{"value": "/refresh", "method": "POST", "bypass": 1},

		{"value": "/users", "method": "GET", "bypass": 0},
		{"value": "/users", "method": "POST", "bypass": 0},
		{"value": "/users/:id", "method": "PUT", "bypass": 0},
		{"value": "/users/:id", "method": "DELETE", "bypass": 0},
		{"value": "/me", "method": "GET", "bypass": 0},
		{"value": "/devices", "method": "GET", "bypass": 0},
	}

	for _, e := range endpoints {

		var count int64
		DB.Table("endpoint").
			Where("value = ? AND method = ?", e["value"], e["method"]).
			Count(&count)

		if count == 0 {

			bypass := 0
			if v, ok := e["bypass"]; ok {
				bypass = v.(int)
			}

			DB.Table("endpoint").Create(map[string]interface{}{
				"endpoint_id": endpoint.NewEndpointID(),
				"value":       e["value"],
				"method":      e["method"],
				"type":        "API",
				"bypass":      bypass,
			})
		}
	}
}

// ================= ADMIN =================
func SeedAdmin() {

	var count int64
	DB.Table("users").
		Where("username = ?", DefaultUsername).
		Count(&count)

	if count > 0 {
		return
	}

	pass, _ := auth.HashPassword(DefaultPassword)
	isAdmin := int16(1)

	var groupID string
	DB.Table("groups").
		Select("group_id").
		Where(WhereGroupByDomain, GroupNameSuperAdmin, DefaultDomainID).
		Scan(&groupID)

	DB.Table("users").Create(map[string]interface{}{
		"user_id":   user.NewUserID(),
		"name":      "Admin",
		"username":  DefaultUsername,
		"password":  pass,
		"group_id":  groupID,
		"is_admin":  isAdmin,
		"domain_id": DefaultDomainID,
	})
}

// ================= GROUP ENDPOINT =================
func SeedGroupEndpoint() {

	var groupID string
	DB.Table("groups").
		Select("group_id").
		Where(WhereGroupByDomain, GroupNameSuperAdmin, DefaultDomainID).
		Scan(&groupID)

	var endpoints []struct {
		EndpointID string
	}

	DB.Table("endpoint").Select("endpoint_id").Find(&endpoints)

	for _, e := range endpoints {

		var count int64

		DB.Table("group_endpoint").
			Where("group_id = ? AND endpoint_id = ?", groupID, e.EndpointID).
			Count(&count)

		if count == 0 {
			DB.Table("group_endpoint").Create(map[string]interface{}{
				"id":          groupendpoint.NewGroupEndpointID(),
				"group_id":    groupID,
				"endpoint_id": e.EndpointID,
			})
		}
	}
}