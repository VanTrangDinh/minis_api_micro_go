package main

import (
	"log"
	"os"

	_ "minisapi/services/notification/docs"
	"minisapi/services/notification/internal/config"
	"minisapi/services/notification/internal/container"
	"minisapi/services/notification/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           MinisAPI Notification Service
// @version         1.0
// @description     A notification service for handling emails, SMS, and push notifications.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8082
// @BasePath  /api

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize config
	cfg := config.New()

	// Initialize container
	container := container.NewContainer(cfg)

	// Set up Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, container)

	// Get port from environment variable or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8082"
	}

	// Start server
	log.Printf("Notification service starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
