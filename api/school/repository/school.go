package repository

import (
	"fmt"
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

func (rs *School) FindPagination(limit, page int, name string) ([]models.School, error) {
	var schools []models.School
	offset := limit * (page - 1)
	db := rs.db.Limit(limit).Offset(offset)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	if err := db.Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}

func (rs *School) FindAll() ([]models.School, error) {
	var schools []models.School
	if err := rs.db.Preload("Students").Find(&schools).Error; err != nil {
		return nil, err
	}
	return schools, nil
}
