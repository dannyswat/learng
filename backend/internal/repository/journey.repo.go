package repository

import (
	"github.com/learng/backend/internal/models"
	"gorm.io/gorm"
)

type JourneyRepository interface {
	Create(journey *models.Journey) error
	GetByID(id string) (*models.Journey, error)
	GetAll(filters map[string]interface{}, page, limit int) ([]models.Journey, int64, error)
	Update(journey *models.Journey) error
	Delete(id string) error
	GetByIDWithScenarios(id string) (*models.Journey, error)
	GetByCreator(creatorID string, status string) ([]models.Journey, error)
}

type journeyRepository struct {
	db *gorm.DB
}

func NewJourneyRepository(db *gorm.DB) JourneyRepository {
	return &journeyRepository{db: db}
}

func (r *journeyRepository) Create(journey *models.Journey) error {
	return r.db.Create(journey).Error
}

func (r *journeyRepository) GetByID(id string) (*models.Journey, error) {
	var journey models.Journey
	if err := r.db.First(&journey, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &journey, nil
}

func (r *journeyRepository) GetAll(filters map[string]interface{}, page, limit int) ([]models.Journey, int64, error) {
	var journeys []models.Journey
	var total int64

	query := r.db.Model(&models.Journey{})

	// Apply filters
	if status, ok := filters["status"].(string); ok && status != "" {
		query = query.Where("status = ?", status)
	}
	if createdBy, ok := filters["createdBy"].(string); ok && createdBy != "" {
		query = query.Where("created_by = ?", createdBy)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&journeys).Error; err != nil {
		return nil, 0, err
	}

	return journeys, total, nil
}

func (r *journeyRepository) Update(journey *models.Journey) error {
	return r.db.Save(journey).Error
}

func (r *journeyRepository) Delete(id string) error {
	return r.db.Delete(&models.Journey{}, "id = ?", id).Error
}

func (r *journeyRepository) GetByIDWithScenarios(id string) (*models.Journey, error) {
	var journey models.Journey
	if err := r.db.Preload("Scenarios", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).First(&journey, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &journey, nil
}

func (r *journeyRepository) GetByCreator(creatorID string, status string) ([]models.Journey, error) {
	var journeys []models.Journey
	query := r.db.Where("created_by = ?", creatorID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at DESC").Find(&journeys).Error; err != nil {
		return nil, err
	}
	return journeys, nil
}
