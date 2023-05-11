package repository

import (
	"u-future-api/models"

	"gorm.io/gorm"
)

type Jurusan struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Jurusan {
	return &Jurusan{db}
}

func (rs *Jurusan) Create(arg *models.JurusanStudentCompare) error {
	return rs.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (rs *Jurusan) Read(id uint) (*models.JurusanStudentCompare, error) {
	var result models.JurusanStudentCompare
	if err := rs.db.First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (rs *Jurusan) QueryByName(name string) ([]*models.JurusanStudentCompare, error) {
	var results []*models.JurusanStudentCompare
	if err := rs.db.Where("nama_jurusan = ?", name).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (rs *Jurusan) QueryByIDUser(idUser string) ([]*models.JurusanStudentCompare, error) {
	var results []*models.JurusanStudentCompare
	if err := rs.db.Where("user_id = ?", idUser).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (rs *Jurusan) FindAll() ([]*models.JurusanStudentCompare, error) {
	var results []*models.JurusanStudentCompare
	if err := rs.db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (rs *Jurusan) FindAllJurusan() ([]*models.Jurusan, error) {
	var results []*models.Jurusan
	if err := rs.db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
