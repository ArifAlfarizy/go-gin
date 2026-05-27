package main

import (
	"gin-test/config"
	"gin-test/internal/migrates"
	"gin-test/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.LoadEnv()

	config.ConnectDB()

	migrates.Migrate()

	routes.TodoRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
