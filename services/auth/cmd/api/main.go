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
	"minisapi/services/auth/internal/infrastructure/database"
	"minisapi/services/auth/internal/infrastructure/jwt"
	"minisapi/services/auth/internal/infrastructure/redis"
	"minisapi/services/auth/internal/interfaces/http/handler"
	"minisapi/services/auth/internal/interfaces/http/routes"
	"minisapi/services/auth/internal/pkg/logger"
	"minisapi/services/auth/internal/repository"
	"minisapi/services/auth/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger
	logger, err := logger.NewLogger(cfg.Log.Level)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		logger.Fatal(context.Background(), "Failed to connect to database", logger.LogFields{
			Error: err,
		})
	}
	defer db.Close()

	// Initialize Redis
	redisClient, err := redis.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Fatal(context.Background(), "Failed to connect to Redis", logger.LogFields{
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
	jwtManager := jwt.NewJWTManager(cfg.JWT)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(
		userRepo,
		roleRepo,
		permissionRepo,
		tokenRepo,
		sessionRepo,
		jwtManager,
		redisClient,
	)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	roleHandler := handler.NewRoleHandler(authUseCase)
	permissionHandler := handler.NewPermissionHandler(authUseCase)
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
		logger,
	)

	// Create server
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Info(context.Background(), "Starting server", logger.LogFields{
			Path: cfg.Server.Port,
		})
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(context.Background(), "Failed to start server", logger.LogFields{
				Error: err,
			})
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	logger.Info(context.Background(), "Shutting down server", logger.LogFields{})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(context.Background(), "Server forced to shutdown", logger.LogFields{
			Error: err,
		})
	}

	logger.Info(context.Background(), "Server exiting", logger.LogFields{})
}
