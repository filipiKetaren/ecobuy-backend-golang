package auth

import (
	"ecobuy/entities"
	"ecobuy/middlewares"
	"ecobuy/repositories/auth"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceInterface interface {
	RegisterUser(user entities.User) (entities.User, error)
	LoginUser(user entities.User) (entities.User, error)
}

type AuthService struct {
	AuthRepository auth.AuthRepositoryInterface
	jwtInterface   middlewares.JwtInterface
}

func NewAuthService(ar auth.AuthRepositoryInterface, jt middlewares.JwtInterface) *AuthService {
	return &AuthService{
		AuthRepository: ar,
		jwtInterface:   jt,
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

func (as *AuthService) LoginUser(user entities.User) (entities.User, error) {
	if user.Email == "" {
		return entities.User{}, errors.New("email kosong")
	} else if user.Password == "" {
		return entities.User{}, errors.New("password kosong")
	}

	oldPassword := user.Password

	// Cari pengguna berdasarkan email
	user, err := as.AuthRepository.LoginUser(user)
	if err != nil {
		return entities.User{}, err
	}

	match := CheckPasswordHash(oldPassword, user.Password)
	if !match {
		return entities.User{}, errors.New("password salah")
	}

	// Generate token JWT
	token, err := as.jwtInterface.GenerateJWT(user.ID)
	if err != nil {
		return entities.User{}, err
	}

	user.Token = token
	return user, nil
}
