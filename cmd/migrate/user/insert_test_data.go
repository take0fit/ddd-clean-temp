package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("USER_DB_USER")
	dbPassword := os.Getenv("USER_DB_PASSWORD")
	dbHost := os.Getenv("USER_DB_HOST")
	dbPort := "5432"
	dbName := os.Getenv("USER_DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Unable to create database connection: %v\n", err)
	}
	defer sqlDB.Close()

	insertTestData(sqlDB)
}

func insertTestData(db *sql.DB) {
	users := []struct {
		firstName string
		lastName  string
		age       int
		birthday  time.Time
	}{
		{"John", "Doe", 30, time.Date(1992, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"Jane", "Smith", 25, time.Date(1997, 5, 15, 0, 0, 0, 0, time.UTC)},
	}

	for _, user := range users {
		_, err := db.Exec(
			`INSERT INTO users (first_name, last_name, age, birthday, created_at, updated_at)
             VALUES ($1, $2, $3, $4, $5, $6)`,
			user.firstName, user.lastName, user.age, user.birthday, time.Now(), time.Now())

		if err != nil {
			log.Fatalf("Failed to insert test data into users table: %v\n", err)
		}
	}

	profiles := []struct {
		userID  int
		hobbies string
		bio     string
	}{
		{1, "Reading, Traveling", "Bio for John Doe"},
		{2, "Cooking, Hiking", "Bio for Jane Smith"},
	}

	for _, profile := range profiles {
		_, err := db.Exec(
			`INSERT INTO profiles (user_id, hobbies, bio, created_at, updated_at)
             VALUES ($1, $2, $3, $4, $5)`,
			profile.userID, profile.hobbies, profile.bio, time.Now(), time.Now())

		if err != nil {
			log.Fatalf("Failed to insert test data into profiles table: %v\n", err)
		}
	}

	log.Println("Test data inserted successfully into users and profiles tables!")
}
