package repository

import (
	"github.com/learng/backend/internal/models"
	"gorm.io/gorm"
)

type ScenarioRepository interface {
	Create(scenario *models.Scenario) error
	GetByID(id string) (*models.Scenario, error)
	GetByJourneyID(journeyID string) ([]models.Scenario, error)
	Update(scenario *models.Scenario) error
	Delete(id string) error
	GetByIDWithWords(id string) (*models.Scenario, error)
}

type scenarioRepository struct {
	db *gorm.DB
}

func NewScenarioRepository(db *gorm.DB) ScenarioRepository {
	return &scenarioRepository{db: db}
}

func (r *scenarioRepository) Create(scenario *models.Scenario) error {
	return r.db.Create(scenario).Error
}

func (r *scenarioRepository) GetByID(id string) (*models.Scenario, error) {
	var scenario models.Scenario
	if err := r.db.First(&scenario, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &scenario, nil
}

func (r *scenarioRepository) GetByJourneyID(journeyID string) ([]models.Scenario, error) {
	var scenarios []models.Scenario
	if err := r.db.Where("journey_id = ?", journeyID).Order("display_order ASC").Find(&scenarios).Error; err != nil {
		return nil, err
	}
	return scenarios, nil
}

func (r *scenarioRepository) Update(scenario *models.Scenario) error {
	return r.db.Save(scenario).Error
}

func (r *scenarioRepository) Delete(id string) error {
	return r.db.Delete(&models.Scenario{}, "id = ?", id).Error
}

func (r *scenarioRepository) GetByIDWithWords(id string) (*models.Scenario, error) {
	var scenario models.Scenario
	if err := r.db.Preload("Words", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).First(&scenario, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &scenario, nil
}
