package database

import (
	"fmt"

	"github.com/fiankasepman/go-gin-template/internal/database/migrations"
)

func RunMigrations() {

	DB.Exec(`
	CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	for _, m := range migrations.Migrations {

		var count int64

		DB.Table("migrations").
			Where("name = ?", m.Name).
			Count(&count)

		if count > 0 {
			continue
		}

		fmt.Println("running:", m.Name)

		m.Up(DB)

		DB.Exec("INSERT INTO migrations (name) VALUES (?)", m.Name)
	}
}