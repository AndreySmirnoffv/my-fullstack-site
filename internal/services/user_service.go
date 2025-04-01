package services

import (
	"fmt"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/repositories"
	"github.com/jinzhu/gorm"
)

func RegisterUser(db *gorm.DB, user *models.User) error {
	existingUser, err := repositories.GetUserByEmail(db, user.Email)
	if err != nil && err.Error() != "record not found" {
		return fmt.Errorf("error checking existing user: %v", err)
	}

	if existingUser != nil {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	return repositories.SaveUser(db, user)
}
