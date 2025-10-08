package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/services"
	"github.com/learng/backend/internal/utils"
)

type ScenarioHandler struct {
	scenarioService services.ScenarioService
}

func NewScenarioHandler(scenarioService services.ScenarioService) *ScenarioHandler {
	return &ScenarioHandler{scenarioService: scenarioService}
}

// CreateScenario handles POST /api/v1/scenarios
func (h *ScenarioHandler) CreateScenario(c echo.Context) error {
	var req struct {
		JourneyID    string `json:"journeyId"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		DisplayOrder int    `json:"displayOrder"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	scenario := &models.Scenario{
		JourneyID:    req.JourneyID,
		Title:        req.Title,
		Description:  req.Description,
		DisplayOrder: req.DisplayOrder,
	}

	if err := h.scenarioService.CreateScenario(scenario); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse(scenario))
}

// GetScenarioByID handles GET /api/v1/scenarios/:id
func (h *ScenarioHandler) GetScenarioByID(c echo.Context) error {
	id := c.Param("id")

	scenario, err := h.scenarioService.GetScenarioWithWords(id)
	if err != nil {
		if err.Error() == "scenario not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Scenario not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch scenario"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(scenario))
}

// UpdateScenario handles PUT /api/v1/scenarios/:id
func (h *ScenarioHandler) UpdateScenario(c echo.Context) error {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	updatedScenario, err := h.scenarioService.UpdateScenario(id, updates)
	if err != nil {
		if err.Error() == "scenario not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Scenario not found"))
		}
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(updatedScenario))
}

// DeleteScenario handles DELETE /api/v1/scenarios/:id
func (h *ScenarioHandler) DeleteScenario(c echo.Context) error {
	id := c.Param("id")

	if err := h.scenarioService.DeleteScenario(id); err != nil {
		if err.Error() == "scenario not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Scenario not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete scenario"))
	}

	return c.NoContent(http.StatusNoContent)
}
