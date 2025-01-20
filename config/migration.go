package config

import "gorm.io/gorm"

func RunMigrations(db *gorm.DB) {
	// AutoMigrate tabel
	db.AutoMigrate()
}
