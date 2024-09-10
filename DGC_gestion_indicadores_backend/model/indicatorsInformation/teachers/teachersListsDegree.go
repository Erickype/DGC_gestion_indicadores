package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	degree "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"gorm.io/gorm"
	"time"
)

type TeachersListsDegree struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	TeacherID        uint      `gorm:"primary_key" json:"teacher_id"`
	TeachersDegreeID uint      `gorm:"primary_key" json:"teachers_degree_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	Teacher        teacher.Teacher               `json:"-"`
	TeachersDegree degree.TeachersDegree         `json:"-"`
}

type AddDegreeAndTeachersListsDegreeRequest struct {
	AcademicPeriodID uint   `json:"academic_period_id"`
	TeacherID        uint   `json:"teacher_id"`
	DegreeLevelID    uint   `json:"degree_level_id"`
	Name             string `json:"name"`
}

type GetTeachersListsDegreesJoinedResponse struct {
	AcademicPeriodID uint   `json:"academic_period_id"`
	TeacherID        uint   `json:"teacher_id"`
	TeachersDegreeID uint   `json:"teachers_degree_id"`
	DegreeLevelID    uint   `json:"degree_level_id"`
	DegreeLevelName  string `json:"degree_level_name"`
	Name             string `json:"name"`
}

type GetDegreesNotAssignedResponse struct {
	TeachersDegreeID uint   `json:"teachers_degree_id"`
	Abbreviation     string `json:"abbreviation"`
	Name             string `json:"name"`
}

func (tl *TeachersListsDegree) TableName() string {
	return model.IndicatorsInformationSchema + ".teachers_lists_degrees"
}

func AddDegreeAndTeachersListsDegree(request *AddDegreeAndTeachersListsDegreeRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		teachersDegree := degree.TeachersDegree{
			DegreeLevelID: request.DegreeLevelID,
			TeacherID:     request.TeacherID,
			Name:          request.Name,
		}
		if err := tx.Create(&teachersDegree).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("título ya registrado")
			}
			return err
		}
		teachersListsDegree := TeachersListsDegree{
			AcademicPeriodID: request.AcademicPeriodID,
			TeacherID:        request.TeacherID,
			TeachersDegreeID: teachersDegree.ID,
		}
		if err := tx.Create(&teachersListsDegree).Error; err != nil {
			if errors.Is(err, gorm.ErrForeignKeyViolated) {
				return errors.New("claves no encontradas")
			}
			return err
		}
		return nil
	})
}

func GetDegreesNotAssigned(response *[]GetDegreesNotAssignedResponse, academicPeriodID, teacherID int) (err error) {
	err = database.DB.
		Table("teachers_degree td").
		Select(
			`tld.teachers_degree_id,
					dl.abbreviation,
					td.name`).
		Joins(`join indicators_information.teachers_lists_degrees tld on td.id = tld.teachers_degree_id`).
		Joins(`degree_levels dl on td.degree_level_id = dl.id`).
		Where("tdl.academic_period_id = ? and tdl.teacher_id = ? ", academicPeriodID, teacherID).
		Scan(&response).Error

	if err != nil {
		return err
	}
	return nil
}

func GetTeachersListsDegreesJoined(
	academicPeriodID, teacherID int,
	response *[]GetTeachersListsDegreesJoinedResponse) (err error) {
	err = database.DB.Order("tld.updated_at desc").
		Table("indicators_information.teachers_lists_degrees tld").
		Select(`tld.academic_period_id, 
					tld.teacher_id, 
					tld.teachers_degree_id, 
					td.degree_level_id, 
					dl.name as degree_level_name, 
					td.name`).
		Joins("join teachers_degrees td on tld.teachers_degree_id = td.id").
		Joins("join degree_levels dl on dl.id = td.degree_level_id").
		Where("tld.academic_period_id = ? AND tld.teacher_id = ?", academicPeriodID, teacherID).
		Scan(&response).Error
	if err != nil {
		return err
	}
	return nil
}
