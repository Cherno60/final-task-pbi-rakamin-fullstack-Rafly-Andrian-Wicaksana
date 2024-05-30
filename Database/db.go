package database

import (
	"fmt"
	models "goAPI/Models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DBConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}
}
func SyncDB() {
	DB.AutoMigrate(&models.Users{}, &models.Photos{})
}
