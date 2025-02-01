package main

import (
	"ecobuy/config"
	"ecobuy/controllers/auth"
	"ecobuy/controllers/product"
	"ecobuy/middlewares"
	AuthRepositories "ecobuy/repositories/auth"
	ProductRepositories "ecobuy/repositories/product"
	"ecobuy/routes"
	AuthService "ecobuy/services/auth"
	ProductService "ecobuy/services/product"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	LoadEnv()
	// Menghubungkan ke database
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Jalankan Migrasi Database
	config.RunMigrations(db)

	// Membuat instance Echo
	e := echo.New()

	jwtUser := middlewares.JwtUser{}
	authRepo := AuthRepositories.NewAuthRepository(db)
	authService := AuthService.NewAuthService(authRepo, jwtUser)
	authController := auth.NewAuthController(authService)

	productRepo := ProductRepositories.NewProductRepository(db)
	productService := ProductService.NewProductService(productRepo)
	productController := product.NewProductController(productService)

	routeController := routes.RouteController{
		AuthController:    *authController,
		ProductController: *productController,
	}
	routeController.RegisterRoutes(e)

	e.Start(":8000")
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
