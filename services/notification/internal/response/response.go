package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents the standard API response structure
type Response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// getStatusCode maps response codes to HTTP status codes
func getStatusCode(code Code) int {
	statusMap := map[Code]int{
		BadRequest:          http.StatusBadRequest,
		Unauthorized:        http.StatusUnauthorized,
		Forbidden:           http.StatusForbidden,
		NotFound:            http.StatusNotFound,
		ValidationError:     http.StatusUnprocessableEntity,
		DuplicateEntry:      http.StatusConflict,
		InvalidCredentials:  http.StatusUnauthorized,
		AccountDisabled:     http.StatusForbidden,
		TokenExpired:        http.StatusUnauthorized,
		TokenInvalid:        http.StatusUnauthorized,
		InternalServerError: http.StatusInternalServerError,
		DatabaseError:       http.StatusInternalServerError,
		CacheError:          http.StatusServiceUnavailable,
		ServiceUnavailable:  http.StatusServiceUnavailable,
	}

	if status, ok := statusMap[code]; ok {
		return status
	}
	return http.StatusInternalServerError
}

// sendResponse is a helper function to send JSON responses
func sendResponse(c *gin.Context, status int, response Response) {
	c.JSON(status, response)
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, data interface{}) {
	sendResponse(c, http.StatusOK, Response{
		Code:    Success,
		Message: Success.Message(),
		Data:    data,
	})
}

// SendCreated sends a created response
func SendCreated(c *gin.Context, data interface{}) {
	sendResponse(c, http.StatusCreated, Response{
		Code:    Created,
		Message: Created.Message(),
		Data:    data,
	})
}

// SendUpdated sends an updated response
func SendUpdated(c *gin.Context, data interface{}) {
	sendResponse(c, http.StatusOK, Response{
		Code:    Updated,
		Message: Updated.Message(),
		Data:    data,
	})
}

// SendDeleted sends a deleted response
func SendDeleted(c *gin.Context) {
	sendResponse(c, http.StatusOK, Response{
		Code:    Deleted,
		Message: Deleted.Message(),
	})
}

// SendError sends an error response
func SendError(c *gin.Context, code Code, err interface{}) {
	status := getStatusCode(code)
	sendResponse(c, status, Response{
		Code:    code,
		Message: code.Message(),
		Error:   err,
	})
}
