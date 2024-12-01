package main

import (
	"golang-web/config"
	"golang-web/internal/database"
	"golang-web/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to database
	database.ConnectDatabase()

	// Set up Gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start server
	router.Run(":" + config.AppConfig.ServerPort)
}
