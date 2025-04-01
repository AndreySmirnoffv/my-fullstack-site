package handlers

import (
	"net/http"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateUserHandler(c *gin.Context, db *gorm.DB) {
	var requestUser models.RequestUser

	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := requestUser.ToUser()

	createdUser, err := models.CreateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created succesfully",
		"user":    createdUser,
	})
}
