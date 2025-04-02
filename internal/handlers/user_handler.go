package handlers

import (
	"net/http"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/services"
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

	token, err := services.RegisterUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"token":   token,
	})
}

func LoginUserHandler(c *gin.Context, db *gorm.DB) {
	var requestUser models.RequestUser
	if err := c.ShouldBindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.LoginUser(db, requestUser.Email, requestUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user logged in successfully",
		"token":   token,
	})
}
