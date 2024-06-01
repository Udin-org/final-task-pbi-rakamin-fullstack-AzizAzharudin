package config

import (
	"log"
	"myapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/dasarcrudgo?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database 
	if err := DB.AutoMigrate(&models.User{}, &models.Photo{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}