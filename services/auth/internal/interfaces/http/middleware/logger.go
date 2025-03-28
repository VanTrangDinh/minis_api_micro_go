package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"minisapi/services/auth/internal/pkg/logger"
)

type LoggerMiddleware struct {
	logger *logger.Logger
}

func NewLoggerMiddleware(logger *logger.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger: logger,
	}
}

func (m *LoggerMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()

		// Set request ID in context
		c.Set("request_id", requestID)

		// Process request
		c.Next()

		// Log request details
		fields := logger.LogFields{
			RequestID: requestID,
			IP:        c.ClientIP(),
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Status:    c.Writer.Status(),
			Duration:  time.Since(start),
		}

		// Get user ID if available
		if userID, exists := c.Get("user_id"); exists {
			fields.UserID = userID.(uint)
		}

		// Log based on status code
		switch {
		case c.Writer.Status() >= 500:
			m.logger.Error(c.Request.Context(), "Server error", fields)
		case c.Writer.Status() >= 400:
			m.logger.Warn(c.Request.Context(), "Client error", fields)
		default:
			m.logger.Info(c.Request.Context(), "Request processed", fields)
		}
	}
}
