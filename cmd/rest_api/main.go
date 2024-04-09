package main

import (
	"github.com/gin-gonic/gin"
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
	gin.SetMode(os.Getenv("GIN_MODE"))
	engine := gin.Default()

	http.RegisterHandlers(engine)

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRDB() {
	if err := db.RDBConnect(); err != nil {
		log.Fatal(err)
	}
}
