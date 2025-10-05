package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Scenario struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	JourneyID    string         `gorm:"not null;index" json:"journeyId"`
	Title        string         `gorm:"not null" json:"title"`
	Description  string         `json:"description"`
	DisplayOrder int            `gorm:"not null" json:"displayOrder"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Journey Journey `gorm:"foreignKey:JourneyID" json:"-"`
	Words   []Word  `gorm:"foreignKey:ScenarioID" json:"words,omitempty"`
	Quizzes []Quiz  `gorm:"foreignKey:ScenarioID" json:"quizzes,omitempty"`
}

func (s *Scenario) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

func (Scenario) TableName() string {
	return "scenarios"
}
