package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/utils"
)

// AuthMiddleware validates JWT tokens
func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Missing authorization header",
				})
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid authorization header format",
				})
			}

			claims, err := utils.ValidateToken(tokenString, jwtSecret)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid or expired token",
				})
			}

			// Store user info in context
			c.Set("userId", claims.UserID)
			c.Set("userEmail", claims.Email)
			c.Set("userRole", claims.Role)

			return next(c)
		}
	}
}

// RequireRole ensures the user has a specific role
func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("userRole")
			if userRole == nil || userRole != role {
				return c.JSON(http.StatusForbidden, map[string]string{
					"error": "Insufficient permissions",
				})
			}
			return next(c)
		}
	}
}

// GetUserID retrieves the user ID from context
func GetUserID(c echo.Context) string {
	if userID := c.Get("userId"); userID != nil {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return ""
}

// GetUserRole retrieves the user role from context
func GetUserRole(c echo.Context) string {
	if role := c.Get("userRole"); role != nil {
		if r, ok := role.(string); ok {
			return r
		}
	}
	return ""
}
