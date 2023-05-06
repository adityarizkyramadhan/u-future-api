package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Title     string         `json:"title"`
	Questions []Question     `gorm:"foreignKey:QuizID" json:"questions"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Question struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Text      string         `json:"text"`
	QuizID    uint           `json:"-"`
	Options   []Option       `gorm:"foreignKey:QuestionID" json:"options"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Option struct {
	ID         uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Text       string         `json:"text"`
	QuestionID uint           `json:"-"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type QuizResult struct {
	ID                 uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	QuizID             uuid.UUID `gorm:"type:uuid" json:"quiz_id"`
	UserID             uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ResultSectionOne   string    `json:"result_section_one"`
	ResultSectionTwo   string    `json:"result_section_two"`
	ResultSectionThree string    `json:"result_section_three"`
}
