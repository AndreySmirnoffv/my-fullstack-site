package routes

import (
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/handlers"
	"github.com/AndreySmirnoffv/my-fullstack-site/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		api.POST("/register", func(c *gin.Context) {
			handlers.CreateUserHandler(c, db)
		})
		api.POST("/login", func(c *gin.Context) {
			handlers.LoginUserHandler(c, db)
		})

		protected := api.Group("/protected")
		protected.Use(middlewares.VerifyJWTMiddleware())
		{
			protected.GET("/profile", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Welcome to your profile!"})
			})
		}
	}
}
