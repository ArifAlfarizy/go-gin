package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct { // JSON tag				// GORM tag
	ID        uint   `json:"id"` // field named `ID` will be used as a primary field by default
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=alfarizy dbname=gin_test port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Database connected successfully")
}
