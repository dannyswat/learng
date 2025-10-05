package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizAttempt struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	UserID         string    `gorm:"not null;index" json:"userId"`
	QuizID         string    `gorm:"not null;index" json:"quizId"`
	Score          float64   `gorm:"not null" json:"score"` // Percentage (0-100)
	TotalQuestions int       `gorm:"not null" json:"totalQuestions"`
	CorrectAnswers int       `gorm:"not null" json:"correctAnswers"`
	Answers        string    `gorm:"not null" json:"answers"` // JSON array of {question_id, answer, is_correct}
	CompletedAt    time.Time `json:"completedAt"`

	// Associations
	User User `gorm:"foreignKey:UserID" json:"-"`
	Quiz Quiz `gorm:"foreignKey:QuizID" json:"-"`
}

func (qa *QuizAttempt) BeforeCreate(tx *gorm.DB) error {
	if qa.ID == "" {
		qa.ID = uuid.New().String()
	}
	if qa.CompletedAt.IsZero() {
		qa.CompletedAt = time.Now()
	}
	return nil
}

func (QuizAttempt) TableName() string {
	return "quiz_attempts"
}
