package migrate

import (
	"goAPI/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDB() {
	DB.AutoMigrate(&models.Users{}, &models.Photos{})
}
