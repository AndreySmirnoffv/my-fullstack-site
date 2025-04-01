package routes

import (
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		api.POST("/register", func(c *gin.Context) {
			handlers.CreateUserHandler(c, db)
		})
	}
}
