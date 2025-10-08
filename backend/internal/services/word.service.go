package services

import (
	"errors"

	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/repository"
	"gorm.io/gorm"
)

type WordService interface {
	CreateWord(word *models.Word) error
	GetWordByID(id string) (*models.Word, error)
	GetWordsByScenarioID(scenarioID string) ([]models.Word, error)
	UpdateWord(id string, updates map[string]interface{}) (*models.Word, error)
	DeleteWord(id string) error
}

type wordService struct {
	wordRepo     repository.WordRepository
	scenarioRepo repository.ScenarioRepository
}

func NewWordService(wordRepo repository.WordRepository, scenarioRepo repository.ScenarioRepository) WordService {
	return &wordService{
		wordRepo:     wordRepo,
		scenarioRepo: scenarioRepo,
	}
}

func (s *wordService) CreateWord(word *models.Word) error {
	// Validate required fields
	if word.TargetText == "" {
		return errors.New("target text is required")
	}
	if word.ScenarioID == "" {
		return errors.New("scenario ID is required")
	}

	// Verify scenario exists
	_, err := s.scenarioRepo.GetByID(word.ScenarioID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("scenario not found")
		}
		return err
	}

	return s.wordRepo.Create(word)
}

func (s *wordService) GetWordByID(id string) (*models.Word, error) {
	return s.wordRepo.GetByID(id)
}

func (s *wordService) GetWordsByScenarioID(scenarioID string) ([]models.Word, error) {
	return s.wordRepo.GetByScenarioID(scenarioID)
}

func (s *wordService) UpdateWord(id string, updates map[string]interface{}) (*models.Word, error) {
	word, err := s.wordRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("word not found")
		}
		return nil, err
	}

	// Apply updates
	if targetText, ok := updates["targetText"].(string); ok {
		word.TargetText = targetText
	}
	if sourceText, ok := updates["sourceText"].(string); ok {
		word.SourceText = sourceText
	}
	if displayOrder, ok := updates["displayOrder"].(float64); ok {
		word.DisplayOrder = int(displayOrder)
	}
	if imageURL, ok := updates["imageUrl"].(string); ok {
		word.ImageURL = &imageURL
	}
	if audioURL, ok := updates["audioUrl"].(string); ok {
		word.AudioURL = &audioURL
	}
	if generationMethod, ok := updates["generationMethod"].(string); ok {
		word.GenerationMethod = generationMethod
	}

	if err := s.wordRepo.Update(word); err != nil {
		return nil, err
	}

	return word, nil
}

func (s *wordService) DeleteWord(id string) error {
	// Check if word exists
	_, err := s.wordRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("word not found")
		}
		return err
	}

	return s.wordRepo.Delete(id)
}
