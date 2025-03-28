package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	VALIDATION_ERROR = gin.H{
		"error": "Validation error",
		"code":  http.StatusBadRequest,
	}

	DB_ERROR = gin.H{
		"error": "Database error",
		"code":  http.StatusInternalServerError,
	}

	PARAMS_ERROR = gin.H{
		"error": "Invalid parameters",
		"code":  http.StatusBadRequest,
	}

	PERMISSION_EXISTS = gin.H{
		"error": "Permission already exists",
		"code":  http.StatusBadRequest,
	}

	PERMISSION_NOT_FOUND = gin.H{
		"error": "Permission not found",
		"code":  http.StatusNotFound,
	}

	ROLE_EXISTS = gin.H{
		"error": "Role already exists",
		"code":  http.StatusBadRequest,
	}

	ROLE_NOT_FOUND = gin.H{
		"error": "Role not found",
		"code":  http.StatusNotFound,
	}

	INVALID_PERMISSION = gin.H{
		"error": "Invalid permission",
		"code":  http.StatusBadRequest,
	}

	UNAUTHORIZED = gin.H{
		"error": "Unauthorized",
		"code":  http.StatusUnauthorized,
	}

	FORBIDDEN = gin.H{
		"error": "Forbidden",
		"code":  http.StatusForbidden,
	}

	SUCCESS = gin.H{
		"message": "Success",
		"code":    http.StatusOK,
	}
)
