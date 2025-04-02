package main

import (
	"log"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func init() {
	dsn := "user=postgres password=postgres dbname=go sslmode=disable"
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()

	routes.SetupRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
