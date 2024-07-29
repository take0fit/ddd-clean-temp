package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("AUTH_DB_USER")
	dbPassword := os.Getenv("AUTH_DB_PASSWORD")
	dbHost := os.Getenv("AUTH_DB_HOST")
	dbPort := os.Getenv("AUTH_DB_PORT")
	dbName := os.Getenv("AUTH_DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Unable to create database connection: %v\n", err)
	}
	defer sqlDB.Close()

	insertTestData(sqlDB)
}

func insertTestData(db *sql.DB) {
	authUsers := []struct {
		userID   int
		email    string
		password string
	}{
		{1, "authuser1@example.com", "password1"},
		{2, "authuser2@example.com", "password2"},
	}

	for _, authUser := range authUsers {
		passwordHash, err := hashPassword(authUser.password)
		if err != nil {
			log.Fatalf("Failed to hash password: %v\n", err)
		}

		_, err = db.Exec(
			`INSERT INTO auth_users (user_id, email, password_hash, created_at, updated_at)
             VALUES ($1, $2, $3, $4, $5)`,
			authUser.userID, authUser.email, passwordHash, time.Now(), time.Now())

		if err != nil {
			log.Fatalf("Failed to insert test data into auth_users table: %v\n", err)
		}
	}

	log.Println("Test data inserted successfully into auth_users table!")
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
