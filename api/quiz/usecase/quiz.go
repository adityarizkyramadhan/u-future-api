package usecase

import (
	"errors"
	"u-future-api/api/quiz/repository"
	"u-future-api/models"

	"gorm.io/gorm"
)

type Quiz struct {
	uc *repository.Quiz
}

func New(uc *repository.Quiz) *Quiz {
	return &Quiz{uc}
}

func (uq *Quiz) GenerateFaker() {

}

func (uq *Quiz) FindByName(name string) (*models.Quiz, error) {
	return uq.uc.GetByName(name)
}

func (uq *Quiz) SearchTestUser(id string) (bool, error) {
	result, err := uq.uc.SearchByUserID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil || result == nil {
		return false, err
	}
	return true, nil
}
