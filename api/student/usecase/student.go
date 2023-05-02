package usecase

import (
	"u-future-api/api/school/usecase"
	"u-future-api/api/student/repository"
	"u-future-api/middleware"
	"u-future-api/models"
	"u-future-api/util/exception"

	"github.com/go-faker/faker/v4"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	rs *repository.Student
}

func New(rs *repository.Student) *Student {
	return &Student{rs}
}
func (us *Student) create(arg *models.StudentInput, shoolId uuid.UUID) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	student := &models.Student{
		ID:       id,
		Name:     arg.Name,
		Email:    arg.Email,
		Password: string(hashPass),
		Type:     arg.Type,
		SchoolID: shoolId,
	}
	return us.rs.Create(student)
}

func (us *Student) FindById(id string) (*models.Student, error) {
	return us.rs.FindById(id)
}

func (us *Student) Login(arg *models.StudentLogin) (string, error) {
	if arg == nil {
		return "", exception.ErrNullPointer
	}
	student, err := us.rs.FindByEmail(arg.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(arg.Password)); err != nil {
		return "", exception.ErrIncorectPassword
	}
	token, err := middleware.GenerateJWToken(student.ID.String())
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *Student) Register(arg *models.StudentRegister) (string, error) {
	idSchool, err := uuid.FromString(arg.SchoolID)
	if err != nil {
		return "", err
	}
	var input models.StudentInput
	if err := copier.Copy(&input, arg); err != nil {
		return "", err
	}
	if err := us.create(&input, idSchool); err != nil {
		return "", err
	}
	student, err := us.rs.FindByEmail(arg.Email)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWToken(student.ID.String())
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *Student) GenerateFaker(ucSchool *usecase.School) error {
	schools, err := ucSchool.FindAll()
	if err != nil {
		return err
	}
	var tempSchId uuid.UUID
	for _, v := range schools {
		for i := 0; i < 10; i++ {
			var studentInput models.StudentInput
			err = faker.FakeData(&studentInput)
			if err != nil {
				return err
			}
			studentInput.Type = "student"
			err = us.create(&studentInput, v.ID)
			if err != nil {
				return err
			}
		}
		tempSchId = v.ID
	}
	// Test Account
	err = us.create(&models.StudentInput{
		Name:     "Test Account",
		Email:    "testing@gmail.com",
		Password: "12345678",
		Type:     "student",
	}, tempSchId)
	if err != nil {
		return err
	}
	return nil
}
