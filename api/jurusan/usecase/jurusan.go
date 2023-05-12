package usecase

import (
	"sort"
	repJur "u-future-api/api/jurusan/repository"
	repQ "u-future-api/api/quiz/repository"
	"u-future-api/models"
	"u-future-api/util/logic"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
)

type Jurusan struct {
	rj *repJur.Jurusan
	uc *repQ.Quiz
}

func New(rj *repJur.Jurusan, uc *repQ.Quiz) *Jurusan {
	return &Jurusan{rj, uc}
}

func (uj *Jurusan) getAnalisis(idUSer string) error {
	var analisis models.QuizResultAnalisis
	userId, err := uuid.FromString(idUSer)
	if err != nil {
		return err
	}
	history, err := uj.uc.GetHistoryByUserId(idUSer)
	if err != nil {
		return err
	}
	result, err := uj.uc.GetResultByUserID(idUSer)
	if err != nil {
		return err
	}
	userRaisecArr := logic.TrimAndChangeStringToArray(history.ResultSectionOne)

	analisis.Tag = logic.GetMostFrequentItems(userRaisecArr)
	analisis.Realistic = logic.CalculateResult("R", userRaisecArr)
	analisis.Investigative = logic.CalculateResult("I", userRaisecArr)
	analisis.Artistic = logic.CalculateResult("A", userRaisecArr)
	analisis.Social = logic.CalculateResult("S", userRaisecArr)
	analisis.Enterprising = logic.CalculateResult("E", userRaisecArr)
	analisis.Conventional = logic.CalculateResult("C", userRaisecArr)

	persentase := make(map[string]float64)
	persentase["R"] = analisis.Realistic
	persentase["I"] = analisis.Investigative
	persentase["A"] = analisis.Artistic
	persentase["S"] = analisis.Social
	persentase["E"] = analisis.Enterprising
	persentase["C"] = analisis.Conventional
	var top3 []string
	for i := 0; i < 3; i++ {
		maxPersentase := 0.0
		maxKepribadian := ""
		for k, v := range persentase {
			if v > maxPersentase {
				maxPersentase = v
				maxKepribadian = k
			}
		}
		top3 = append(top3, maxKepribadian)
		delete(persentase, maxKepribadian)
	}

	// Ubah array hasilnya menjadi string
	resultString := ""
	for i, k := range top3 {
		resultString += k
		if i < len(top3)-1 {
			resultString += ", "
		}
	}
	jurusans, err := uj.rj.FindAllJurusan()
	if err != nil {
		return err
	}
	for _, v := range jurusans {
		data := models.JurusanStudentCompare{
			ID:          uuid.Must(uuid.NewV4()),
			UserID:      userId,
			NamaJurusan: v.NamaJurusan,
			TagJurusan:  v.TagJurusan,
			Jurusan:     v.Jurusan,
			Percentage:  float64((logic.CompareJurusanAndHistory(v.RIASECCode, resultString) + logic.CompareJurusanAndHistory(v.WorkActivities, result.ResultSectionTwo) + logic.CompareJurusanAndHistory(v.Values, result.ResultSectionThree)/3.0)),
		}
		err = uj.rj.Create(&data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (uj *Jurusan) GetResult(idUser string) ([]*models.JurusanStudentCompare, error) {
	var jurusanCount int64
	if err := uj.rj.GetDB().Model(&models.JurusanStudentCompare{}).Where("user_id = ?", idUser).Count(&jurusanCount).Error; err != nil {
		return nil, err
	}
	if jurusanCount != 0 {
		return uj.rj.QueryByIDUser(idUser)
	}
	if err := uj.getAnalisis(idUser); err != nil {
		return nil, err
	}
	return uj.rj.QueryByIDUser(idUser)
}

func (uj *Jurusan) GetRekomendasi(idUser string) ([]*models.JurusanStudentCompare, error) {
	jurusan, err := uj.GetResult(idUser)
	if err != nil {
		return nil, err
	}
	sort.Slice(jurusan, func(i, j int) bool {
		return jurusan[i].Percentage > jurusan[j].Percentage
	})
	return jurusan[:3], err
}

func (uj *Jurusan) GetComparationData(name, idUser string) (*models.ComparationData, error) {
	var comparation models.ComparationData
	results, err := uj.rj.QueryByNameAndIdUSer(name, idUser)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&comparation, results)
	if err != nil {
		return nil, err
	}
	jurusan, err := uj.rj.GetByNameJurusan(results.NamaJurusan)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(&comparation, jurusan)
	if err != nil {
		return nil, err
	}
	return &comparation, nil
}
