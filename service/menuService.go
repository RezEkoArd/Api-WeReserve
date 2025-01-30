package service

import (
	"wereserve/models"
	"wereserve/repository"
)

type MenuService struct {
	MenuRepository *repository.MenuRepository
}

func NewMenuService(menuRepo *repository.MenuRepository) *MenuService{
	return &MenuService{MenuRepository: menuRepo}
}



// Admin Role
func (s *MenuService) CreateMenu(name,description,price string) (*models.Menu, error) {
	//create menu
	menu := &models.Menu{
		Name:        name,
		Description: description,
		Price:       price,
	}

	//save to db
	err := s.MenuRepository.CreateMenu(menu)
	if err != nil {
		return nil, err
	}

	return menu, err
}

func (s *MenuService) DeleteMenu(id uint) error {
	return s.MenuRepository.DeleteMenu(id)
}

// User Role

func (s *MenuService) GetAllMenu() ([]models.Menu, error) {
	return s.MenuRepository.GetAllMenu()
}