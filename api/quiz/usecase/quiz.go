package usecase

import (
	"errors"
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

	realistic := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "R",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa senang bekerja dengan mesin atau peralatan yang rumit, seperti alat berat atau kendaraan berat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "operating, maintaining, repairing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka mengatur proyek atau kegiatan yang melibatkan banyak orang, seperti acara besar atau proyek konstruksi yang kompleks",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "organizing, coordinating, delegating",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya menikmati bekerja dalam tim atau kelompok yang memiliki tujuan atau target tertentu",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "collaborating, team-building, problem-solving",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa terpanggil untuk membantu orang lain dalam hal teknis atau mekanis, seperti memperbaiki mesin atau alat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "troubleshooting, repairing, servicing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya memiliki ketertarikan atau keterampilan dalam bidang teknologi, seperti pengembangan perangkat lunak atau desain sistem, dan menikmati bekerja dengan komputer dan perangkat digital lainnya",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "programming, designing, utilizing software",
						Description: "",
					},
				},
			},
		},
	}

	for i := range realistic.Questions {
		realistic.Questions[i].QuizID = realistic.ID
		for j := range realistic.Questions[i].Options {
			realistic.Questions[i].Options[j].QuestionID = realistic.Questions[i].ID
		}
	}

	var realInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", realistic.Title).Count(&realInt).Error; err != nil {
		return err
	}
	if realInt == 0 {
		err := uq.uc.Create(realistic)
		if err != nil {
			return err
		}
	}

	investigative := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "I",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka menemukan jawaban atas pertanyaan-pertanyaan yang kompleks, melakukan riset, atau mengembangkan teori?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "researching, collecting, analyzing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dengan data atau informasi teknis dan dapat melakukan analisis yang kompleks, seperti membuat laporan atau melakukan evaluasi",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "data processing, analyzing, reporting",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka mengeksplorasi konsep baru atau ide-ide yang belum pernah ditemukan sebelumnya?",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "exploring, brainstorming, developing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya tertarik untuk mempelajari hal-hal yang kompleks atau terperinci, seperti teori matematika atau prinsip ilmu fisika",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "learning, understanding, analyzing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka menyelesaikan masalah teknis atau masalah yang memerlukan analisis logis dan rasional",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "problem-solving, logical-thinking, troubleshooting",
						Description: "",
					},
				},
			},
		},
	}

	for i := range investigative.Questions {
		investigative.Questions[i].QuizID = investigative.ID
		for j := range investigative.Questions[i].Options {
			investigative.Questions[i].Options[j].QuestionID = investigative.Questions[i].ID
		}
	}

	var invInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", investigative.Title).Count(&invInt).Error; err != nil {
		return err
	}
	if invInt == 0 {
		err := uq.uc.Create(investigative)
		if err != nil {
			return err
		}
	}

	artistic := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "A",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang seni rupa, seperti lukisan atau patung",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "visual, fine-art, painting, sculpture",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang desain, seperti desain grafis atau desain interior",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "design, graphic design, interior design",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang musik, seperti komposisi atau pertunjukan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "music, composition, perform",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang sastra, seperti penulisan atau penyuntingan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "writing, editing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang fotografi atau videografi",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "photography, videography, creative-thinking",
						Description: "",
					},
				},
			},
		},
	}
	for i := range artistic.Questions {
		artistic.Questions[i].QuizID = artistic.ID
		for j := range artistic.Questions[i].Options {
			artistic.Questions[i].Options[j].QuestionID = artistic.Questions[i].ID
		}
	}

	var artInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", artistic.Title).Count(&artInt).Error; err != nil {
		return err
	}
	if artInt == 0 {
		err := uq.uc.Create(artistic)
		if err != nil {
			return err
		}
	}

	social := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "S",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka mengajar atau membimbing orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "teaching, instructing, coaching",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka merawat dan membantu orang yang membutuhkan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "caring, helping, assisting",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka memberikan nasihat dan dukungan kepada orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "counseling, advising, supporting",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka melatih dan mengembangkan karyawan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "training, developing, coaching",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam tim atau kelompok",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "teamwork, collaboration, cooperation",
						Description: "",
					},
				},
			},
		},
	}
	for i := range social.Questions {
		social.Questions[i].QuizID = social.ID
		for j := range social.Questions[i].Options {
			social.Questions[i].Options[j].QuestionID = social.Questions[i].ID
		}
	}

	var socialInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", social.Title).Count(&socialInt).Error; err != nil {
		return err
	}
	if socialInt == 0 {
		err := uq.uc.Create(social)
		if err != nil {
			return err
		}
	}

	enterprising := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "E",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang penjualan atau pemasaran",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "sales, marketing, advertising",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang manajemen atau kepemimpinan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "management, leadership, supervision",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang keuangan atau investasi",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "finance, investment, accounting",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang kewirausahaan atau bisnis",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "entrepreneurship, business development, strategic planning",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang riset pasar atau analisis",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "market research, data analysis, trend forecasting",
						Description: "",
					},
				},
			},
		},
	}

	for i := range enterprising.Questions {
		enterprising.Questions[i].QuizID = enterprising.ID
		for j := range enterprising.Questions[i].Options {
			enterprising.Questions[i].Options[j].QuestionID = enterprising.Questions[i].ID
		}
	}

	var entInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", enterprising.Title).Count(&entInt).Error; err != nil {
		return err
	}
	if entInt == 0 {
		err := uq.uc.Create(enterprising)
		if err != nil {
			return err
		}
	}

	conventional := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "C",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang administrasi atau manajemen kantor",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "administration, management, office",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang akuntansi atau keuangan.",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "accounting, finance, bookkeeping",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang pengolahan data atau informasi.",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "data processing, management, data analysis",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang perencanaan atau pengorganisasian",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "planning, organizing, scheduling",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya suka bekerja dalam bidang sumber daya manusia atau pengembangan karyawan.",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "human resources, employee development, talent management",
						Description: "",
					},
				},
			},
		},
	}

	for i := range conventional.Questions {
		conventional.Questions[i].QuizID = conventional.ID
		for j := range conventional.Questions[i].Options {
			conventional.Questions[i].Options[j].QuestionID = conventional.Questions[i].ID
		}
	}

	var convInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", conventional.Title).Count(&convInt).Error; err != nil {
		return err
	}
	if convInt == 0 {
		err := uq.uc.Create(conventional)
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
	if query == "SectionOne" {
		arr := data.([]models.InputQuizString)
		stringArr := make([]string, 0, len(arr))
		for _, item := range arr {
			stringArr = append(stringArr, item.Data)
		}
		result := logic.MostFrequentElements(stringArr)

		idUser, err := uuid.FromString(userId)
		if err != nil {
			return err
		}

		resultUser := &models.QuizResult{
			ID:               uuid.Must(uuid.NewV4()),
			UserID:           idUser,
			ResultSectionOne: result,
		}
		return uq.uc.CreateResult(resultUser)
	}
	return exception.ErrNoQuery
}
