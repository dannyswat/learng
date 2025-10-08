package repository

import (
	"github.com/learng/backend/internal/models"
	"gorm.io/gorm"
)

type WordRepository interface {
	Create(word *models.Word) error
	GetByID(id string) (*models.Word, error)
	GetByScenarioID(scenarioID string) ([]models.Word, error)
	Update(word *models.Word) error
	Delete(id string) error
}

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) WordRepository {
	return &wordRepository{db: db}
}

func (r *wordRepository) Create(word *models.Word) error {
	return r.db.Create(word).Error
}

func (r *wordRepository) GetByID(id string) (*models.Word, error) {
	var word models.Word
	if err := r.db.First(&word, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &word, nil
}

func (r *wordRepository) GetByScenarioID(scenarioID string) ([]models.Word, error) {
	var words []models.Word
	if err := r.db.Where("scenario_id = ?", scenarioID).Order("display_order ASC").Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (r *wordRepository) Update(word *models.Word) error {
	return r.db.Save(word).Error
}

func (r *wordRepository) Delete(id string) error {
	return r.db.Delete(&models.Word{}, "id = ?", id).Error
}
