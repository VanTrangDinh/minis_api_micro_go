package routes

import (
	"net/http"

	"minisapi/services/notification/internal/container"
	"minisapi/services/notification/internal/handlers"
	"minisapi/services/notification/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(r *gin.Engine, container *container.Container) {
	// Add metrics middleware to all routes
	r.Use(middleware.MetricsMiddleware())

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Metrics endpoint for Prometheus
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Initialize handlers
	notificationHandler := handlers.NewNotificationHandler(container)

	// API routes
	api := r.Group("/api")
	{
		// Notification routes
		notifications := api.Group("/notifications")
		notifications.Use(middleware.AuthMiddleware())
		{
			// @Summary      Send email notification
			// @Description  Send an email to specified recipients
			// @Tags         notifications
			// @Accept       json
			// @Produce      json
			// @Param        notification  body      models.EmailRequest  true  "Email notification details"
			// @Success      200           {object}  models.Response
			// @Failure      400           {object}  models.ErrorResponse
			// @Failure      500           {object}  models.ErrorResponse
			// @Router       /notifications/email [post]
			notifications.POST("/email", notificationHandler.SendEmail)

			// @Summary      Send SMS notification
			// @Description  Send an SMS to specified phone number
			// @Tags         notifications
			// @Accept       json
			// @Produce      json
			// @Param        notification  body      models.SMSRequest  true  "SMS notification details"
			// @Success      200          {object}  models.Response
			// @Failure      400          {object}  models.ErrorResponse
			// @Failure      500          {object}  models.ErrorResponse
			// @Router       /notifications/sms [post]
			notifications.POST("/sms", notificationHandler.SendSMS)

			// @Summary      Send push notification
			// @Description  Send a push notification to specified device
			// @Tags         notifications
			// @Accept       json
			// @Produce      json
			// @Param        notification  body      models.PushRequest  true  "Push notification details"
			// @Success      200          {object}  models.Response
			// @Failure      400          {object}  models.ErrorResponse
			// @Failure      500          {object}  models.ErrorResponse
			// @Router       /notifications/push [post]
			notifications.POST("/push", notificationHandler.SendPush)

			// @Summary      Get notification status
			// @Description  Get the status of a notification by ID
			// @Tags         notifications
			// @Accept       json
			// @Produce      json
			// @Param        id  path      string  true  "Notification ID"
			// @Success      200  {object}  models.NotificationStatus
			// @Failure      404  {object}  models.ErrorResponse
			// @Failure      500  {object}  models.ErrorResponse
			// @Router       /notifications/{id}/status [get]
			notifications.GET("/:id/status", notificationHandler.GetStatus)

			// @Summary      List notifications
			// @Description  Get a list of notifications with pagination
			// @Tags         notifications
			// @Accept       json
			// @Produce      json
			// @Param        page  query     int  false  "Page number"
			// @Param        limit query     int  false  "Items per page"
			// @Success      200   {object}  models.NotificationList
			// @Failure      500   {object}  models.ErrorResponse
			// @Router       /notifications [get]
			notifications.GET("", notificationHandler.List)
		}
	}
}
