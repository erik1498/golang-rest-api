package database

import (
	"fmt"
	"log"
	"rest-api/config"
	"rest-api/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var DB_USER = config.Config("DB_USER")
	var DB_PASSWORD = config.Config("DB_PASSWORD")
	var DB_HOST = config.Config("DB_HOST")
	var DB_PORT = config.Config("DB_PORT")
	var DB_NAME = config.Config("DB_NAME")

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database Idiot")
	}

	log.Println("Database Connect")

	DB.AutoMigrate(&model.Kelas{})
	log.Println("Migrate Success")
}
