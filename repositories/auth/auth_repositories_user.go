package auth

import (
	"ecobuy/entities"
	"ecobuy/repositories/models"
	"errors"

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
	LoginUser(user entities.User) (entities.User, error)
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

func (ar *AuthRepository) LoginUser(user entities.User) (entities.User, error) {
	userDB := models.FromEntitiesUser(user)
	err := ar.db.First(&userDB, "email = ?", userDB.Email)
	if err.Error != nil {
		return entities.User{}, errors.New("email tidak ditemukan")
	}
	return userDB.ToEntities(), nil
}
