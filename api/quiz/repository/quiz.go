package repository

import (
	"u-future-api/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Quiz {
	return &Quiz{db}
}

func (rq *Quiz) Create(arg *models.Quiz) error {
	return rq.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (q *Quiz) Get(id uuid.UUID) (*models.Quiz, error) {
	var quiz models.Quiz
	err := q.db.
		Preload("Options").
		Preload("Questions").First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (q *Quiz) GetByName(name string) (*models.Quiz, error) {
	var quiz models.Quiz
	err := q.db.
		Preload("Options").
		Preload("Questions").
		Where("title = ?", name).
		First(&quiz).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}
