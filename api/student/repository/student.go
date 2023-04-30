package repository

import (
	"u-future-api/models"

	"gorm.io/gorm"
)

type Student struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Student {
	return &Student{db}
}

func (rs *Student) Create(arg *models.Student) error {
	return rs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (rs *Student) Update(arg *models.Student) error {
	return rs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(arg).Error
	})
}

func (rs *Student) FindByEmail(email string) (*models.Student, error) {
	var student models.Student
	if err := rs.db.Where("email = ?", email).Take(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (rs *Student) FindById(id string) (*models.Student, error) {
	var student models.Student
	if err := rs.db.Preload("School").Where("id = ?", id).Take(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (rs *Student) FindBySchoolId(id string) (*[]models.Student, error) {
	var student []models.Student
	if err := rs.db.Preload("School").Where("school_id = ?", id).Take(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}
