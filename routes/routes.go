package routes

import (
	"ecobuy/controllers/auth"
	"ecobuy/controllers/product"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	AuthController    auth.AuthController
	ProductController product.ProductController
}

func (rc RouteController) RegisterRoutes(e *echo.Echo) {
	// endpoint user
	e.POST("/register", rc.AuthController.RegisterController)
	e.POST("/login", rc.AuthController.LoginController)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	// Endpoint Produk (Bisa filter kategori & support pagination)
	eUserProduct := eJwt.Group("/product")
	eUserProduct.GET("", rc.ProductController.GetProductsController)
}
