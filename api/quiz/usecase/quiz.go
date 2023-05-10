package usecase

import (
	"errors"
	"sort"
	"u-future-api/api/quiz/repository"
	"u-future-api/models"
	"u-future-api/util/exception"
	"u-future-api/util/logic"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	uc *repository.Quiz
}

func New(uc *repository.Quiz) *Quiz {
	return &Quiz{uc}
}

func (uq *Quiz) FindByName(name string) (*models.Quiz, error) {
	if name == "" {
		return nil, exception.ErrNoQuery
	}
	return uq.uc.GetByName(name)
}

func (uq *Quiz) SectionTwoQuiz(id string) (*models.Quiz, error) {
	result, err := uq.uc.SearchByUserID(id)
	if err != nil {
		return nil, err
	}
	quiz, err := uq.uc.GetQuestion(result.ResultSectionOne)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}

func (uq *Quiz) SectionThreeQuiz(id string) (*models.Quiz, error) {
	result, err := uq.uc.SearchByUserID(id)
	if err != nil {
		return nil, err
	}
	resultSectionTwo := logic.TrimAndChangeStringToArray(result.ResultSectionTwo)
	for i := range resultSectionTwo {
		quiz, err := uq.uc.GetTheme(resultSectionTwo[i])
		if err != nil {
			err = nil
			continue
		}
		if quiz != nil {
			return quiz, nil
		}
	}
	return nil, exception.ErrNoQuery

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

func (uq *Quiz) UpdateResult(query string, data interface{}, userId string) error {
	idUser, err := uuid.FromString(userId)
	if err != nil {
		return err
	}
	if query == "SectionOne" {
		arr := data.([]models.InputQuizString)
		stringArr := make([]string, 0, len(arr))
		for _, item := range arr {
			stringArr = append(stringArr, item.Data)
		}
		result := logic.MostFrequentElements(stringArr)

		resultUser := &models.QuizResult{
			ID:               uuid.Must(uuid.NewV4()),
			UserID:           idUser,
			ResultSectionOne: result,
		}
		return uq.uc.CreateResult(resultUser)
	} else if query == "SectionTwo" {
		arr := data.([]models.InputQuizInteger)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Data > arr[j].Data
		})
		questionId := arr[0].QuestionId
		result, err := uq.uc.GetQuestionById(questionId)
		if err != nil {
			return err
		}
		resultUser := &models.QuizResult{
			UserID:           idUser,
			ResultSectionTwo: result.Options[0].Text,
		}
		return uq.uc.UpdateResult(resultUser, "two")
	} else if query == "SectionThree" {
		arr := data.([]models.InputQuizInteger)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Data > arr[j].Data
		})
		questionId := arr[0].QuestionId
		result, err := uq.uc.GetQuestionById(questionId)
		if err != nil {
			return err
		}
		resultUser := &models.QuizResult{
			UserID:             idUser,
			ResultSectionThree: result.Options[0].Text,
		}
		return uq.uc.UpdateResult(resultUser, "three")
	}
	return exception.ErrNoQuery
}
