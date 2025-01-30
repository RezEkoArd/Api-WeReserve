package repository

import (
	"wereserve/models"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	DB *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository{
	return &ReservationRepository{DB: db}
}


// Create Reserve
func (r *ReservationRepository) CreateReservation(reservation *models.Reservation) (*models.Reservation, error) {
	err := r.DB.Create(reservation).Error
    if err != nil {
        return nil, err
    }

	//Ambil ulang data dengan preload
	var newReservation models.Reservation 
	err = r.DB.Preload("User").Preload("Table").First(&newReservation, reservation.ID).Error
	return &newReservation, err
}

//	Get User Reservation
func (r *ReservationRepository) GetUserReservation(id uint) ([]models.Reservation, error){
	var reservations []models.Reservation
	if err := r.DB.Preload("User").Preload("Table").Where("user_id = ?", id).
	Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

// Update Status Table on Reservation
func (r *ReservationRepository) UpdateTableStatus(idTable uint, status string) error {
	return r.DB.Model(&models.Table{}).Where("id = ?", idTable).Update("status", status).Error
}


// A
func (r *ReservationRepository) GetReservationByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation

    // Gunakan Preload untuk mengambil relasi User dan Table
    err := r.DB.Preload("User").Preload("Table").First(&reservation, id).Error
    if err != nil {
        return nil, err
    }

    return &reservation, nil
}


// Get All Reserve
func (r *ReservationRepository) GetAllReserver() ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := r.DB.Preload("User").Preload("Table").Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}