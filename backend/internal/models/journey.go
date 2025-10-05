package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Journey struct {
	ID             string         `gorm:"primaryKey" json:"id"`
	Title          string         `gorm:"not null" json:"title"`
	Description    string         `json:"description"`
	SourceLanguage string         `gorm:"not null" json:"sourceLanguage"`       // ISO 639-1 code
	TargetLanguage string         `gorm:"not null" json:"targetLanguage"`       // ISO 639-1 code
	Status         string         `gorm:"not null;default:draft" json:"status"` // 'draft' | 'published' | 'archived'
	CreatedBy      string         `gorm:"not null" json:"createdBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Creator   User       `gorm:"foreignKey:CreatedBy" json:"-"`
	Scenarios []Scenario `gorm:"foreignKey:JourneyID" json:"scenarios,omitempty"`
}

func (j *Journey) BeforeCreate(tx *gorm.DB) error {
	if j.ID == "" {
		j.ID = uuid.New().String()
	}
	if j.Status == "" {
		j.Status = "draft"
	}
	return nil
}

func (Journey) TableName() string {
	return "journeys"
}
