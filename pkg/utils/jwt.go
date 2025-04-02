package utils

import (
	"os"
	"time"

	"github.com/AndreySmirnoffv/my-fullstack-site/internal/models"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return t, nil

}
