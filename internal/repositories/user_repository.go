package repositories

import (
	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/jinzhu/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(db *gorm.DB, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
