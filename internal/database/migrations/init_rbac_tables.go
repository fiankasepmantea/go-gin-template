package migrations

import (
	"log"

	"gorm.io/gorm"
)

var InitRBACMigration = Migration{
	Name: "init_rbac_tables",
	Up: func(db *gorm.DB) {

		mustExec(db, `
		CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(200) PRIMARY KEY,
			user_unique SERIAL UNIQUE,
			group_id VARCHAR(200),
			domain_id INT NOT NULL,

			name VARCHAR(255) NOT NULL,
			email VARCHAR(200),
			username VARCHAR(200) NOT NULL,
			password VARCHAR(200) NOT NULL,

			avatar VARCHAR(200),
			status INT2,
			is_admin INT2 DEFAULT 0,

			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			login_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			access_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)

		mustExec(db, `
		CREATE TABLE IF NOT EXISTS groups (
			group_id VARCHAR(200) PRIMARY KEY,
			domain_id INT NOT NULL,
			group_name VARCHAR(200),
			status INT4,
			avatar VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)

		mustExec(db, `
		CREATE TABLE IF NOT EXISTS endpoint (
			endpoint_id VARCHAR(200) PRIMARY KEY,
			value VARCHAR(200) NOT NULL,
			method VARCHAR(20) NOT NULL,
			type VARCHAR(20) NOT NULL,
			bypass INT NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(value, method)
		)`)

		// ✅ FIXED (missing closing bracket)
		mustExec(db, `
		CREATE TABLE IF NOT EXISTS group_endpoint (
			id VARCHAR(200) PRIMARY KEY,
			group_id VARCHAR(200) NOT NULL,
			endpoint_id VARCHAR(200) NOT NULL,
			UNIQUE(group_id, endpoint_id)
		)`)

		mustExec(db, `
		CREATE TABLE IF NOT EXISTS user_tokens (
			id VARCHAR(200) PRIMARY KEY,
			user_id VARCHAR(200) NOT NULL,
			refresh_token VARCHAR(500) NOT NULL,
			device VARCHAR(200),
			user_agent TEXT,
			ip_address VARCHAR(100),
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
	},
}

func mustExec(db *gorm.DB, query string) {
	if err := db.Exec(query).Error; err != nil {
		log.Fatal("migration error:", err)
	}
}