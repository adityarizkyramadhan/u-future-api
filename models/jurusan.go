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
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type JurusanStudentCompare struct {
}
