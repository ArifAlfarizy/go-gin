package config

import (
	"fmt"
	"log"
	"os"

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
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		dbname,
		port,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Database connected successfully")
}
