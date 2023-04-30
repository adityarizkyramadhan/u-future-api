package repository

import (
	"u-future-api/models"

	"gorm.io/gorm"
)

type School struct {
	db *gorm.DB
}

func New(db *gorm.DB) *School {
	return &School{db}
}

func (rs *School) Create(arg *models.School) error {
	return rs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (rs *School) FindById(id string) (*models.School, error) {
	var school models.School
	if err := rs.db.Preload("Students").Where("id = ?", id).Take(&school).Error; err != nil {
		return nil, err
	}
	return &school, nil
}

func (rs *School) FindAll() ([]models.School, error) {
	var schools []models.School
	if err := rs.db.Preload("Students").Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}
