package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/take0fit/ddd-clean-temp/internal/common/db"
	"github.com/take0fit/ddd-clean-temp/internal/common/http"
	"log"
	"os"
)

func main() {
	setupRDB()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")
	pemBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("Failed to read public key file: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))
	engine := gin.Default()

	http.RegisterHandlers(engine, publicKey)

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB() {
	if err := db.RDBConnect(); err != nil {
		log.Fatal(err)
	}
}
