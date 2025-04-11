package main

import (
	"log"
	"net/http"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/infrastructure/redis"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var (
	db          *gorm.DB
	err         error
	redisClient = redis.InitRedis()
)

func init() {
	initDB()
	redis.InitRedis()
}

func initDB() {
	dsn := "user=postgres password=postgres dbname=go sslmode=disable"
	db, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db.AutoMigrate(&models.User{})
}

func main() {
	r := gin.Default()

	cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	})

	r.OPTIONS("/api/auth/*action", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Status(http.StatusOK)
	})

	routes.SetupRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
