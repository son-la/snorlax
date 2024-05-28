package repositories

import (
	"github.com/son-la/snorlax/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo ..
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindByEmail(email string) *models.User {
	var user *models.User
	r.db.Where("email = ?", email).First(&user)

	return user

}

// Save ..
func (r *UserRepo) Save(user *models.User) error {

	record := r.db.Save(&user)
	return record.Error
}
