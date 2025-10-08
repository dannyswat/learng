package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/services"
	"github.com/learng/backend/internal/utils"
)

type JourneyHandler struct {
	journeyService services.JourneyService
}

func NewJourneyHandler(journeyService services.JourneyService) *JourneyHandler {
	return &JourneyHandler{journeyService: journeyService}
}

// CreateJourney handles POST /api/v1/journeys
func (h *JourneyHandler) CreateJourney(c echo.Context) error {
	userID := c.Get("userId").(string)

	var req struct {
		Title          string `json:"title"`
		Description    string `json:"description"`
		SourceLanguage string `json:"sourceLanguage"`
		TargetLanguage string `json:"targetLanguage"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	journey := &models.Journey{
		Title:          req.Title,
		Description:    req.Description,
		SourceLanguage: req.SourceLanguage,
		TargetLanguage: req.TargetLanguage,
		CreatedBy:      userID,
		Status:         "draft",
	}

	if err := h.journeyService.CreateJourney(journey); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse(journey))
}

// GetJourneys handles GET /api/v1/journeys
func (h *JourneyHandler) GetJourneys(c echo.Context) error {
	// Parse query parameters
	status := c.QueryParam("status")
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 20

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	filters := make(map[string]interface{})
	if status != "" {
		filters["status"] = status
	}

	journeys, total, err := h.journeyService.GetAllJourneys(filters, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch journeys"))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"journeys": journeys,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

// GetJourneyByID handles GET /api/v1/journeys/:id
func (h *JourneyHandler) GetJourneyByID(c echo.Context) error {
	id := c.Param("id")

	journey, err := h.journeyService.GetJourneyWithScenarios(id)
	if err != nil {
		if err.Error() == "journey not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Journey not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch journey"))
	}

	// Add counts
	scenarioCount := len(journey.Scenarios)
	wordCount := 0
	for _, scenario := range journey.Scenarios {
		wordCount += len(scenario.Words)
	}

	response := map[string]interface{}{
		"id":             journey.ID,
		"title":          journey.Title,
		"description":    journey.Description,
		"sourceLanguage": journey.SourceLanguage,
		"targetLanguage": journey.TargetLanguage,
		"status":         journey.Status,
		"createdBy":      journey.CreatedBy,
		"createdAt":      journey.CreatedAt,
		"updatedAt":      journey.UpdatedAt,
		"scenarios":      journey.Scenarios,
		"scenarioCount":  scenarioCount,
		"wordCount":      wordCount,
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(response))
}

// UpdateJourney handles PUT /api/v1/journeys/:id
func (h *JourneyHandler) UpdateJourney(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("userId").(string)

	// First check if journey exists and user owns it
	journey, err := h.journeyService.GetJourneyByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Journey not found"))
	}

	if journey.CreatedBy != userID {
		return c.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to update this journey"))
	}

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	updatedJourney, err := h.journeyService.UpdateJourney(id, updates)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(updatedJourney))
}

// DeleteJourney handles DELETE /api/v1/journeys/:id
func (h *JourneyHandler) DeleteJourney(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("userId").(string)

	// First check if journey exists and user owns it
	journey, err := h.journeyService.GetJourneyByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Journey not found"))
	}

	if journey.CreatedBy != userID {
		return c.JSON(http.StatusForbidden, utils.ErrorResponse("You don't have permission to delete this journey"))
	}

	if err := h.journeyService.DeleteJourney(id); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete journey"))
	}

	return c.NoContent(http.StatusNoContent)
}
