package config

import (
	"ecobuy/repositories/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	// AutoMigrate tabel
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.ImpactData{})
}
