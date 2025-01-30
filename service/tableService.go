package service

import (
	"wereserve/models"
	"wereserve/repository"
)

type TableService struct {
	TableRepo *repository.TableRepository
}

func NewTableService(tableRepo *repository.TableRepository) *TableService{
	return &TableService{TableRepo: tableRepo}
}

//ADMIN ROLE
func (s *TableService) CreateTable(table_number,status string, capacity int) (*models.Table, error) {
	//create table
	table := &models.Table{
		TableNumber: table_number,
		Capacity:    capacity,
		Status:      status,
	}

	// saved to DB
	err := s.TableRepo.CreateTable(table)
	if err != nil {
		return nil, err
	}

	return table,err
}

func (s *TableService) UpdateStatusTable(id uint, updateStatus *models.Table) (*models.Table, error) {
	// Cari ID table
	table, err := s.TableRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// update Status Table
	table.Status = updateStatus.Status

	// saved to dB
	if err := s.TableRepo.UpdateTable(table); err != nil {
		return nil, err
	}
	return table, err
}

func (s *TableService) DeleteTable(id uint) error{
	return s.TableRepo.DeleteTable(id)
}

// USER ROLE
func (s *TableService) FindAllTables() ([]models.Table, error){
	return s.TableRepo.FindAllTable()
}

func (s *TableService) FindByID(id uint)  (*models.Table, error) {
	return s.TableRepo.FindByID(id)
}