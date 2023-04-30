package usecase

import (
	"u-future-api/api/school/repository"
	"u-future-api/models"

	"github.com/gofrs/uuid"
)

type School struct {
	rs *repository.School
}

func New(rs *repository.School) *School {
	return &School{rs}
}

func (us *School) Create(arg *models.SchoolInput) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	school := &models.School{
		ID:   id,
		Name: arg.Name,
	}
	return us.rs.Create(school)
}

func (us *School) FindAll() ([]models.School, error) {
	return us.rs.FindAll()
}

func (us *School) FindById(id string) (*models.School, error) {
	return us.rs.FindById(id)
}

func (us *School) GenerateFaker() error {
	schools := []models.SchoolInput{
		{
			Name: "SMAN 1 Malang",
		},
		{
			Name: "SMAN 1 Sidoarjo",
		},
		{
			Name: "SMAN 1 Surabaya",
		},
		{
			Name: "SMAN 5 Surabaya",
		},
		{
			Name: "SMAN 3 Sidoarjo",
		},
		{
			Name: "SMAN 1 Jakarta",
		},
	}
	for _, v := range schools {
		err := us.Create(&v)
		if err != nil {
			return err
		}
	}
	return nil
}
