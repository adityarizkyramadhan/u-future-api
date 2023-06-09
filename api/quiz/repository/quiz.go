package repository

import (
	"fmt"
	"u-future-api/models"
	"u-future-api/util/exception"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Quiz {
	return &Quiz{db}
}

func (rq *Quiz) Create(arg any) error {
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
		Where("title LIKE ?", name).
		Take(&quiz).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (q *Quiz) GetTheme(name string) (*models.Quiz, error) {
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
	existingResult := &models.QuizResult{}
	if err := q.db.Where("user_id = ?", arg.UserID).First(existingResult).Error; err == nil {
		return q.db.Save(arg).Error
	}
	return q.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(arg).Error
	})
}

func (q *Quiz) UpdateResult(arg *models.QuizResult, quiz string) error {
	return q.db.Transaction(func(tx *gorm.DB) error {
		if quiz == "two" {
			return tx.Model(&models.QuizResult{}).Where("user_id = ?", arg.UserID).Update("result_section_two", arg.ResultSectionTwo).Error
		} else if quiz == "three" {
			return tx.Model(&models.QuizResult{}).Where("user_id = ?", arg.UserID).Update("result_section_three", arg.ResultSectionThree).Error
		}
		return exception.ErrNoQuery
	})
}

func (q *Quiz) UpdateHistory(value string, userId string, quiz string) error {
	return q.db.Transaction(func(tx *gorm.DB) error {
		if quiz == "two" {
			return tx.Model(&models.QuizResult{}).Where("user_id = ?", userId).Update("result_section_two", value).Error
		} else if quiz == "three" {
			return tx.Model(&models.QuizResult{}).Where("user_id = ?", userId).Update("result_section_three", value).Error
		}
		return exception.ErrNoQuery
	})
}

func (q *Quiz) GetQuestionById(id string) (*models.Question, error) {
	var question models.Question
	if err := q.db.Model(&models.Question{}).Preload("Options").Where("id = ?", id).Take(&question).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (q *Quiz) GetHistoryByUserId(id string) (*models.QuizResultRiwayat, error) {
	var history models.QuizResultRiwayat
	if err := q.db.Model(&models.QuizResultRiwayat{}).Where("user_id = ?", id).Take(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}

func (q *Quiz) GetAllJurusan() ([]models.Jurusan, error) {
	var jurusans []models.Jurusan
	if err := q.db.Find(&jurusans).Error; err != nil {
		return nil, err
	}
	return jurusans, nil
}

func (q *Quiz) GetResultByUserID(userId string) (*models.QuizResult, error) {
	var result models.QuizResult
	if err := q.db.Where("user_id", userId).Take(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (q *Quiz) GetAnalisisRiasec(tag string) (*models.AnalisisRiasec, error) {
	var result models.AnalisisRiasec
	if err := q.db.Where("tag LIKE ?", tag).Take(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
