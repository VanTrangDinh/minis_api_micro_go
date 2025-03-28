package routes

import (
	"context"
	"fmt"
	"net/http"

	"minisapi/services/auth/internal/configs"
	"minisapi/services/auth/internal/domain/service"
	"minisapi/services/auth/internal/domain/usecase"
	"minisapi/services/auth/internal/infrastructure/jwt"
	"minisapi/services/auth/internal/infrastructure/redis"
	"minisapi/services/auth/internal/infrastructure/repository"
	"minisapi/services/auth/internal/interfaces/http/handler"
	"minisapi/services/auth/internal/interfaces/http/middleware"
	"minisapi/services/auth/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *configs.Config, db *gorm.DB) *Server {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewTokenRepository(db)
	sessionRepo := repository.NewSessionRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	// Initialize Redis client
	if _, err := redis.NewRedisClient(cfg.Redis); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	// Initialize JWT manager
	jwtManager, err := jwt.NewJWTManager(cfg.JWT)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize JWT manager: %v", err))
	}

	// Initialize auth service
	authService := service.NewAuthService(jwtManager)

	// Initialize use cases
	authUseCase := usecase.NewAuthUseCase(userRepo, tokenRepo, sessionRepo, authService)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authUseCase)
	healthHandler := handler.NewHealthHandler(db)
	roleHandler := handler.NewRoleHandler(roleRepo, permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionRepo)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(10), 100)

	log, err := logger.NewLogger("auth-service")
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	loggerMiddleware := middleware.NewLoggerMiddleware(log)

	// Initialize router
	router := gin.Default()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(loggerMiddleware.Logger())
	router.Use(middleware.RateLimit(rateLimiter))

	// Health check
	router.GET("/health", healthHandler.Check)

	// Public routes
	public := router.Group("/api/v1/auth")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.POST("/refresh", authHandler.RefreshToken)
		public.POST("/reset-password", authHandler.ResetPassword)
		public.POST("/verify-email", authHandler.VerifyEmail)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(authMiddleware.Authenticate())
	{
		// Auth routes
		auth := protected.Group("/auth")
		{
			auth.POST("/logout", authHandler.Logout)
			auth.POST("/change-password", authHandler.ChangePassword)
			auth.POST("/enable-2fa", authHandler.EnableTwoFactor)
			auth.POST("/disable-2fa", authHandler.DisableTwoFactor)
			auth.POST("/verify-2fa", authHandler.VerifyTwoFactor)
		}

		// Role routes
		roles := protected.Group("/roles")
		roles.Use(authMiddleware.RequireRole("admin"))
		{
			roles.POST("", roleHandler.Create)
			roles.GET("", roleHandler.List)
			roles.GET("/:id", roleHandler.Get)
			roles.PUT("/:id", roleHandler.Update)
			roles.DELETE("/:id", roleHandler.Delete)
		}

		// Permission routes
		permissions := protected.Group("/permissions")
		permissions.Use(authMiddleware.RequireRole("admin"))
		{
			permissions.POST("", permissionHandler.Create)
			permissions.GET("", permissionHandler.List)
			permissions.GET("/:id", permissionHandler.Get)
			permissions.PUT("/:id", permissionHandler.Update)
			permissions.DELETE("/:id", permissionHandler.Delete)
		}
	}

	// Admin routes
	admin := router.Group("/api/v1/admin")
	admin.Use(authMiddleware.Authenticate())
	admin.Use(authMiddleware.RequireRole("admin"))
	{
		// Add admin routes here
	}

	// Initialize HTTP server
	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func SetupRoutes(
	router *gin.Engine,
	authHandler *handler.AuthHandler,
	roleHandler *handler.RoleHandler,
	permissionHandler *handler.PermissionHandler,
	healthHandler *handler.HealthHandler,
	logger *logger.Logger,
) {
	// Initialize middleware
	rateLimiter := middleware.NewIPRateLimiter(rate.Limit(10), 100)
	loggerMiddleware := middleware.NewLoggerMiddleware(logger)

	// Initialize JWT manager and auth service
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}
	jwtManager, err := jwt.NewJWTManager(cfg.JWT)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize JWT manager: %v", err))
	}
	authService := service.NewAuthService(jwtManager)
	authMiddleware := middleware.NewAuthMiddleware(authService)

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(loggerMiddleware.Logger())
	router.Use(middleware.RateLimit(rateLimiter))

	// Health check
	router.GET("/health", healthHandler.Check)

	// Public routes
	public := router.Group("/api/v1")
	{
		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/forgot-password", authHandler.ForgotPassword)
			auth.POST("/reset-password", authHandler.ResetPassword)
			auth.POST("/verify-email", authHandler.VerifyEmail)
		}
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(authMiddleware.Authenticate())
	{
		// Auth routes
		auth := protected.Group("/auth")
		{
			auth.POST("/logout", authHandler.Logout)
			auth.POST("/change-password", authHandler.ChangePassword)
			auth.POST("/enable-2fa", authHandler.EnableTwoFactor)
			auth.POST("/disable-2fa", authHandler.DisableTwoFactor)
			auth.POST("/verify-2fa", authHandler.VerifyTwoFactor)
		}

		// Role routes
		roles := protected.Group("/roles")
		roles.Use(authMiddleware.RequireRole("admin"))
		{
			roles.POST("", roleHandler.Create)
			roles.GET("", roleHandler.List)
			roles.GET("/:id", roleHandler.Get)
			roles.PUT("/:id", roleHandler.Update)
			roles.DELETE("/:id", roleHandler.Delete)
		}

		// Permission routes
		permissions := protected.Group("/permissions")
		permissions.Use(authMiddleware.RequireRole("admin"))
		{
			permissions.POST("", permissionHandler.Create)
			permissions.GET("", permissionHandler.List)
			permissions.GET("/:id", permissionHandler.Get)
			permissions.PUT("/:id", permissionHandler.Update)
			permissions.DELETE("/:id", permissionHandler.Delete)
		}
	}
}
