package main

import (
	"ecobuy/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	config.RunMigrations(db)

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
