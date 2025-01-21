package auth

import (
	"ecobuy/entities"
	"ecobuy/repositories/models"

	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepositoryInterface interface {
	RegisterUser(entities.User) (entities.User, error)
	CheckEmailExists(email string) (bool, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func (ar *AuthRepository) RegisterUser(user entities.User) (entities.User, error) {
	userDB := models.FromEntitiesUser(user)
	err := ar.db.Create(&userDB)
	if err.Error != nil {
		return entities.User{}, err.Error
	}
	return userDB.ToEntities(), nil
}

func (ar *AuthRepository) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := ar.db.Model(&models.User{}).Where("email = ?", email).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
