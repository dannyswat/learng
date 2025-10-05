package utils

import "github.com/labstack/echo/v4"

// SuccessResponse creates a standardized success response
func SuccessResponse(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, map[string]interface{}{
		"success": true,
		"data":    data,
	})
}

// ErrorResponse creates a standardized error response
func ErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]interface{}{
		"success": false,
		"error":   message,
	})
}

// ValidationErrorResponse creates a response for validation errors
func ValidationErrorResponse(c echo.Context, errors map[string]string) error {
	return c.JSON(400, map[string]interface{}{
		"success": false,
		"error":   "Validation failed",
		"errors":  errors,
	})
}
