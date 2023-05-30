package database

import (
	"log"
	"os"

	"github.com/RianIhsan/ApiGoJwt/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Tidak Bisa connect ke database")

	} else {
		log.Println("Berhasil terhubung ke database")
	}

	DB = database

	database.AutoMigrate(&models.User{})
}
