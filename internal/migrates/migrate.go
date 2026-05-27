package main

import (
	"gin-test/config"
	"gin-test/internal/models"
	"log"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	var err error

	err = config.DB.AutoMigrate(&models.Todo{})

	if err != nil {
		log.Fatal("Migrate failed")
	}

	log.Println("Migrate success")
}
