package repository

import (
	"wereserve/models"

	"gorm.io/gorm"
)


type UserRepository struct {
	DB *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	result := 	r.DB.Create(user)
	return result.Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}