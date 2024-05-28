package models

import (
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	Name     string `json:"name" binding:"required"`
	Username string `json:"username" gorm:"unique" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required"`
}

type UserRepository interface {
	FindByEmail(email string) *User
	Save(user *User) error
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
