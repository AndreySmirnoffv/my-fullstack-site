package services

import (
	"fmt"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/repositories"
	"github.com/AndreySmirnoffv/my-fullstack-site/pkg/utils"
	"github.com/jinzhu/gorm"
)

func RegisterUser(db *gorm.DB, user *models.User) (string, error) {
	existingUser, err := repositories.GetUserByEmail(db, user.Email)
	if err != nil && err.Error() != "record not found" {
		return "", fmt.Errorf("error checking existing user: %v", err)
	}

	if existingUser != nil {
		return "", fmt.Errorf("user with email %s already exists", user.Email)
	}

	if err := user.HashPassword(); err != nil {
		return "", fmt.Errorf("error hashing password: %v", err)
	}

	if err := repositories.SaveUser(db, user); err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("error generating JWT: %v", err)
	}

	return token, nil
}

func LoginUser(db *gorm.DB, email, password string) (string, error) {
	user, err := repositories.GetUserByEmail(db, email)
	if err != nil {
		return "", fmt.Errorf("error finding user: %v", err)
	}

	if !user.CheckPassword(password) {
		return "", fmt.Errorf("incorrect password")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("error generating JWT: %v", err)
	}

	return token, nil
}
