package main

import (
	"log"
	"os"

	_ "minisapi/services/auth/docs"
	"minisapi/services/auth/internal/configs"
	"minisapi/services/auth/internal/infrastructure/database"
	"minisapi/services/auth/internal/interfaces/http/routes"

	"github.com/joho/godotenv"
)

// @title           MinisAPI Auth Service
// @version         1.0
// @description     A JWT-based authentication service with RBAC support.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description JWT token
// @tokenUrl /login
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Log environment variables
	log.Printf("DB_HOST: %s", os.Getenv("DB_HOST"))
	log.Printf("DB_PORT: %s", os.Getenv("DB_PORT"))
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))

	// Initialize container
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Check if database connection is successful
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Test database connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Printf("Successfully connected to database")

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Printf("Successfully ran migrations")

	// Initialize HTTP server
	server := routes.NewServer(cfg, db)

	// Start server
	log.Printf("Auth service starting on port %s", cfg.Server.Port)
	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
