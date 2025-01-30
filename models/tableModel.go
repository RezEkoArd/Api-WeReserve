package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	ID	uint `gorm:"primaryKey"`
	TableNumber string	`gorm:"unique;notnull"`
	Capacity	int		`gorm:"notnull"`;
	Status		string	`gorm:"default:'available'"`
	Reservation	[]Reservation
}