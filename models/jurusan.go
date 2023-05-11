package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Jurusan struct {
	ID                    uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	NamaJurusan           string         `json:"nama_jurusan"`
	RIASECCode            string         `json:"RIASEC_code"`
	WorkActivities        string         `json:"work_activities"`
	Values                string         `json:"values"`
	TingkatKeketatan      float64        `json:"tingkat_keketatan"`
	TingkatKeselarasan    float64        `json:"tingkat_keselarasan"`
	TingkatDapatPekerjaan float64        `json:"tingkat_dapat_pekerjaan"`
	TingkatProspekKerja   float64        `json:"tingkat_prospek_kerja"`
	Gaji                  string         `json:"gaji"`
	Pekerjaan             string         `json:"pekerjaan"`
	Deskripsi             string         `json:"deskripsi"`
	TagJurusan            string         `json:"tag_jurusan"`
	Jurusan               string         `json:"jurusan"`
	MataKuliah            string         `json:"mata_kuliah"`
	IDKampus              string         `json:"id_kampus"`
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type JurusanStudentCompare struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID      uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Percentage  float64   `json:"percentage"`
	NamaJurusan string    `json:"nama_jurusan"`
	TagJurusan  string    `json:"tag_jurusan"`
	Jurusan     string    `json:"jurusan"`
}

type ComparationData struct {
	Percentage            float64 `json:"percentage"`
	TingkatKeketatan      float64 `json:"tingkat_keketatan"`
	TingkatKeselarasan    float64 `json:"tingkat_keselarasan"`
	TingkatDapatPekerjaan float64 `json:"tingkat_dapat_pekerjaan"`
	TingkatProspekKerja   float64 `json:"tingkat_prospek_kerja"`
	Gaji                  string  `json:"gaji"`
	Pekerjaan             string  `json:"pekerjaan"`
	Deskripsi             string  `json:"deskripsi"`
	TagJurusan            string  `json:"tag_jurusan"`
	Jurusan               string  `json:"jurusan"`
	MataKuliah            string  `json:"mata_kuliah"`
}
