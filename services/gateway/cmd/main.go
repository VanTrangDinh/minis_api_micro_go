package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"minisapi/services/gateway/docs"
	_ "minisapi/services/gateway/docs" // Import swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ServiceConfig holds the configuration for a microservice
type ServiceConfig struct {
	Name    string
	BaseURL string
	Timeout time.Duration
}

// @title           MinisAPI Gateway
// @version         1.0
// @description     Gateway service for MinisAPI
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize service configurations
	services := map[string]ServiceConfig{
		"auth": {
			Name:    "auth-service",
			BaseURL: os.Getenv("AUTH_SERVICE_URL"),
			Timeout: 5 * time.Second,
		},
		// Add other services here
	}

	// Set up Gin router
	r := gin.Default()

	// Add middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register Swagger
	docs.RegisterSwagger(r)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// Metrics endpoint for Prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	api := r.Group("/api")
	{
		// Auth service routes
		auth := api.Group("/auth")
		{
			// @Summary      Register a new user
			// @Description  Register a new user with the provided information
			// @Tags         auth
			// @Accept       json
			// @Produce      json
			// @Param        user  body      models.RegisterRequest  true  "User registration information"
			// @Success      201   {object}  models.RegisterResponse
			// @Failure      400   {object}  models.ErrorResponse
			// @Failure      500   {object}  models.ErrorResponse
			// @Router       /auth/register [post]
			auth.POST("/register", proxyRequest(services["auth"], "POST", "/api/register"))

			// @Summary      Login user
			// @Description  Login with username and password
			// @Tags         auth
			// @Accept       json
			// @Produce      json
			// @Param        credentials  body      models.LoginRequest  true  "Login credentials"
			// @Success      200          {object}  models.LoginResponse
			// @Failure      400          {object}  models.ErrorResponse
			// @Failure      401          {object}  models.ErrorResponse
			// @Failure      500          {object}  models.ErrorResponse
			// @Router       /auth/login [post]
			auth.POST("/login", proxyRequest(services["auth"], "POST", "/api/login"))

			// Protected routes
			protected := auth.Group("")
			protected.Use(authMiddleware())
			{
				// @Summary      Get current user
				// @Description  Get information about the currently logged-in user
				// @Tags         auth
				// @Accept       json
				// @Produce      json
				// @Security     BearerAuth
				// @Success      200  {object}  models.User
				// @Failure      401  {object}  models.ErrorResponse
				// @Failure      500  {object}  models.ErrorResponse
				// @Router       /auth/me [get]
				protected.GET("/me", proxyRequest(services["auth"], "GET", "/api/me"))

				// @Summary      Get user roles
				// @Description  Get all roles for the current user
				// @Tags         auth
				// @Accept       json
				// @Produce      json
				// @Security     BearerAuth
				// @Success      200  {array}   models.Role
				// @Failure      401  {object}  models.ErrorResponse
				// @Failure      500  {object}  models.ErrorResponse
				// @Router       /auth/roles [get]
				protected.GET("/roles", proxyRequest(services["auth"], "GET", "/api/roles"))

				// @Summary      Get user permissions
				// @Description  Get all permissions for the current user
				// @Tags         auth
				// @Accept       json
				// @Produce      json
				// @Security     BearerAuth
				// @Success      200  {array}   models.Permission
				// @Failure      401  {object}  models.ErrorResponse
				// @Failure      500  {object}  models.ErrorResponse
				// @Router       /auth/permissions [get]
				protected.GET("/permissions", proxyRequest(services["auth"], "GET", "/api/permissions"))
			}
		}
	}

	// Get port from environment variable or use default
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Gateway service starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// proxyRequest creates a handler that proxies requests to the target service
func proxyRequest(service ServiceConfig, method, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create HTTP client with timeout
		client := &http.Client{
			Timeout: service.Timeout,
		}

		// Create request
		req, err := http.NewRequest(method, service.BaseURL+path, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
		}

		// Copy headers from original request
		for key, values := range c.Request.Header {
			// Skip Authorization header as we'll handle it separately
			if key == "Authorization" {
				continue
			}
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}

		// Get Authorization header and add Bearer prefix if needed
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			if !strings.HasPrefix(authHeader, "Bearer ") {
				authHeader = "Bearer " + authHeader
			}
			req.Header.Set("Authorization", authHeader)
		}

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to proxy request"})
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		for key, values := range resp.Header {
			for _, value := range values {
				c.Header(key, value)
			}
		}

		// Send response
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}

// authMiddleware validates the JWT token
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Code":    2012,
				"Message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Add Bearer prefix if not present
		if !strings.HasPrefix(token, "Bearer ") {
			token = "Bearer " + token
		}

		// Pass the token to auth service
		c.Header("Authorization", token)
		c.Next()
	}
}
