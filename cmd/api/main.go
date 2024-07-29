package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	authdb "github.com/take0fit/ddd-clean-temp/internal/auth/infrastructure/db"
	"log"
	"os"

	commonHttp "github.com/take0fit/ddd-clean-temp/cmd/api/http"
	userdb "github.com/take0fit/ddd-clean-temp/internal/user/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	setupRDB() // DB接続を確立

	publicKeyPath := os.Getenv("PUBLIC_KEY_PATH")
	pemBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("Failed to read public key file: %v", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		log.Fatalf("Failed to parse public key: %v", err)
	}
	engine := gin.Default()

	commonHttp.RegisterHandlers(engine, publicKey) // ルーティング情報を定義

	// 8080ポートでアプリケーションを起動
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB() {
	authDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("AUTH_DB_HOST"), os.Getenv("AUTH_DB_PORT"), os.Getenv("AUTH_DB_USER"), os.Getenv("AUTH_DB_NAME"), os.Getenv("AUTH_DB_PASSWORD"))
	if err := authdb.RDBConnect(authDSN); err != nil {
		log.Fatal(err)
	}

	userDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("USER_DB_HOST"), "5432", os.Getenv("USER_DB_USER"), os.Getenv("USER_DB_NAME"), os.Getenv("USER_DB_PASSWORD"))
	if err := userdb.RDBConnect(userDSN); err != nil {
		log.Fatal(err)
	}
}
