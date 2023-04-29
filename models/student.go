package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Type      string         `json:"type"`
	SchoolID  uuid.UUID      `json:"school_id" gorm:"type:uuid"`
	School    *School        `json:"school" gorm:"foreignKey:SchoolID"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
