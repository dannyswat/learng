package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Word struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	ScenarioID       string         `gorm:"not null;index" json:"scenarioId"`
	TargetText       string         `gorm:"not null" json:"targetText"`
	SourceText       string         `json:"sourceText"`
	DisplayOrder     int            `gorm:"not null" json:"displayOrder"`
	ImageURL         *string        `json:"imageUrl"`
	AudioURL         *string        `json:"audioUrl"`
	GenerationMethod string         `gorm:"default:manual" json:"generationMethod"` // 'manual' | 'ai_image' | 'ai_audio' | 'ai_both'
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Scenario Scenario `gorm:"foreignKey:ScenarioID" json:"-"`
}

func (w *Word) BeforeCreate(tx *gorm.DB) error {
	if w.ID == "" {
		w.ID = uuid.New().String()
	}
	if w.GenerationMethod == "" {
		w.GenerationMethod = "manual"
	}
	return nil
}

func (Word) TableName() string {
	return "words"
}
