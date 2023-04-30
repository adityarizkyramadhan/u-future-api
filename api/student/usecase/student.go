package usecase

import (
	"u-future-api/api/school/usecase"
	"u-future-api/api/student/repository"
	"u-future-api/models"

	"github.com/go-faker/faker/v4"
	"github.com/gofrs/uuid"
)

type Student struct {
	rs *repository.Student
}

func New(rs *repository.Student) *Student {
	return &Student{rs}
}
func (us *Student) Create(arg *models.StudentInput, shoolId uuid.UUID) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	student := &models.Student{
		ID:       id,
		Name:     arg.Name,
		Email:    arg.Email,
		Password: arg.Password,
		Type:     arg.Type,
		SchoolID: shoolId,
	}
	return us.rs.Create(student)
}

func (us *Student) FindById(id string) (*models.Student, error) {
	return us.rs.FindById(id)
}

func (us *Student) GenerateFaker(ucSchool *usecase.School) error {
	schools, err := ucSchool.FindAll()
	if err != nil {
		return err
	}
	for _, v := range schools {
		for i := 0; i < 300; i++ {
			var studentInput models.StudentInput
			err = faker.FakeData(&studentInput)
			if err != nil {
				return err
			}
			studentInput.Type = "student"
			err = us.Create(&studentInput, v.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
