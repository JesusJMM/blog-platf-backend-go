package users

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt the password
func EncryptPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

// Check if the given password is correct
// return true on success and false on failure
func ValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
