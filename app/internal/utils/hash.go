package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const DEFAULT_COST = 12

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), DEFAULT_COST)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(passwordHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
