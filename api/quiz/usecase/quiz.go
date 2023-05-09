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

	sectionThreeOne := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "teaching, coaching, leadership, counseling",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menyampaikan informasi dengan jelas",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Communication, Public Speaking, Clarity, Presentation",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memahami kebutuhan dan perasaan orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Empathy, Emotional Intelligence, Understanding, Listening",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menyesuaikan metode pengajaran sesuai dengan kebutuhan individu",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Adaptability, Customization, Personalization, Curriculum Development",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memotivasi dan menginspirasi orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Motivation, Inspiration, Leadership, Coaching, Mentoring",
						Description: "",
					},
				},
			},
		},
	}

	for i := range sectionThreeOne.Questions {
		sectionThreeOne.Questions[i].QuizID = sectionThreeOne.ID
		for j := range sectionThreeOne.Questions[i].Options {
			sectionThreeOne.Questions[i].Options[j].QuestionID = sectionThreeOne.Questions[i].ID
		}
	}

	var secThreeeOneInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeOne.Title).Count(&secThreeeOneInt).Error; err != nil {
		return err
	}
	if secThreeeOneInt == 0 {
		err := uq.uc.Create(sectionThreeOne)
		if err != nil {
			return err
		}
	}

	sectionThreeTwo := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "operating, repairing, maintaining",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengoperasikan suatu alat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Equipment Operation, Machinery Handling, Technical Skill, Technical Understanding, Instruction Following, Manual Reading",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk bekerja dengan teliti dan akurat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Attention to Detail, Precision, Accuracy, Troubleshooting",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memperbaiki dan merawat peralatan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Equipment Repair, Maintenance, Servicing, Troubleshooting",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeTwo.Questions {
		sectionThreeTwo.Questions[i].QuizID = sectionThreeTwo.ID
		for j := range sectionThreeTwo.Questions[i].Options {
			sectionThreeTwo.Questions[i].Options[j].QuestionID = sectionThreeTwo.Questions[i].ID
		}
	}

	var secThreeTwoInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeTwo.Title).Count(&secThreeTwoInt).Error; err != nil {
		return err
	}
	if secThreeTwoInt == 0 {
		err := uq.uc.Create(sectionThreeTwo)
		if err != nil {
			return err
		}
	}

	sectionThreeThree := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "research, analyzing",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk merancang dan melaksanakan suatu penelitian",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Experimental Design, Scientific Method, Research Skill, Planning, Analytical Thinking",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menganalisis dan menafsirkan data hasil penelitian",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Data Analysis, Statistics, Interpretation, Critical Thinking, Problem Solving",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya meresa memiliki kemampuan untuk bekerja dengan teliti dan akurat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Attention to Detail, Precision, Accuracy, Logical Thinking",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memecahkan masalah yang kompleks",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Problem Solving, Critical Thinking, Innovation, Troubleshooting, Analytical Thinking",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeThree.Questions {
		sectionThreeThree.Questions[i].QuizID = sectionThreeThree.ID
		for j := range sectionThreeThree.Questions[i].Options {
			sectionThreeThree.Questions[i].Options[j].QuestionID = sectionThreeThree.Questions[i].ID
		}
	}

	var secThreeThreeInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeThree.Title).Count(&secThreeThreeInt).Error; err != nil {
		return err
	}
	if secThreeThreeInt == 0 {
		err := uq.uc.Create(sectionThreeThree)
		if err != nil {
			return err
		}
	}

	sectionThreeFour := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "programming",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memahami logika pemrograman dan bahasa pemrograman",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Programming Languages, Logical Thinking, Algorithm Design",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menyelesaikan masalah teknis yang rumit dalam pemrograman",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Problem Solving, Troubleshooting, Debugging, Software Development",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeFour.Questions {
		sectionThreeFour.Questions[i].QuizID = sectionThreeFour.ID
		for j := range sectionThreeFour.Questions[i].Options {
			sectionThreeFour.Questions[i].Options[j].QuestionID = sectionThreeFour.Questions[i].ID
		}
	}

	var secThreeFourInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeFour.Title).Count(&secThreeFourInt).Error; err != nil {
		return err
	}
	if secThreeFourInt == 0 {
		err := uq.uc.Create(sectionThreeFour)
		if err != nil {
			return err
		}
	}

	sectionThreeFive := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "writing",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menulis dengan gaya dan nada yang tepat untuk audiens yang dituju",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Writing Style, Audience Understanding, Content Creation, Creative Writing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk merancang konten yang informatif dan menarik",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Content Creation, Blogging, Copywriting, Journalism",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk meneliti topik dan menulis artikel yang informatif",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Research, Article Writing, Interviewing, Fact-Checking",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengedit dan menyunting tulisan untuk memastikan kualitasnya",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "Editing, Proofreading, Copyediting, Grammar and Spelling, Style Guide Compliance",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeFive.Questions {
		sectionThreeFive.Questions[i].QuizID = sectionThreeFive.ID
		for j := range sectionThreeFive.Questions[i].Options {
			sectionThreeFive.Questions[i].Options[j].QuestionID = sectionThreeFive.Questions[i].ID
		}
	}

	var secThreeFiveInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeFive.Title).Count(&secThreeFiveInt).Error; err != nil {
		return err
	}
	if secThreeFiveInt == 0 {
		err := uq.uc.Create(sectionThreeFive)
		if err != nil {
			return err
		}
	}

	sectionThreeSix := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "designing",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk berpikir kreatif dan menghasilkan ide-ide yang unik",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "creative thinking, innovative, artistic, aesthetic, visual thinking",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya terampil dalam mengoperasikan software dan teknologi terkait visual, seperti Adobe Photoshop, Illustrator, atau program editing video",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "visual, creative, communicative, technical, utilizing software",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya mampu memahami brief dari klien dan menghasilkan karya yang sesuai dengan keinginan mereka",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "logic, technical, problem-solving, organizing, coordination, utilizing software",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya memiliki pengalaman dalam mengekspresikan ide-ide atau gagasan melalui media visual seperti lukisan, fotografi, atau desain grafis",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "visual, creative, communicative, fine-art, painting, sculpture, design, graphic design, interior design, music, composition, performing, photography, videography, utilizing software",
						Description: "",
					},
				},
			},
		},
	}

	for i := range sectionThreeSix.Questions {
		sectionThreeSix.Questions[i].QuizID = sectionThreeSix.ID
		for j := range sectionThreeSix.Questions[i].Options {
			sectionThreeSix.Questions[i].Options[j].QuestionID = sectionThreeSix.Questions[i].ID
		}
	}

	var secThreeSixInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeSix.Title).Count(&secThreeSixInt).Error; err != nil {
		return err
	}
	if secThreeSixInt == 0 {
		err := uq.uc.Create(sectionThreeSix)
		if err != nil {
			return err
		}
	}

	secThreeSeven := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "administration, management, office, accounting",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk merencanakan dan mengorganisir tugas-tugas kantor",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "planning, organizing, scheduling, management",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengolah dan menganalisis data dengan akurasi",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "data processing, data analysis, statistics",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengelola sumber daya manusia, termasuk merekrut, melatih, dan mengembangkan karyawan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "human resources, employee development, talent management, training, developing",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menjaga buku dan catatan keuangan yang akurat",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "bookkeeping, accounting, finance, management",
						Description: "",
					},
				},
			},
		},
	}

	for i := range secThreeSeven.Questions {
		secThreeSeven.Questions[i].QuizID = secThreeSeven.ID
		for j := range secThreeSeven.Questions[i].Options {
			secThreeSeven.Questions[i].Options[j].QuestionID = secThreeSeven.Questions[i].ID
		}
	}

	var secThreeSevenInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", secThreeSeven.Title).Count(&secThreeSevenInt).Error; err != nil {
		return err
	}
	if secThreeSevenInt == 0 {
		err := uq.uc.Create(secThreeSeven)
		if err != nil {
			return err
		}
	}

	sectionThreeEight := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "managing, organizing, supervising",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengorganisir tugas dan proyek dengan efisien",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "project management, time management, team management, delegation",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menetapkan tujuan dan strategi untuk mencapai hasil yang diinginkan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "leadership, strategic planning",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk melakukan evaluasi dan perbaikan terhadap proses kerja",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "performance evaluation, process improvement",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeEight.Questions {
		sectionThreeEight.Questions[i].QuizID = sectionThreeEight.ID
		for j := range sectionThreeEight.Questions[i].Options {
			sectionThreeEight.Questions[i].Options[j].QuestionID = sectionThreeEight.Questions[i].ID
		}
	}

	var secThreeEightInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeEight.Title).Count(&secThreeEightInt).Error; err != nil {
		return err
	}
	if secThreeEightInt == 0 {
		err := uq.uc.Create(sectionThreeEight)
		if err != nil {
			return err
		}
	}
	sectionThreeNine := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "caring, helping, assisting",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk memahami dan merespons kebutuhan dan masalah orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "empathy, compassion, counseling, support",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk bekerja dengan orang-orang yang memiliki berbagai macam latar belakang dan kebutuhan",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "social skills, problem-solving, counseling, support, assistance, diversity, inclusion",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeNine.Questions {
		sectionThreeNine.Questions[i].QuizID = sectionThreeNine.ID
		for j := range sectionThreeNine.Questions[i].Options {
			sectionThreeNine.Questions[i].Options[j].QuestionID = sectionThreeNine.Questions[i].ID
		}
	}

	var secThreeNineInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeNine.Title).Count(&secThreeNineInt).Error; err != nil {
		return err
	}
	if secThreeNineInt == 0 {
		err := uq.uc.Create(sectionThreeNine)
		if err != nil {
			return err
		}
	}
	sectionThreeTen := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "creating, innovating, ideating",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk berpikir kreatif dan menghasilkan ide-ide baru",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "creativity, innovation, ideation",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk mengembangkan solusi baru dan inovatif untuk masalah yang kompleks",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "brainstorming, problem-solving, design thinking, product development",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeTen.Questions {
		sectionThreeTen.Questions[i].QuizID = sectionThreeTen.ID
		for j := range sectionThreeTen.Questions[i].Options {
			sectionThreeTen.Questions[i].Options[j].QuestionID = sectionThreeTen.Questions[i].ID
		}
	}

	var secThreeTenInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeTen.Title).Count(&secThreeTenInt).Error; err != nil {
		return err
	}
	if secThreeTenInt == 0 {
		err := uq.uc.Create(sectionThreeTen)
		if err != nil {
			return err
		}
	}

	sectionThreeEleven := &models.Quiz{
		ID:    uuid.Must(uuid.NewV4()),
		Title: "performing, entertaining, directing",
		Questions: []models.Question{
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk tampil dan menghibur orang lain",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "performing arts, entertainment industry, creative expression",
						Description: "",
					},
				},
			},
			{
				ID:   uuid.Must(uuid.NewV4()),
				Text: "Saya merasa memiliki kemampuan untuk menciptakan dan mengembangkan karya seni dan hiburan yang bermakna dan berpengaruh",
				Options: []models.Option{
					{
						ID:          uuid.Must(uuid.NewV4()),
						Text:        "producing, direction, leadership",
						Description: "",
					},
				},
			},
		},
	}
	for i := range sectionThreeEleven.Questions {
		sectionThreeEleven.Questions[i].QuizID = sectionThreeEleven.ID
		for j := range sectionThreeEleven.Questions[i].Options {
			sectionThreeEleven.Questions[i].Options[j].QuestionID = sectionThreeEleven.Questions[i].ID
		}
	}

	var secThreeElevenInt int64
	if err := db.Model(&models.Quiz{}).Where("title = ?", sectionThreeEleven.Title).Count(&secThreeElevenInt).Error; err != nil {
		return err
	}
	if secThreeElevenInt == 0 {
		err := uq.uc.Create(sectionThreeEleven)
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
