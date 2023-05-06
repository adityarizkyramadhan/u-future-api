package usecase

import (
	"errors"
	"u-future-api/api/quiz/repository"
	"u-future-api/models"
	"u-future-api/util/exception"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Quiz struct {
	uc *repository.Quiz
}

func New(uc *repository.Quiz) *Quiz {
	return &Quiz{uc}
}

func (uq *Quiz) GenerateQuestion(db *gorm.DB) error {
	// Section One
	sectionOne := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "SectionOne",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya suka mengeksplorasi ide baru dan mencari solusi kreatif untuk menyelesaikan masalah.",
						Description: "I",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka mengikuti prosedur yang telah terbukti efektif.",
						Description: "R",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya tertarik untuk mengikuti kegiatan seni.",
						Description: "A",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Seni bukan bidang yang menarik bagi saya.",
						Description: "C",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya menikmati bekerja dengan secara langsung di lapangan.",
						Description: "R",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka bekerja di kantor.",
						Description: "I",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya suka mengambil peran kepemimpinan dalam kelompok.",
						Description: "E",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka menjadi anggota kelompok daripada pemimpin.",
						Description: "S",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya tertarik untuk mempelajari teori dan konsep ilmiah yang kompleks.",
						Description: "I",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka mempelajari topik yang praktis dan langsung terkait dengan kehidupan sehari-hari.",
						Description: "C",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya menikmati membantu orang lain dan memberikan dukungan emosional.",
						Description: "S",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka fokus pada tugas dan proyek daripada mendukung orang lain.",
						Description: "E",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya suka bekerja dengan teknologi dan perangkat lunak komputer.",
						Description: "R",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka bekerja dengan orang atau alam.",
						Description: "A",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya senang berbicara di depan umum dan mempresentasikan ide-ide saya.",
						Description: "E",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka menyampaikan ide-ide saya secara tertulis atau dalam diskusi kecil.",
						Description: "R",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka bekerja dengan angka dan data.",
						Description: "I",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka berinteraksi dengan orang.",
						Description: "S",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Mana yang lebih sesuai dengan Anda?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya tertarik untuk mempelajari tentang budaya dan sejarah dunia.",
						Description: "A",
					},
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Saya lebih suka mempelajari topik yang terkait dengan saat ini dan praktis.",
						Description: "C",
					},
				},
			},
		},
	}
	for i := range sectionOne.Questions {
		sectionOne.Questions[i].QuizID = sectionOne.ID
		for j := range sectionOne.Questions[i].Options {
			sectionOne.Questions[i].Options[j].QuestionID = sectionOne.Questions[i].ID
		}
	}
	var secOne int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionOne.Title).Count(&secOne).Error; err != nil {
		return err
	}
	if secOne == 0 {
		err := uq.uc.Create(sectionOne)
		if err != nil {
			return err
		}
	}
	return nil

}

func (uq *Quiz) FindByName(name string) (*models.Quiz, error) {
	if name == "" {
		return nil, exception.ErrNoQuery
	}
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
