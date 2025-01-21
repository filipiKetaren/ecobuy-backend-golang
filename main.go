package main

import (
	"ecobuy/config"
	"ecobuy/controllers/auth"
	AuthRepositories "ecobuy/repositories/auth"
	"ecobuy/routes"
	AuthService "ecobuy/services/auth"
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
	authRepo := AuthRepositories.NewAuthRepository(db)
	authService := AuthService.NewAuthService(authRepo)
	authController := auth.NewAuthController(authService)

	routeController := routes.RouteController{
		AuthController: *authController,
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
