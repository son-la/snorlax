package database

import (
	"github.com/son-la/snorlax/internal/models"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

func NewMySQLDB(connectionString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
