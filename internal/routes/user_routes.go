package routes

import (
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		api.POST("/send-verification", handlers.SendVerificationHandler)
		api.POST("/verify-code", handlers.VerifyCodeHandler)

		api.POST("/register", handlers.RegisterUserHandler(db))
		api.POST("/login", handlers.LoginHandler(db))
	}
}
