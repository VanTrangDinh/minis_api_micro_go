package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"minisapi/services/auth/internal/configs"
	"minisapi/services/auth/internal/domain/service"
	"minisapi/services/auth/internal/domain/usecase"
	"minisapi/services/auth/internal/infrastructure/database"
	"minisapi/services/auth/internal/infrastructure/jwt"
	"minisapi/services/auth/internal/infrastructure/redis"
	"minisapi/services/auth/internal/infrastructure/repository"
	"minisapi/services/auth/internal/interfaces/http/handler"
	"minisapi/services/auth/internal/interfaces/http/routes"
	"minisapi/services/auth/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	log, err := logger.NewLogger(cfg.Log.Level)
	if err != nil {
		log.Fatal(context.Background(), "Failed to initialize logger", logger.LogFields{
			Error: err,
		})
	}
	defer log.Sync()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal(context.Background(), "Failed to connect to database", logger.LogFields{
			Error: err,
		})
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(context.Background(), "Failed to get database instance", logger.LogFields{
			Error: err,
		})
	}
	defer sqlDB.Close()

	// Initialize Redis
	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatal(context.Background(), "Failed to connect to Redis", logger.LogFields{
			Error: err,
		})
	}
	defer redisClient.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	// Initialize JWT manager
	jwtManager, err := jwt.NewJWTManager(cfg.JWT)
	if err != nil {
		log.Fatal(context.Background(), "Failed to initialize JWT manager", logger.LogFields{
			Error: err,
		})
	}

	// Initialize use cases
	authService := service.NewAuthService(jwtManager)
	authUseCase := usecase.NewAuthUseCase(
		userRepo,
		tokenRepo,
		sessionRepo,
		authService,
	)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	roleHandler := handler.NewRoleHandler(roleRepo, permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionRepo)
	healthHandler := handler.NewHealthHandler(db)

	// Initialize router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(
		router,
		authHandler,
		roleHandler,
		permissionHandler,
		healthHandler,
		log,
	)

	// Create server
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Info(context.Background(), "Starting server", logger.LogFields{
			Path: cfg.Server.Port,
		})
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(context.Background(), "Failed to start server", logger.LogFields{
				Error: err,
			})
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	log.Info(context.Background(), "Shutting down server", logger.LogFields{})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error(context.Background(), "Server forced to shutdown", logger.LogFields{
			Error: err,
		})
	}

	log.Info(context.Background(), "Server exiting", logger.LogFields{})
}
