package migrations

import "gorm.io/gorm"

func Up{{.FuncName}}(db *gorm.DB) {
	// TODO: create table
}

func Down{{.FuncName}}(db *gorm.DB) {
	// TODO: drop table
}