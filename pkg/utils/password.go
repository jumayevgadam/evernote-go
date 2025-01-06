package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword func hashes password.
func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}

	return string(hashedPass), nil
}

// CheckAndComparePassword func checks and compares password.
func CheckAndComparePassword(password string, hashedPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		return fmt.Errorf("error comparing password: %w", err)
	}

	return nil
}
