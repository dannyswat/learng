package services

import (
	"errors"

	"github.com/learng/backend/internal/models"
	"github.com/learng/backend/internal/repository"
	"gorm.io/gorm"
)

type ScenarioService interface {
	CreateScenario(scenario *models.Scenario) error
	GetScenarioByID(id string) (*models.Scenario, error)
	GetScenariosByJourneyID(journeyID string) ([]models.Scenario, error)
	UpdateScenario(id string, updates map[string]interface{}) (*models.Scenario, error)
	DeleteScenario(id string) error
	GetScenarioWithWords(id string) (*models.Scenario, error)
}

type scenarioService struct {
	scenarioRepo repository.ScenarioRepository
	journeyRepo  repository.JourneyRepository
}

func NewScenarioService(scenarioRepo repository.ScenarioRepository, journeyRepo repository.JourneyRepository) ScenarioService {
	return &scenarioService{
		scenarioRepo: scenarioRepo,
		journeyRepo:  journeyRepo,
	}
}

func (s *scenarioService) CreateScenario(scenario *models.Scenario) error {
	// Validate required fields
	if scenario.Title == "" {
		return errors.New("title is required")
	}
	if scenario.JourneyID == "" {
		return errors.New("journey ID is required")
	}

	// Verify journey exists
	_, err := s.journeyRepo.GetByID(scenario.JourneyID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("journey not found")
		}
		return err
	}

	return s.scenarioRepo.Create(scenario)
}

func (s *scenarioService) GetScenarioByID(id string) (*models.Scenario, error) {
	return s.scenarioRepo.GetByID(id)
}

func (s *scenarioService) GetScenariosByJourneyID(journeyID string) ([]models.Scenario, error) {
	return s.scenarioRepo.GetByJourneyID(journeyID)
}

func (s *scenarioService) UpdateScenario(id string, updates map[string]interface{}) (*models.Scenario, error) {
	scenario, err := s.scenarioRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("scenario not found")
		}
		return nil, err
	}

	// Apply updates
	if title, ok := updates["title"].(string); ok {
		scenario.Title = title
	}
	if description, ok := updates["description"].(string); ok {
		scenario.Description = description
	}
	if displayOrder, ok := updates["displayOrder"].(float64); ok {
		scenario.DisplayOrder = int(displayOrder)
	}

	if err := s.scenarioRepo.Update(scenario); err != nil {
		return nil, err
	}

	return scenario, nil
}

func (s *scenarioService) DeleteScenario(id string) error {
	// Check if scenario exists
	_, err := s.scenarioRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("scenario not found")
		}
		return err
	}

	return s.scenarioRepo.Delete(id)
}

func (s *scenarioService) GetScenarioWithWords(id string) (*models.Scenario, error) {
	scenario, err := s.scenarioRepo.GetByIDWithWords(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("scenario not found")
		}
		return nil, err
	}
	return scenario, nil
}
