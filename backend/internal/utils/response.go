package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// SuccessResponse creates a standardized success response
func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"data":    data,
	}
}

// ErrorResponse creates a standardized error response
func ErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"error": message,
	}
}

// ValidationErrorResponse creates a response for validation errors
func ValidationErrorResponse(errors map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"error":  "Validation failed",
		"errors": errors,
	}
}

// GetUserID retrieves the user ID from Echo context
func GetUserID(c echo.Context) (string, error) {
	userID := c.Get("userId")
	if userID == nil {
		return "", errors.New("user ID not found in context")
	}

	id, ok := userID.(string)
	if !ok {
		return "", errors.New("invalid user ID type")
	}

	return id, nil
}

// GetUserRole retrieves the user role from Echo context
func GetUserRole(c echo.Context) (string, error) {
	role := c.Get("userRole")
	if role == nil {
		return "", errors.New("user role not found in context")
	}

	r, ok := role.(string)
	if !ok {
		return "", errors.New("invalid user role type")
	}

	return r, nil
}
