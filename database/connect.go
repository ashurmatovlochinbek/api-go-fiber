package database

import (
	"api-go-fiber/config"
	"api-go-fiber/internals/models"
	"fmt"
	"log"
	"strconv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	
	if err != nil {
		log.Println("Couldnt parse in ParseUint function")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Config("DB_HOST"),
		port, config.Config("DB_USER"), 
		config.Config("DB_PASSWORD"), 
		config.Config("DB_NAME"))
	
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection Opened to database")

	DB.AutoMigrate(&models.Note{})
	log.Println("Database migrated")
}