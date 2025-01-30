package initializer

import "wereserve/models"

func SyncMigrate () {
	DB.AutoMigrate(&models.User{},&models.Table{},&models.Reservation{},models.Menu{});
}