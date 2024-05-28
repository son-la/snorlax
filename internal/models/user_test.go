package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUserHashPassword(t *testing.T) {

	u := User{
		Name:     "Jari",
		Username: "jari.larson",
		Email:    "jari.larson@tomorrow.com",
	}

	passwordList := []string{
		"passwordnothashesyet",
		"password123",
	}

	for _, password := range passwordList {
		u.HashPassword(password)

		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
			t.Errorf("got %v want nil", err)
		}
	}
}
