package email

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	gomail "gopkg.in/mail.v2"
)

func GenerateVerificationCode() int {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGenerator.Intn(900000) + 100000
}

func SendVerificationEmail(email string, code int) error {
	EMAIL := os.Getenv("EMAIL")
	EMAIL_PASSWORD := os.Getenv(EMAIL)
	message := gomail.NewMessage()
	message.SetHeader("From", "smirnoffa675@gmail.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Verification Code")
	message.SetBody("text/plain", fmt.Sprintf("Your verification code is: %d", code))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, EMAIL, EMAIL_PASSWORD)

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
