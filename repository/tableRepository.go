package repository

import (
	"wereserve/models"

	"gorm.io/gorm"
)

type TableRepository struct {
	DB *gorm.DB 
}

func NewTableRepository(db *gorm.DB) *TableRepository{
	return &TableRepository{DB: db}
}

//ADMIN ROLE
func (r *TableRepository) CreateTable(table *models.Table) error {
	result := r.DB.Create(table)
	return result.Error
}

func (r *TableRepository) FindByID(id uint) (*models.Table, error) {
	var table models.Table
	err := r.DB.Model(&models.Table{}).First(&table, id).Error
	if err != nil {
		return nil ,err
	}
	return &table, err
}

func (r *TableRepository) UpdateTable(table *models.Table) error {
	return r.DB.Model(&models.Table{}).Where("id = ?", table.ID).Updates(table).Error
}

func (r *TableRepository) DeleteTable(id uint) error  {
	return r.DB.Delete(&models.User{}, id).Error
}


func (r *TableRepository) FindAllTable()  ([]models.Table, error) {
	var tables []models.Table
	if err := r.DB.Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

