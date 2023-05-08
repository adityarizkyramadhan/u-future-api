package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Title     string         `json:"title"`
	Questions []Question     `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"questions"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Question struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Text      string         `json:"text"`
	QuizID    uuid.UUID      `json:"-" gorm:"type:char(36);"`
	Quiz      Quiz           `gorm:"foreignKey:QuizID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Options   []Option       `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"options"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Option struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Text        string         `json:"text"`
	QuestionID  uuid.UUID      `json:"-" gorm:"type:char(36);"`
	Question    Question       `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type QuizResult struct {
	ID                 uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID             uuid.UUID `gorm:"type:uuid" json:"user_id"`
	ResultSectionOne   string    `json:"result_section_one"`
	ResultSectionTwo   string    `json:"result_section_two"`
	ResultSectionThree string    `json:"result_section_three"`
}

type InputQuizString struct {
	Data string `json:"data" binding:"required"`
}

type InputQuizInteger struct {
	Data int `json:"data" binding:"required"`
}
