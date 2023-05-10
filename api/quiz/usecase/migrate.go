package usecase

import (
	"u-future-api/models"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

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

func (uq *Quiz) GenerateJurusan(db *gorm.DB) error {
	jurusans := []models.Jurusan{
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Akuntansi",
			RIASECCode:            "C, E, S",
			WorkActivities:        "bookkeeping, accounting, data processing, data analysis, management, finance, planning, organizing",
			Values:                "accuracy, attention to detail, precision, process improvement, time management",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.9,
			TingkatProspekKerja:   0.85,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Auditor, Tax Consultant, Financial Controller, Investment Banker, Management Accountant",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Antropologi Sosial",
			RIASECCode:            "S, E, A",
			WorkActivities:        "exploring, researching, analyzing, data processing, reporting, collecting, brainstorming, learning",
			Values:                "creativity, innovation, diversity, inclusion, empathy",
			TingkatKeketatan:      0.16,
			TingkatKeselarasan:    0.6,
			TingkatDapatPekerjaan: 0.5,
			TingkatProspekKerja:   0.55,
			Gaji:                  "75 - 90 Juta",
			Pekerjaan:             "Cultural Resource Manager, Human Resource Manager, Social Worker, Public Policy Analyst, Market Research Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Arkeologi",
			RIASECCode:            "S, R, A",
			WorkActivities:        "exploring, researching, analyzing, data processing, reporting, collecting, brainstorming, learning",
			Values:                "creativity, innovation, diversity, inclusion, empathy",
			TingkatKeketatan:      0.13,
			TingkatKeselarasan:    0.65,
			TingkatDapatPekerjaan: 0.45,
			TingkatProspekKerja:   0.55,
			Gaji:                  "80 - 100 Juta",
			Pekerjaan:             "Cultural Resource Manager, Museum Curator, Heritage Manager, Archaeologist, Art Historian",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Arsitektur",
			RIASECCode:            "A, E, R",
			WorkActivities:        "designing, utilizing software, researching, analyzing, exploring, brainstorming, developing, understanding",
			Values:                "creativity, innovation, visual, aesthetic, technical",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.78,
			Gaji:                  "150 - 250 Juta",
			Pekerjaan:             "Architect, Urban Designer, Interior Designer, Landscape Architect, Construction Manager",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Biologi",
			RIASECCode:            "I, S, R",
			WorkActivities:        "exploring, researching, collecting, analyzing, data processing, reporting, learning",
			Values:                "critical thinking, analytical thinking, precision, accuracy",
			TingkatKeketatan:      0.14,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.6,
			TingkatProspekKerja:   0.65,
			Gaji:                  "80 - 120 Juta",
			Pekerjaan:             "Biotechnologist, Bioinformatician, Biomedical Scientist, Ecologist, Conservationist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Farmasi",
			RIASECCode:            "I, S, E",
			WorkActivities:        "researching, analyzing, dispensing medication, counseling, communicating",
			Values:                "attention to detail, organization, problem-solving",
			TingkatKeketatan:      0.06,
			TingkatKeselarasan:    0.93,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.89,
			Gaji:                  "150 - 200 Juta",
			Pekerjaan:             "Apoteker, Clinical Researcher, Farmasi Industri",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Fisika",
			RIASECCode:            "I, R, A",
			WorkActivities:        "researching, analyzing, problem-solving, developing theories, conducting experiments",
			Values:                "logical thinking, mathematical skills",
			TingkatKeketatan:      0.44,
			TingkatKeselarasan:    0.87,
			TingkatDapatPekerjaan: 0.6,
			TingkatProspekKerja:   0.74,
			Gaji:                  "100 - 150 Juta",
			Pekerjaan:             "Fisikawan, Ahli Geofisika, Ahli Kimia",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Fisioterapi",
			RIASECCode:            "S, E, A",
			WorkActivities:        "assessing patients, designing treatment plans, implementing therapies, educating patients",
			Values:                "compassion, empathy, attention to detail",
			TingkatKeketatan:      0.09,
			TingkatKeselarasan:    0.77,
			TingkatDapatPekerjaan: 0.5,
			TingkatProspekKerja:   0.64,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Fisioterapis, Ahli Rehabilitasi, Ahli Kesehatan Olahraga",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Geofisika",
			RIASECCode:            "I, R, A",
			WorkActivities:        "conducting research, analyzing data, using geophysical equipment, interpreting geological phenomena",
			Values:                "analytical thinking, mathematical skills, problem-solving",
			TingkatKeketatan:      0.4,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.6,
			TingkatProspekKerja:   0.7,
			Gaji:                  "100 - 150 Juta",
			Pekerjaan:             "Ahli Geofisika, Ahli Geologi, Peneliti Bumi",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Geografi",
			RIASECCode:            "I, S, E",
			WorkActivities:        "collecting data, analyzing data, conducting research, creating maps and charts",
			Values:                "analytical thinking, attention to detail, problem-solving",
			TingkatKeketatan:      0.17,
			TingkatKeselarasan:    0.73,
			TingkatDapatPekerjaan: 0.45,
			TingkatProspekKerja:   0.59,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Kartografer, Pengamat Lingkungan, Ahli Pemetaan",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Geologi",
			RIASECCode:            "R, I, E",
			WorkActivities:        "exploring, collecting, analyzing, problem-solving, troubleshooting",
			Values:                "creativity, innovation, ideation",
			TingkatKeketatan:      0.4,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.83,
			Gaji:                  "125 - 175 Juta",
			Pekerjaan:             "Geologist, Mining Engineer, Environmental Scientist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Gizi",
			RIASECCode:            "S, E, C",
			WorkActivities:        "counseling, advising, supporting, caring, helping, assisting",
			Values:                "compassion, empathy, diversity",
			TingkatKeketatan:      0.03,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.9,
			TingkatProspekKerja:   0.83,
			Gaji:                  "50 - 75 Juta",
			Pekerjaan:             "Nutritionist, Dietitian, Food Scientist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Hubungan Internasional",
			RIASECCode:            "S, E, A",
			WorkActivities:        "teaching, instructing, coaching, networking, relationship building",
			Values:                "leadership, strategic planning, performance evaluation",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.83,
			Gaji:                  "100 - 125 Juta",
			Pekerjaan:             "Diplomat, International Trade Specialist, Policy Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Hubungan Masyarakat",
			RIASECCode:            "S, E, A",
			WorkActivities:        "public speaking, networking, sales, marketing, advertising",
			Values:                "creativity, strategic planning, teamwork",
			TingkatKeketatan:      0.12,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.75,
			TingkatProspekKerja:   0.73,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Public Relations Specialist, Communications Specialist, Event Planner",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Aktuaria",
			RIASECCode:            "I, R, C",
			WorkActivities:        "data processing, analyzing, problem-solving, statistical modeling",
			Values:                "precision, accuracy, data analysis",
			TingkatKeketatan:      0.27,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.83,
			Gaji:                  "200 - 250 Juta",
			Pekerjaan:             "Actuary, Risk Manager, Data Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Ekonomi",
			RIASECCode:            "I, R, E",
			WorkActivities:        "analyzing, researching, collecting, interpreting, reporting",
			Values:                "creativity, innovation, ideation",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.78,
			Gaji:                  "100 - 125 Juta",
			Pekerjaan:             "Economist, Market Research Analyst, Investment Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Filsafat",
			RIASECCode:            "A, I, E",
			WorkActivities:        "researching, analyzing, interpreting, writing, teaching",
			Values:                "creativity, innovation, ideation",
			TingkatKeketatan:      0.8,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.75,
			TingkatProspekKerja:   0.73,
			Gaji:                  "50 - 75 Juta",
			Pekerjaan:             "Philosopher, Researcher, Writer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Hukum",
			RIASECCode:            "S, E, A",
			WorkActivities:        "analyzing, researching, interpreting, advising, arguing",
			Values:                "leadership, strategic planning, performance evaluation",
			TingkatKeketatan:      0.05,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.83,
			Gaji:                  "100 - 125 Juta",
			Pekerjaan:             "Lawyer, Judge, Legal Consultant",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Keperawatan",
			RIASECCode:            "S, E, C",
			WorkActivities:        "caring, helping, assisting, counseling, instructing",
			Values:                "compassion, empathy, diversity",
			TingkatKeketatan:      0.04,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.9,
			TingkatProspekKerja:   0.88,
			Gaji:                  "50 - 75 Juta",
			Pekerjaan:             "Nurse, Healthcare Administrator, Clinical Research Coordinator",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Komputer",
			RIASECCode:            "R, I, E",
			WorkActivities:        "programming, designing, troubleshooting, researching, analyzing",
			Values:                "creativity, innovation, ideation",
			TingkatKeketatan:      0.06,
			TingkatKeselarasan:    0.9,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.88,
			Gaji:                  "175 - 225 Juta",
			Pekerjaan:             "Software Engineer, Cybersecurity Analyst, Data Scientist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Komunikasi",
			RIASECCode:            "R, E, S",
			WorkActivities:        "networking, public speaking, counseling, advising, assisting, collaborating",
			Values:                "communication, empathy, creativity",
			TingkatKeketatan:      0.17,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.82,
			TingkatProspekKerja:   0.79,
			Gaji:                  "75 - 90 Juta",
			Pekerjaan:             "PR Specialist, Journalist, Advertising Executive, Communication Specialist, Event Organizer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Politik",
			RIASECCode:            "S, E, A",
			WorkActivities:        "researching, collecting, analyzing, data processing, reporting, exploring, brainstorming, developing",
			Values:                "inclusion, diversity, problem-solving",
			TingkatKeketatan:      0.13,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.83,
			Gaji:                  "100 - 150 Juta",
			Pekerjaan:             "Political Analyst, Policy Researcher, Lobbyist, Political Campaign Manager, Legislative Aide",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Ilmu Sejarah",
			RIASECCode:            "I, A, E",
			WorkActivities:        "researching, collecting, analyzing, data processing, reporting, exploring, brainstorming, developing",
			Values:                "creativity, diversity, innovation",
			TingkatKeketatan:      0.25,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.5,
			TingkatProspekKerja:   0.6,
			Gaji:                  "80 - 100 Juta",
			Pekerjaan:             "Historian, Archivist, Museum Curator, Researcher, Writer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Kesehatan Masyarakat",
			RIASECCode:            "S, I, E",
			WorkActivities:        "researching, collecting, analyzing, data processing, reporting, exploring, brainstorming, developing, instructing, coaching, helping, assisting",
			Values:                "empathy, inclusion, creativity",
			TingkatKeketatan:      0.05,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.95,
			TingkatProspekKerja:   0.9,
			Gaji:                  "120 - 200 Juta",
			Pekerjaan:             "Epidemiologist, Health Educator, Public Health Analyst, Health Administrator, Biostatistician",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Kewirausahaan",
			RIASECCode:            "E, S, C",
			WorkActivities:        "entrepreneurship, business development, strategic planning, market research, data analysis, administration, management, leadership",
			Values:                "innovation, creativity, leadership",
			TingkatKeketatan:      0.16,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.8,
			Gaji:                  "150 - 300 Juta",
			Pekerjaan:             "Entrepreneur, Business Consultant, Marketing Manager, Project Manager, Sales Executive",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Kimia Terapan",
			RIASECCode:            "R, I, C",
			WorkActivities:        "analyzing, troubleshooting, servicing, exploring, collecting, researching",
			Values:                "innovation, ideation, brainstorming, problem-solving",
			TingkatKeketatan:      0.5,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.65,
			TingkatProspekKerja:   0.7,
			Gaji:                  "75 - 90 Juta",
			Pekerjaan:             "Chemical Technician, Quality Control Analyst, Research Scientist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Kimia",
			RIASECCode:            "I, R, C",
			WorkActivities:        "analyzing, researching, collecting, exploring, troubleshooting, servicing",
			Values:                "innovation, ideation, brainstorming, problem-solving",
			TingkatKeketatan:      0.18,
			TingkatKeselarasan:    0.82,
			TingkatDapatPekerjaan: 0.73,
			TingkatProspekKerja:   0.78,
			Gaji:                  "85 - 110 Juta",
			Pekerjaan:             "Chemist, Research Scientist, Quality Control Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Manajemen Bisnis",
			RIASECCode:            "E, S, C",
			WorkActivities:        "planning, organizing, scheduling, managing, strategizing, networking, collaborating, supervising, team-building",
			Values:                "leadership, delegation, time management",
			TingkatKeketatan:      0.09,
			TingkatKeselarasan:    0.9,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.88,
			Gaji:                  "150 - 250 Juta",
			Pekerjaan:             "Business Analyst, Project Manager, Marketing Manager",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Matematika",
			RIASECCode:            "I, R, A",
			WorkActivities:        "analyzing, researching, collecting, exploring, problem-solving, troubleshooting, data processing, statistical analysis",
			Values:                "innovation, ideation, brainstorming, problem-solving",
			TingkatKeketatan:      0.13,
			TingkatKeselarasan:    0.78,
			TingkatDapatPekerjaan: 0.69,
			TingkatProspekKerja:   0.74,
			Gaji:                  "85 - 110 Juta",
			Pekerjaan:             "Actuary, Data Analyst, Research Scientist",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Pariwisata",
			RIASECCode:            "S, E, A",
			WorkActivities:        "planning, organizing, scheduling, coordinating, presenting, public speaking, networking, collaborating, assisting, counseling",
			Values:                "diversity, inclusion, empathy, compassion",
			TingkatKeketatan:      0.1,
			TingkatKeselarasan:    0.68,
			TingkatDapatPekerjaan: 0.58,
			TingkatProspekKerja:   0.63,
			Gaji:                  "80 - 95 Juta",
			Pekerjaan:             "Tourism Manager, Event Planner, Travel Consultant",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Pendidikan Dokter Gigi",
			RIASECCode:            "S, I, A",
			WorkActivities:        "analyzing, problem-solving, collaborating, organizing, communication",
			Values:                "compassion, teamwork, leadership",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.95,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.9,
			Gaji:                  "200 - 300 Juta",
			Pekerjaan:             "Dokter gigi umum, dokter gigi anak, dokter gigi spesialis, dosen kedokteran gigi, peneliti kedokteran gigi",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Pendidikan Dokter",
			RIASECCode:            "S, I, A",
			WorkActivities:        "analyzing, problem-solving, collaborating, organizing, communication",
			Values:                "compassion, teamwork, leadership",
			TingkatKeketatan:      0.04,
			TingkatKeselarasan:    0.97,
			TingkatDapatPekerjaan: 0.87,
			TingkatProspekKerja:   0.92,
			Gaji:                  "300 - 400 Juta",
			Pekerjaan:             "Dokter umum, dokter spesialis, dokter gigi umum, dosen kedokteran, peneliti kedokteran",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Pendidikan Teknologi Informasi",
			RIASECCode:            "R, I, A",
			WorkActivities:        "programming, designing, utilizing software, researching, analyzing, data processing, troubleshooting",
			Values:                "logical-thinking, problem-solving, teamwork",
			TingkatKeketatan:      0.15,
			TingkatKeselarasan:    0.92,
			TingkatDapatPekerjaan: 0.78,
			TingkatProspekKerja:   0.85,
			Gaji:                  "80 - 120 Juta",
			Pekerjaan:             "Programmer, web developer, system analyst, data analyst, database administrator",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Psikologi",
			RIASECCode:            "S, E, C",
			WorkActivities:        "counseling, advising, supporting, teaching, instructing, coaching, training, researching, analyzing, collecting",
			Values:                "empathy, compassion, social skills",
			TingkatKeketatan:      0.03,
			TingkatKeselarasan:    0.89,
			TingkatDapatPekerjaan: 0.82,
			TingkatProspekKerja:   0.86,
			Gaji:                  "80 - 150 Juta",
			Pekerjaan:             "Psikolog klinis, psikolog organisasi, konsultan psikologi, peneliti psikologi, pengajar psikologi",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sastra Arab",
			RIASECCode:            "A, S, E",
			WorkActivities:        "writing, editing, translating, researching, analyzing, instructing, teaching",
			Values:                "creativity, knowledge-sharing, critical thinking",
			TingkatKeketatan:      1.03,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.67,
			TingkatProspekKerja:   0.71,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Penulis, penerjemah, pengajar bahasa Arab, konsultan bahasa, peneliti sastra Arab",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sastra Cina",
			RIASECCode:            "A, S, E",
			WorkActivities:        "writing, editing, translating, researching, analyzing, instructing, teaching",
			Values:                "creativity, knowledge-sharing, critical thinking",
			TingkatKeketatan:      1.75,
			TingkatKeselarasan:    0.73,
			TingkatDapatPekerjaan: 0.65,
			TingkatProspekKerja:   0.69,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Penulis, penerjemah, pengajar bahasa Mandarin, konsultan bahasa, peneliti sastra Cina",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sastra Indonesia",
			RIASECCode:            "A, E, I",
			WorkActivities:        "writing, editing, translating, interpreting, publishing, presenting, public speaking, creative-thinking, imaginative, story-telling, research",
			Values:                "creativity, empathy, open-mindedness",
			TingkatKeketatan:      0.29,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.65,
			TingkatProspekKerja:   0.75,
			Gaji:                  "75 - 100 Juta",
			Pekerjaan:             "Penulis, Penerjemah, Editor",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sastra Inggris",
			RIASECCode:            "A, E, I",
			WorkActivities:        "writing, editing, translating, interpreting, publishing, presenting, public speaking, creative-thinking, imaginative, story-telling, research",
			Values:                "creativity, empathy, open-mindedness",
			TingkatKeketatan:      0.18,
			TingkatKeselarasan:    0.83,
			TingkatDapatPekerjaan: 0.68,
			TingkatProspekKerja:   0.76,
			Gaji:                  "75 - 120 Juta",
			Pekerjaan:             "Penulis, Editor, Translator",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sastra Jepang",
			RIASECCode:            "A, E, I",
			WorkActivities:        "writing, editing, translating, interpreting, publishing, presenting, public speaking, creative-thinking, imaginative, story-telling, research",
			Values:                "creativity, empathy, open-mindedness",
			TingkatKeketatan:      0.58,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.7,
			TingkatProspekKerja:   0.75,
			Gaji:                  "75 - 120 Juta",
			Pekerjaan:             "Penulis, Editor, Penerjemah",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sistem Informasi",
			RIASECCode:            "R, I, A",
			WorkActivities:        "analyzing, problem-solving, collaborating, organizing, management",
			Values:                "communication, empathy, leadership",
			TingkatKeketatan:      0.05,
			TingkatKeselarasan:    0.98,
			TingkatDapatPekerjaan: 0.77,
			TingkatProspekKerja:   0.88,
			Gaji:                  "80 - 150 Juta",
			Pekerjaan:             "System Analyst, Business Analyst, Proudct Manager, UI Designer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Sosiologi",
			RIASECCode:            "S, E, C",
			WorkActivities:        "team-building, collaborating, research, collecting, analyzing, reporting, counseling, advising, supporting",
			Values:                "diversity, inclusion, empathy, compassion",
			TingkatKeketatan:      0.07,
			TingkatKeselarasan:    0.8,
			TingkatDapatPekerjaan: 0.85,
			TingkatProspekKerja:   0.83,
			Gaji:                  "70 - 110 Juta",
			Pekerjaan:             "Social Worker, Community Organizer, Researcher",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Statistika",
			RIASECCode:            "I, A, S",
			WorkActivities:        "analyzing, collecting, organizing, interpreting data, modeling, data visualization, programming",
			Values:                "precision, accuracy, attention to detail",
			TingkatKeketatan:      0.45,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.75,
			TingkatProspekKerja:   0.8,
			Gaji:                  "90 - 130 Juta",
			Pekerjaan:             "Data Analyst, Statistician, Actuary",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Elektro",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, analyzing, troubleshooting, testing, installing, maintaining, repairing, researching, programming",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.2,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.7,
			TingkatProspekKerja:   0.73,
			Gaji:                  "110 - 180 Juta",
			Pekerjaan:             "Electrical Engineer, Control System Engineer, Telecommunications Engineer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Industri",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, implementing, analyzing, improving, testing, evaluating, researching, project management",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.21,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.65,
			TingkatProspekKerja:   0.68,
			Gaji:                  "90 - 150 Juta",
			Pekerjaan:             "Industrial Engineer, Quality Control Engineer, Production Manager",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Informatika",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, testing, debugging, troubleshooting, analyzing, researching, programming",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.03,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.83,
			Gaji:                  "110 - 190 Juta",
			Pekerjaan:             "Software Engineer, Web Developer, Database Administrator",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Kimia",
			RIASECCode:            "R, I, E",
			WorkActivities:        "analyzing, problem-solving, experimenting, designing, researching",
			Values:                "technical, analytical, critical thinking",
			TingkatKeketatan:      0.39,
			TingkatKeselarasan:    0.93,
			TingkatDapatPekerjaan: 0.75,
			TingkatProspekKerja:   0.84,
			Gaji:                  "100 - 150 Juta",
			Pekerjaan:             "Process Engineer, Product Engineer, Researcher",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Komputer",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, testing, debugging, troubleshooting, analyzing, researching, programming",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.49,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.83,
			Gaji:                  "120 - 200 Juta",
			Pekerjaan:             "Computer Hardware Engineer, Network and Systems Administrator, Information Security Analyst",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Mesin",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, analyzing, testing, troubleshooting, installing, maintaining, repairing, researching",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.22,
			TingkatKeselarasan:    0.75,
			TingkatDapatPekerjaan: 0.7,
			TingkatProspekKerja:   0.73,
			Gaji:                  "90 - 150 Juta",
			Pekerjaan:             "Mechanical Engineer, Automotive Engineer, Manufacturing Engineer",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknik Sipil",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, analyzing, testing, troubleshooting, estimating, managing, researching",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.07,
			TingkatKeselarasan:    0.7,
			TingkatDapatPekerjaan: 0.65,
			TingkatProspekKerja:   0.68,
			Gaji:                  "90 - 160 Juta",
			Pekerjaan:             "Civil Engineer, Construction Manager, Urban and Regional Planner",
		},
		{
			ID:                    uuid.Must(uuid.NewV4()),
			NamaJurusan:           "Teknologi Informasi",
			RIASECCode:            "I, R, A",
			WorkActivities:        "designing, developing, testing, debugging, troubleshooting, analyzing, researching, programming",
			Values:                "precision, accuracy, technical skill",
			TingkatKeketatan:      0.08,
			TingkatKeselarasan:    0.85,
			TingkatDapatPekerjaan: 0.8,
			TingkatProspekKerja:   0.83,
			Gaji:                  "100 - 180 Juta",
			Pekerjaan:             "Software Developer, System Analyst, IT Project Manager",
		},
	}
	var jurusanInt int64
	if err := db.Model(&models.Jurusan{}).Count(&jurusanInt).Error; err != nil {
		return err
	}
	if jurusanInt == 0 {
		for _, v := range jurusans {
			err := db.Transaction(func(tx *gorm.DB) error {
				return tx.Create(&v).Error
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
