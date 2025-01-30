package service

import (
	"wereserve/models"
	"wereserve/repository"
)

type ReserveService struct {
	ReserveRepo *repository.ReservationRepository
}

func NewReserveService(reserveRepo *repository.ReservationRepository) *ReserveService{
	return &ReserveService{ReserveRepo: reserveRepo}
}

// Create Reserver
func (s *ReserveService) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	//Create Reserver
	createReservation, err := s.ReserveRepo.CreateReservation(reservation)
	if err != nil {
		return nil, err
	}

	if err := s.ReserveRepo.UpdateTableStatus(reservation.TableID, "booked"); err != nil {
		return nil, err
	}

	// ambil data ulang
	return s.ReserveRepo.GetReservationByID(createReservation.ID)
}

// Get MY Reserve
func (s *ReserveService) GetUserReservation(id uint) ([]models.Reservation, error) {
	return s.ReserveRepo.GetUserReservation(id)
}

// Get All
func (s *ReserveService) GetAllReserve() ([]models.Reservation, error) {
	return s.ReserveRepo.GetAllReserver()
}