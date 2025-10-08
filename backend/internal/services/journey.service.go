package services

import (
	"errors"

	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/repository"
	"gorm.io/gorm"
)

type JourneyService interface {
	CreateJourney(journey *models.Journey) error
	GetJourneyByID(id string) (*models.Journey, error)
	GetAllJourneys(filters map[string]interface{}, page, limit int) ([]models.Journey, int64, error)
	UpdateJourney(id string, updates map[string]interface{}) (*models.Journey, error)
	DeleteJourney(id string) error
	GetJourneyWithScenarios(id string) (*models.Journey, error)
}

type journeyService struct {
	journeyRepo  repository.JourneyRepository
	scenarioRepo repository.ScenarioRepository
}

func NewJourneyService(journeyRepo repository.JourneyRepository, scenarioRepo repository.ScenarioRepository) JourneyService {
	return &journeyService{
		journeyRepo:  journeyRepo,
		scenarioRepo: scenarioRepo,
	}
}

func (s *journeyService) CreateJourney(journey *models.Journey) error {
	// Validate required fields
	if journey.Title == "" {
		return errors.New("title is required")
	}
	if journey.SourceLanguage == "" {
		return errors.New("source language is required")
	}
	if journey.TargetLanguage == "" {
		return errors.New("target language is required")
	}
	if journey.CreatedBy == "" {
		return errors.New("created by is required")
	}

	return s.journeyRepo.Create(journey)
}

func (s *journeyService) GetJourneyByID(id string) (*models.Journey, error) {
	return s.journeyRepo.GetByID(id)
}

func (s *journeyService) GetAllJourneys(filters map[string]interface{}, page, limit int) ([]models.Journey, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}
	return s.journeyRepo.GetAll(filters, page, limit)
}

func (s *journeyService) UpdateJourney(id string, updates map[string]interface{}) (*models.Journey, error) {
	journey, err := s.journeyRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("journey not found")
		}
		return nil, err
	}

	// Apply updates
	if title, ok := updates["title"].(string); ok {
		journey.Title = title
	}
	if description, ok := updates["description"].(string); ok {
		journey.Description = description
	}
	if status, ok := updates["status"].(string); ok {
		// Validate status
		if status != "draft" && status != "published" && status != "archived" {
			return nil, errors.New("invalid status")
		}
		journey.Status = status
	}
	if sourceLang, ok := updates["sourceLanguage"].(string); ok {
		journey.SourceLanguage = sourceLang
	}
	if targetLang, ok := updates["targetLanguage"].(string); ok {
		journey.TargetLanguage = targetLang
	}

	if err := s.journeyRepo.Update(journey); err != nil {
		return nil, err
	}

	return journey, nil
}

func (s *journeyService) DeleteJourney(id string) error {
	// Check if journey exists
	_, err := s.journeyRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("journey not found")
		}
		return err
	}

	return s.journeyRepo.Delete(id)
}

func (s *journeyService) GetJourneyWithScenarios(id string) (*models.Journey, error) {
	journey, err := s.journeyRepo.GetByIDWithScenarios(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("journey not found")
		}
		return nil, err
	}

	// Add word count to each scenario
	for i := range journey.Scenarios {
		words, err := s.scenarioRepo.GetByIDWithWords(journey.Scenarios[i].ID)
		if err == nil && words != nil {
			journey.Scenarios[i].Words = words.Words
		}
	}

	return journey, nil
}
