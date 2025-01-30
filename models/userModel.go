package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser Role = "customer"

)

type User struct {
	*gorm.Model
	ID 			uint	`gorm:"primaryKey"`
	Name 		string `gorm:"unique"`
	Email 		string `gorm:"unique"`
	Password	string	
	Role		Role 	`gorm:"default:'customer';not null"`
	Reserve		[]Reservation	`gorm:"foreignKey:UserID"`
}