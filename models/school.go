package models

import "github.com/google/uuid"

type School struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name"`
	Students []Student `json:"students" gorm:"foreignKey:SchoolID"`
}
