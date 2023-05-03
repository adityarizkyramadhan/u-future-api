package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Type      string         `json:"type" gorm:"default:personal"`
	SchoolID  uuid.UUID      `json:"school_id" gorm:"type:uuid"`
	School    *School        `json:"school" gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type StudentInput struct {
	Name     string `json:"name" faker:"name"`
	Email    string `json:"email" faker:"email"`
	Password string `json:"password" faker:"password"`
	Type     string `json:"type" faker:"-"`
}

type StudentRegister struct {
	SchoolID string `json:"school_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     string `json:"type"`
}

type StudentLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
