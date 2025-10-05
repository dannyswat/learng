package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LearnerProgress struct {
	ID           string         `gorm:"primaryKey" json:"id"`
	UserID       string         `gorm:"not null;index" json:"userId"`
	WordID       string         `gorm:"not null;index" json:"wordId"`
	MasteryLevel string         `gorm:"default:new" json:"masteryLevel"` // 'new' | 'learning' | 'review' | 'mastered'
	ViewCount    int            `gorm:"default:0" json:"viewCount"`
	LastViewedAt *time.Time     `json:"lastViewedAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	User User `gorm:"foreignKey:UserID" json:"-"`
	Word Word `gorm:"foreignKey:WordID" json:"-"`
}

func (lp *LearnerProgress) BeforeCreate(tx *gorm.DB) error {
	if lp.ID == "" {
		lp.ID = uuid.New().String()
	}
	if lp.MasteryLevel == "" {
		lp.MasteryLevel = "new"
	}
	return nil
}

func (LearnerProgress) TableName() string {
	return "learner_progress"
}

// Ensure unique constraint on (user_id, word_id)
type LearnerProgressIndex struct {
	UserID string `gorm:"uniqueIndex:idx_user_word"`
	WordID string `gorm:"uniqueIndex:idx_user_word"`
}
