package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"gorm.io/gorm"
	"time"
)

type PostgraduateCohortList struct {
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	PostgraduateProgramID uint      `gorm:"primary_key" json:"postgraduate_program_id"`
	Year                  uint      `gorm:"primary_key" json:"year"`
	Name                  string    `json:"name"`
	GraduatedStudents     uint      `json:"graduated_students"`
	TotalStudents         uint      `json:"total_students"`

	PostgraduateProgram PostgraduateProgram `json:"-"`
}

func (c PostgraduateCohortList) TableName() string {
	return model.IndicatorsInformationSchema + ".postgraduate_cohort_lists"
}

func GetPostgraduateCohortListByProgramIDAndYear(programID, year int, postgraduateCohortList *PostgraduateCohortList) (err error) {
	err = database.DB.Model(&PostgraduateCohortList{}).
		Find(&postgraduateCohortList,
			"postgraduate_program_id = ? and year = ?", programID, year).
		Error
	if err != nil {
		return err
	}
	return nil
}

func GetPostgraduateCohortListsByProgramID(programID int, postgraduateCohortLists *[]PostgraduateCohortList) (err error) {
	err = database.DB.Model(&PostgraduateCohortList{}).
		Where("postgraduate_program_id = ?", programID).
		Find(&postgraduateCohortLists).Error
	if err != nil {
		return err
	}
	return nil
}

func PostPostgraduateCohortList(postgraduateCohortList *PostgraduateCohortList) (err error) {
	var postgraduateProgram PostgraduateProgram
	err = GetPostgraduateProgramByID(int(postgraduateCohortList.PostgraduateProgramID), &postgraduateProgram)
	if err != nil {
		return err
	}
	if postgraduateCohortList.Year < postgraduateProgram.StartYear ||
		postgraduateCohortList.Year > postgraduateProgram.EndYear {
		return errors.New("año fuera de intervalo de programa")
	}
	err = database.DB.Create(&postgraduateCohortList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("cohorte ya existe")
		}
		return err
	}
	return nil
}

func UpdatePostgraduateCohortList(postgraduateCohortList *PostgraduateCohortList) (err error) {
	err = database.DB.Save(postgraduateCohortList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("valores para la cohorte ya registrados")
		}
		return err
	}
	return nil
}
