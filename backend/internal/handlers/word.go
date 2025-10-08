package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/services"
	"github.com/learng/backend/internal/utils"
)

type WordHandler struct {
	wordService services.WordService
}

func NewWordHandler(wordService services.WordService) *WordHandler {
	return &WordHandler{wordService: wordService}
}

// CreateWord handles POST /api/v1/words
func (h *WordHandler) CreateWord(c echo.Context) error {
	var req struct {
		ScenarioID       string  `json:"scenarioId"`
		TargetText       string  `json:"targetText"`
		SourceText       string  `json:"sourceText"`
		DisplayOrder     int     `json:"displayOrder"`
		ImageURL         *string `json:"imageUrl"`
		AudioURL         *string `json:"audioUrl"`
		GenerationMethod string  `json:"generationMethod"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	word := &models.Word{
		ScenarioID:       req.ScenarioID,
		TargetText:       req.TargetText,
		SourceText:       req.SourceText,
		DisplayOrder:     req.DisplayOrder,
		ImageURL:         req.ImageURL,
		AudioURL:         req.AudioURL,
		GenerationMethod: req.GenerationMethod,
	}

	if err := h.wordService.CreateWord(word); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse(word))
}

// GetWordByID handles GET /api/v1/words/:id
func (h *WordHandler) GetWordByID(c echo.Context) error {
	id := c.Param("id")

	word, err := h.wordService.GetWordByID(id)
	if err != nil {
		if err.Error() == "word not found" || err.Error() == "record not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Word not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to fetch word"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(word))
}

// UpdateWord handles PUT /api/v1/words/:id
func (h *WordHandler) UpdateWord(c echo.Context) error {
	id := c.Param("id")

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	updatedWord, err := h.wordService.UpdateWord(id, updates)
	if err != nil {
		if err.Error() == "word not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Word not found"))
		}
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(updatedWord))
}

// DeleteWord handles DELETE /api/v1/words/:id
func (h *WordHandler) DeleteWord(c echo.Context) error {
	id := c.Param("id")

	if err := h.wordService.DeleteWord(id); err != nil {
		if err.Error() == "word not found" {
			return c.JSON(http.StatusNotFound, utils.ErrorResponse("Word not found"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to delete word"))
	}

	return c.NoContent(http.StatusNoContent)
}
