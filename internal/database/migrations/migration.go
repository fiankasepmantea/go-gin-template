package migrations

import "gorm.io/gorm"

type Migration struct {
	Name string
	Up   func(db *gorm.DB)
	Down func(db *gorm.DB)
}