package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
    *gorm.Model
    UserID         uint   `gorm:"not null"`
    TableID        uint   `gorm:"not null"`
    Date           time.Time `gorm:"type:date;not null"`
    Time           time.Time `gorm:"type:time;not null"`
    NumberOfPeople int    `gorm:"not null"`
    User           User   `gorm:"foreignKey:UserID"`
    Table          Table  `gorm:"foreignKey:TableID"`
}
