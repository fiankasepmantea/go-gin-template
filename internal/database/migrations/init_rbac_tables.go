package migrations

import (
	"fmt"
	"gorm.io/gorm"
)

var InitRBACMigration = Migration{
	Name: "init_rbac_tables",
	Up: func(db *gorm.DB) {

		// ---------------- USERS ----------------
		resUser := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(200) PRIMARY KEY,
			user_unique SERIAL4 UNIQUE,
			group_id VARCHAR(200) NULL,
			domain_id INT NOT NULL,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(200) NULL,
			username VARCHAR(200) NOT NULL,
			password VARCHAR(200) NOT NULL,
			avatar VARCHAR(200) NULL,
			status INT2 NULL,
			token VARCHAR(200) NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
			login_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
			access_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
			is_admin INT2 DEFAULT 0 NULL,
			no_wa VARCHAR(50) NULL,
			join_date TIMESTAMP NULL,
			gender VARCHAR(30) NULL,
			nik VARCHAR(20) NULL,
			device TEXT NULL
		)
		`)
		fmt.Println("users table error:", resUser.Error)

		// ---------------- GROUPS ----------------
		resGroup := db.Exec(`
		CREATE TABLE IF NOT EXISTS groups (
			group_id VARCHAR(200) PRIMARY KEY,
			domain_id INT NOT NULL,
			group_name VARCHAR(200) NULL,
			status INT4 NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
			avatar VARCHAR(255) NULL
		)
		`)
		fmt.Println("groups table error:", resGroup.Error)

		// ---------------- ENDPOINT ----------------
		resEndpoint := db.Exec(`
		CREATE TABLE IF NOT EXISTS endpoint (
			endpoint_id VARCHAR(200) PRIMARY KEY,
			value VARCHAR(200) NOT NULL,
			description VARCHAR(200) NULL,
			created_by VARCHAR(200) NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NULL,
			method VARCHAR(20) NOT NULL,
			type VARCHAR(20) NOT NULL,
			bypass INT NOT NULL,
			pagination VARCHAR NULL
		)
		`)
		fmt.Println("endpoint table error:", resEndpoint.Error)

		// ---------------- GROUP_ENDPOINT ----------------
		resGroupEndpoint := db.Exec(`
		CREATE TABLE IF NOT EXISTS group_endpoint (
			id VARCHAR(200) PRIMARY KEY,
			group_id VARCHAR(200),
			endpoint_id VARCHAR(200)
		)
		`)
		fmt.Println("group_endpoint table error:", resGroupEndpoint.Error)
	},
}