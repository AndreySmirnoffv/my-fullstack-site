package services

import (
	"sync"
	"time"

	"github.com/AndreySmirnoffv/my-fullstack-site/pkg/services/email"
)

var verificationCodes = make(map[string]int)
var mu sync.Mutex

func GenerateAndStoreCode(emailAddress string) int {
	code := email.GenerateVerificationCode()

	mu.Lock()
	verificationCodes[emailAddress] = code
	mu.Unlock()

	// Удаляем код через 5 минут
	go func() {
		time.Sleep(5 * time.Minute)
		mu.Lock()
		delete(verificationCodes, emailAddress)
		mu.Unlock()
	}()

	return code
}

func VerifyCode(emailAddress string, code int) bool {
	mu.Lock()
	defer mu.Unlock()

	storedCode, exists := verificationCodes[emailAddress]
	if !exists || storedCode != code {
		return false
	}

	delete(verificationCodes, emailAddress)
	return true
}

func SendVerification(emailAddr string) error {
	code := GenerateAndStoreCode(emailAddr)
	return email.SendEmailVerificationCode(emailAddr, code)
}
