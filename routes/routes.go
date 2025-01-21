package routes

import (
	"ecobuy/controllers/auth"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController auth.AuthController
}

func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	// endpoint user
	e.POST("/register", rc.AuthController.RegisterController)
}
