package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/services"
	"github.com/learng/backend/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
// POST /api/v1/auth/register
func (h *AuthHandler) Register(c echo.Context) error {
	var req services.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Register user
	response, err := h.authService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, response)
}

// Login handles user login
// POST /api/v1/auth/login
func (h *AuthHandler) Login(c echo.Context) error {
	var req services.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	// Login user
	response, err := h.authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, response)
}

// GetMe returns the currently authenticated user
// GET /api/v1/auth/me
func (h *AuthHandler) GetMe(c echo.Context) error {
	// Get user ID from context (set by auth middleware)
	userID, err := utils.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
	}

	// Get user from database
	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("User not found"))
	}

	return c.JSON(http.StatusOK, user)
}
