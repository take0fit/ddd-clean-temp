package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v\n", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/user",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v\n", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v\n", err)
	}

	log.Println("User migrations applied successfully!")
}
