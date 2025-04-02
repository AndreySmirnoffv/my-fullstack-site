package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your-secret-key")

type RequestUser struct {
	ID       int     `json:"id"`
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) CheckPassword(providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	return err == nil
}

func (ru *RequestUser) ToUser() *User {
	return &User{
		Name:  *ru.Name,
		Email: ru.Email,
	}
}
