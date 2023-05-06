package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type School struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string         `json:"name"`
	Type      string         `json:"type" gorm:"default:not-subscribe"`
	Students  []Student      `json:"students" gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type SchoolInput struct {
	Name string `json:"name" binding:"required"`
}
