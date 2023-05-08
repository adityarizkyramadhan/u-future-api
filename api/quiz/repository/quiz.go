package repository

import (
	"fmt"
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
		Preload("Questions").
		Preload("Options").
		First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (q *Quiz) GetByName(name string) (*models.Quiz, error) {
	var quiz models.Quiz
	err := q.db.
		Preload("Questions").
		Preload("Questions.Options").
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Take(&quiz).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (q *Quiz) GetQuestion(name string) (*models.Quiz, error) {
	var quiz models.Quiz
	err := q.db.
		Preload("Questions").
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Take(&quiz).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (q *Quiz) SearchByUserID(id string) (*models.QuizResult, error) {
	var result models.QuizResult
	if err := q.db.Model(&models.QuizResult{}).Where("user_id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (q *Quiz) CreateResult(arg *models.QuizResult) error {
	return q.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (q *Quiz) UpdateResult(arg *models.QuizResult) error {
	return q.db.Transaction(func(tx *gorm.DB) error {
		return tx.Where("user_id = ?", arg.UserID).Updates(arg).Error
	})
}
