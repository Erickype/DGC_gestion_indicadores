package model

import (
	"errors"
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ScienceMagazine struct {
	ID                 uint      `gorm:"primary_key"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Name               string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation       string    `gorm:"size:200;not null;unique" json:"abbreviation"`
	Description        string    `gorm:"size:200" json:"description,omitempty"`
	AcademicDatabaseID uint      `json:"academic_database_id"`

	AcademicDatabase AcademicDatabase `json:"academic_database"`
}

type FilterScienceMagazinesRequest struct {
	AcademicDatabase     string `json:"academic_database,omitempty"`
	ScienceMagazine      string `json:"science_magazine,omitempty"`
	MagazineAbbreviation string `json:"magazine_abbreviation,omitempty"`
	MagazineDescription  string `json:"magazine_description,omitempty"`
	PageSize             int    `json:"page_size"`
	Page                 int    `json:"page"`
}

type ScienceMagazinesJoined struct {
	AcademicDatabaseID          uint   `json:"academic_database_id"`
	AcademicDatabase            string `json:"academic_database"`
	ScienceMagazineID           uint   `json:"science_magazine_id"`
	ScienceMagazine             string `json:"science_magazine"`
	ScienceMagazineAbbreviation string `json:"science_magazine_abbreviation"`
	ScienceMagazineDescription  string `json:"science_magazine_description"`
}

type FilterScienceMagazinesResponse struct {
	Count            int64                    `json:"count"`
	PageSize         int                      `json:"page_size"`
	Page             int                      `json:"page"`
	ScienceMagazines []ScienceMagazinesJoined `json:"science_magazines"`
}

func GetScienceMagazines(magazines *[]ScienceMagazine) (err error) {
	err = database.DB.Order("updated_at desc").Find(magazines).Error
	if err != nil {
		return err
	}
	return nil
}

func GetScienceMagazineByID(id int, magazine *ScienceMagazine) (err error) {
	err = database.DB.First(magazine, id).Error
	if err != nil {
		return err
	}
	return nil
}

func PostScienceMagazine(magazine *ScienceMagazine) (err error) {
	err = database.DB.Create(magazine).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("revista científica ya existe")
		}
		return err
	}
	return nil
}

func PutScienceMagazine(magazine *ScienceMagazine) (err error) {
	err = database.DB.Save(magazine).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("programa posgrado ya registrado")
		}
		return err
	}
	return nil
}

func PostFilterScienceMagazines(
	filterScienceMagazinesResponse *FilterScienceMagazinesResponse,
	filterScienceMagazinesRequest *FilterScienceMagazinesRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterScienceMagazinesRequest.AcademicDatabase != "" {
		conditions = append(conditions, "LOWER(people.identity) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterScienceMagazinesRequest.AcademicDatabase)))
	}
	if filterScienceMagazinesRequest.ScienceMagazine != "" {
		conditions = append(conditions, "LOWER(people.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterScienceMagazinesRequest.ScienceMagazine)))
	}
	if filterScienceMagazinesRequest.MagazineAbbreviation != "" {
		conditions = append(conditions, "LOWER(people.lastname) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterScienceMagazinesRequest.MagazineAbbreviation)))
	}
	if filterScienceMagazinesRequest.MagazineDescription != "" {
		conditions = append(conditions, "LOWER(careers.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterScienceMagazinesRequest.MagazineDescription)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`ad.id as academic_database_id,
		ad.name as academic_database,
		sm.id as science_magazine_id,
		sm.name as science_magazine,
		sm.abbreviation as science_magazine_abbreviation,
		sm.description as science_magazine_description`,
	).Joins(
		`join academic_databases ad on sm.academic_database_id = ad.id`,
	)

	var totalCount int64
	err = query.Table("science_magazines sm").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterScienceMagazinesRequest.PageSize
	page := filterScienceMagazinesRequest.Page
	err = query.
		Order("sm.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterScienceMagazinesResponse.ScienceMagazines).Error
	if err != nil {
		return err
	}

	filterScienceMagazinesResponse.Count = totalCount
	filterScienceMagazinesResponse.PageSize = pageSize
	if page <= 0 {
		filterScienceMagazinesResponse.Page = 1
	} else {
		filterScienceMagazinesResponse.Page = page
	}
	return nil
}

func GetScienceMagazineFilterJoinedByScienceMagazineID(scienceMagazineID int, response *ScienceMagazinesJoined) (err error) {
	err = database.DB.Table("science_magazines sm").
		Select(`ad.id as academic_database_id,
		ad.name as academic_database,
		sm.id as science_magazine_id,
		sm.name as science_magazine,
		sm.abbreviation as science_magazine_abbreviation,
		sm.description as science_magazine_description`).
		Joins(`join academic_databases ad on sm.academic_database_id = ad.id`).
		Where("sm.id = ?", scienceMagazineID).
		Scan(&response).Error

	if err != nil {
		return err
	}
	return nil
}
