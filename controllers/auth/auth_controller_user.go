package auth

import (
	"ecobuy/controllers/auth/request"
	"ecobuy/controllers/auth/response"
	"ecobuy/services/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService auth.AuthServiceInterface
}

func NewAuthController(as auth.AuthServiceInterface) *AuthController {
	return &AuthController{
		AuthService: as,
	}
}

func (ac *AuthController) RegisterController(c echo.Context) error {
	userRegister := request.RegisterRequest{}
	c.Bind(&userRegister)

	// Panggil service untuk mendaftarkan pengguna
	user, err := ac.AuthService.RegisterUser(userRegister.ToEntities())
	if err != nil {
		// Periksa error untuk memberikan pesan spesifik
		if err.Error() == "email already exists" {
			return c.JSON(http.StatusConflict, map[string]interface{}{
				"message": "Email already exists",
			})
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Registration Successful",
		"user":    response.RegisterFromEntities(user),
	})
}

func (ac *AuthController) LoginController(c echo.Context) error {
	userLogin := request.LoginRequest{}
	c.Bind(&userLogin)
	user, err := ac.AuthService.LoginUser(userLogin.ToEntities())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user":    response.LoginFromEntities(user),
	})
}
