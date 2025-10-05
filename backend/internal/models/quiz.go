package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	ScenarioID    string         `gorm:"not null;index" json:"scenarioId"`
	Title         string         `gorm:"not null" json:"title"`
	PassThreshold float64        `gorm:"default:70.0" json:"passThreshold"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Scenario  Scenario       `gorm:"foreignKey:ScenarioID" json:"-"`
	Questions []QuizQuestion `gorm:"foreignKey:QuizID" json:"questions,omitempty"`
}

func (q *Quiz) BeforeCreate(tx *gorm.DB) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	if q.PassThreshold == 0 {
		q.PassThreshold = 70.0
	}
	return nil
}

func (Quiz) TableName() string {
	return "quizzes"
}

type QuizQuestion struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	QuizID        string         `gorm:"not null;index" json:"quizId"`
	WordID        string         `gorm:"not null" json:"wordId"`
	QuestionType  string         `gorm:"not null" json:"questionType"` // 'multiple_choice' | 'audio_match' | 'image_match'
	QuestionText  string         `json:"questionText"`
	CorrectAnswer string         `gorm:"not null" json:"correctAnswer"`
	Options       string         `json:"options"` // JSON array of choices
	DisplayOrder  int            `gorm:"not null" json:"displayOrder"`
	CreatedAt     time.Time      `json:"createdAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Quiz Quiz `gorm:"foreignKey:QuizID" json:"-"`
	Word Word `gorm:"foreignKey:WordID" json:"-"`
}

func (qq *QuizQuestion) BeforeCreate(tx *gorm.DB) error {
	if qq.ID == "" {
		qq.ID = uuid.New().String()
	}
	return nil
}

func (QuizQuestion) TableName() string {
	return "quiz_questions"
}
