package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) Check(c *gin.Context) {
	health := map[string]interface{}{
		"status": "ok",
		"time":   time.Now().Format(time.RFC3339),
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		health["database"] = map[string]interface{}{
			"status": "error",
			"error":  err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, health)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		health["database"] = map[string]interface{}{
			"status": "error",
			"error":  err.Error(),
		}
		c.JSON(http.StatusServiceUnavailable, health)
		return
	}

	health["database"] = map[string]interface{}{
		"status": "ok",
	}

	c.JSON(http.StatusOK, health)
}
