package config

import (
	"financify/pkg/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("POSTGRES_URI")
	log.Println(uri)
	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDB() {
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Category{})
    db.Model(&models.User{}).Association("Transactions").Append(&models.Transaction{})
}
