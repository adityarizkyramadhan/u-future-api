package usecase

import (
	"u-future-api/api/quiz/repository"
	"u-future-api/models"
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
