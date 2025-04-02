package email

import (
	"fmt"
	"math/rand"
	"time"

	gomail "gopkg.in/mail.v2"
)

func GenerateVerificationCode() int {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGenerator.Intn(900000) + 100000
}

func SendVerificationEmail(email string, code int) error {
	message := gomail.NewMessage()
	message.SetHeader("From", "your-email@gmail.com")
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Verification Code")
	message.SetBody("text/plain", fmt.Sprintf("Your verification code is: %d", code))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "your-email@gmail.com", "your-email-password")

	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
