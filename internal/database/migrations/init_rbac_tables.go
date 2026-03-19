package migrations

import (
	"fmt"

	"gorm.io/gorm"
)

var InitRBACMigration = Migration{
	Name: "init_rbac_tables",
	Up: func(db *gorm.DB) {

		// USERS
		resultUser := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(50) PRIMARY KEY,
			name VARCHAR(100),
			username VARCHAR(100) UNIQUE,
			password TEXT,
			group_id INT,
			is_admin SMALLINT
		)
		`)
		fmt.Println("users table error:", resultUser.Error)
		// GROUPS
		resultGroup := db.Exec(`
		CREATE TABLE IF NOT EXISTS groups (
			group_id SERIAL PRIMARY KEY,
			group_name VARCHAR(100)
		)
		`)
		fmt.Println("users table error:", resultGroup.Error)
		// ENDPOINT
		resultEndpoint := db.Exec(`
		CREATE TABLE IF NOT EXISTS endpoint (
			endpoint_id SERIAL PRIMARY KEY,
			value VARCHAR(200),
			method VARCHAR(10),
			type VARCHAR(20),
			bypass INT
		)
		`)
		fmt.Println("users table error:", resultEndpoint.Error)
		// GROUP ENDPOINT
		resultGroupEndpoint := db.Exec(`
		CREATE TABLE IF NOT EXISTS group_endpoint (
			id SERIAL PRIMARY KEY,
			group_id INT,
			endpoint_id INT
		)
		`)
		fmt.Println("users table error:", resultGroupEndpoint.Error)
	},
}
