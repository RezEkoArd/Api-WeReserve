package models

import "gorm.io/gorm"

type Menu struct {
	*gorm.Model
	ID 			uint	`gorm:"primaryKey"`
	Name		string	`gorm:"notnull"`
	Description	string	`gorm:"notnull"`
	Price		string	`gorm:"notnull"`
}
