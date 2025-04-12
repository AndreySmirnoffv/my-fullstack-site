package email

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	gomail "gopkg.in/mail.v2"
)

func GenerateVerificationCode() int {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randGenerator.Intn(900000) + 100000
}

func SendEmailVerificationCode(to string, code int) error {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	message := gomail.NewMessage()
	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Your Verification Code")
	message.SetBody("text/plain", fmt.Sprintf("Your verification code is: %d", code))

	dialer := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := dialer.DialAndSend(message); err != nil {
		log.Println(to)
		log.Println("❌ Ошибка при отправке письма:", err)
		return err
	} else {
		log.Println("✔️ Письмо отправлено успешно.")
	}

	return nil
}
