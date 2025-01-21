package auth

import (
	"ecobuy/entities"
	"ecobuy/repositories/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(user entities.User) (entities.User, error)
}

type AuthService struct {
	AuthRepository auth.AuthRepositoryInterface
}

func NewAuthService(ar auth.AuthRepositoryInterface) *AuthService {
	return &AuthService{
		AuthRepository: ar,
	}
}

func (as *AuthService) RegisterUser(user entities.User) (entities.User, error) {
	// Validasi email
	if user.Email == "" {
		return entities.User{}, errors.New("email kosong")
	}

	// Validasi password
	if user.Password == "" {
		return entities.User{}, errors.New("password kosong")
	}

	// Periksa apakah email sudah ada
	exists, err := as.AuthRepository.CheckEmailExists(user.Email)
	if err != nil {
		return entities.User{}, err
	}

	if exists {
		return entities.User{}, errors.New("email sudah ada")
	}

	// hash password
	hash, _ := HashPassword(user.Password)
	user.Password = hash

	user, err = as.AuthRepository.RegisterUser(user)
	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
