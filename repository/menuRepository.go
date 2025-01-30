package repository

import (
	"wereserve/models"

	"gorm.io/gorm"
)


type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository{
	return &MenuRepository{DB: db}
}

// Admin Role
func (r *MenuRepository) CreateMenu(menu *models.Menu) error {
	result := r.DB.Create(menu)
	return result.Error
}

func (r *MenuRepository) DeleteMenu(id uint) error {
	return r.DB.Delete(&models.Menu{}, id).Error
}

//

func (r *MenuRepository) GetAllMenu() ([]models.Menu, error) {
	var menus []models.Menu
	if err := r.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}