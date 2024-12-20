package model

import (
	"errors"
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	evaluationAcademicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	production "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/artisticProduction"
	"gorm.io/datatypes"
)

func CalculateIndicator26(evaluationPeriodID int, indicator *IndicatorsEvaluationPeriod) (err error) {
	var academicPeriodIds []int
	err = getAcademicPeriodsInEvaluationPeriod(evaluationPeriodID, &academicPeriodIds)
	if err != nil {
		return err
	}

	var selectedEvaluationPeriod evaluationPeriod.EvaluationPeriod
	err = getCurrentEvaluationPeriod(evaluationPeriodID, &selectedEvaluationPeriod)
	if err != nil {
		return err
	}

	var countFullTimeTeachers int64
	err = getFullTimeTeachers(academicPeriodIds, &countFullTimeTeachers)
	if err != nil {
		return err
	}

	var countPartTimeTeachers int64
	err = getPartTimeTeachers(academicPeriodIds, &countPartTimeTeachers)
	if err != nil {
		return err
	}

	var academicPublication float64
	err = calculateIndicator26AcademicPublication(selectedEvaluationPeriod.StartYear, selectedEvaluationPeriod.EndYear, &academicPublication)
	if err != nil {
		return err
	}

	var artisticProduction float64
	err = calculateIndicator26ArtisticProduction(academicPeriodIds, &artisticProduction)
	if err != nil {
		return err
	}

	var booksOrChaptersPublication float64
	err = calculateIndicator26BooksOrChapterPublication(
		selectedEvaluationPeriod.StartYear, selectedEvaluationPeriod.EndYear, &booksOrChaptersPublication)
	if err != nil {
		return err
	}

	var intellectualProperty float64
	err = calculateIndicator26IntellectualProperty(academicPeriodIds, &intellectualProperty)
	if err != nil {
		return err
	}

	indicator.ActualValue =
		(academicPublication + artisticProduction + booksOrChaptersPublication + intellectualProperty) /
			(float64(countFullTimeTeachers) + 0.5*float64(countPartTimeTeachers))
	indicator.TargetValue = model.Indicator26TargetValue
	err = RefreshIndicatorEvaluationPeriod(indicator)
	if err != nil {
		return err
	}
	return nil
}

func calculateIndicator26AcademicPublication(startYear, endYear datatypes.Date, academicPublication *float64) (err error) {
	var weights []struct {
		Weight                 float64 `json:"weight"`
		InterculturalComponent bool    `json:"intercultural_component"`
	}
	err = database.DB.Table("indicators_information.academic_production_lists apl").
		Joins("join impact_coefficients ic on apl.impact_coefficient_id = ic.id").
		Joins("join compensation_factors cf on ic.compensation_factor_id = cf.id").
		Where("apl.publication_date between ? and ?", startYear, endYear).
		Scan(&weights).Error

	*academicPublication = 0
	for _, weight := range weights {
		var weightedValue float64 = 0
		if weight.InterculturalComponent {
			weightedValue = weight.Weight + 0.21
			if weightedValue > 1 {
				weightedValue = 1
			}
		} else {
			weightedValue = weight.Weight
		}
		*academicPublication += weightedValue
	}
	return nil
}

func calculateIndicator26ArtisticProduction(
	academicPeriodIds []int, artisticProduction *float64) (err error) {
	var valuesSums production.ArtisticProductionList
	err = database.DB.Table("indicators_information.artistic_production_lists apl").
		Select(`sum(apl.international_artistic_work) as international_artistic_work,
			sum(apl.national_artistic_work) as national_artistic_work`).
		Where("apl.academic_period_id in ?", academicPeriodIds).
		Scan(&valuesSums).Error
	if err != nil {
		return err
	}
	*artisticProduction = float64(valuesSums.InternationalArtisticWork) + 0.9*float64(valuesSums.NationalArtisticWork)
	return nil
}

func calculateIndicator26BooksOrChapterPublication(
	startYear, endYear datatypes.Date, booksOrChaptersPublication *float64) (err error) {
	var countPeerReviewedBooks int64
	err = database.DB.Table("indicators_information.books_or_chapter_production_lists bocpl").
		Where("bocpl.publication_date between ? and ?", startYear, endYear).
		Where("bocpl.is_chapter = false").
		Where("bocpl.peer_reviewed").
		Count(&countPeerReviewedBooks).Error
	if err != nil {
		return err
	}

	var countChapters int64
	err = database.DB.Table("indicators_information.books_or_chapter_production_lists bocpl").
		Where("bocpl.publication_date between ? and ?", startYear, endYear).
		Where("bocpl.is_chapter").
		Count(&countChapters).Error
	if err != nil {
		return err
	}

	var countPeerReviewedChapters int64
	err = database.DB.Table("indicators_information.books_or_chapter_production_lists bocpl").
		Where("bocpl.publication_date between ? and ?", startYear, endYear).
		Where("bocpl.is_chapter").
		Where("bocpl.peer_reviewed").
		Count(&countPeerReviewedChapters).Error
	if err != nil {
		return err
	}
	if countPeerReviewedChapters == 0 {
		return errors.New(fmt.Sprintf("indicador %d: no hay capítulos revisados pares", 26))
	}

	*booksOrChaptersPublication =
		float64(countPeerReviewedBooks) +
			float64(countChapters)/float64(countPeerReviewedChapters)
	return nil
}

func calculateIndicator26IntellectualProperty(
	academicPeriodIds []int, intellectualProperty *float64) (err error) {
	var valuesSums production.ArtisticProductionList
	err = database.DB.Table("indicators_information.artistic_production_lists apl").
		Select(`sum(apl.intellectual_property) as intellectual_property`).
		Where("apl.academic_period_id in ?", academicPeriodIds).
		Scan(&valuesSums).Error
	if err != nil {
		return err
	}
	*intellectualProperty = float64(valuesSums.IntellectualProperty)
	return nil
}

func getFullTimeTeachers(academicPeriodIds []int, countFullTimeTeachers *int64) (err error) {
	err = database.DB.Table("indicators_information.teachers_lists tl").
		Joins("join dedications d on tl.dedication_id = d.id").
		Where("academic_period_id in ?", academicPeriodIds).
		Where("d.name = ?", model.FullTimeDedication).
		Distinct("teacher_id").Count(countFullTimeTeachers).Error
	if err != nil {
		return err
	}
	if *countFullTimeTeachers <= 0 {
		return errors.New(fmt.Sprintf("indicador %d: no hay profesores a tiempo completo", 26))
	}
	return nil
}

func getPartTimeTeachers(academicPeriodIds []int, countFullTimeTeachers *int64) (err error) {
	err = database.DB.Table("indicators_information.teachers_lists tl").
		Joins("join dedications d on tl.dedication_id = d.id").
		Where("academic_period_id in ?", academicPeriodIds).
		Where("d.name = ?", model.PartTimeDedication).
		Distinct("teacher_id").Count(countFullTimeTeachers).Error
	if err != nil {
		return err
	}
	if *countFullTimeTeachers <= 0 {
		return errors.New(fmt.Sprintf("indicador %d: no hay profesores a tiempo parcial", 26))
	}
	return nil
}

func getCurrentEvaluationPeriod(evaluationPeriodID int, currentEvaluationPeriod *evaluationPeriod.EvaluationPeriod) (err error) {
	err = database.DB.Model(&evaluationPeriod.EvaluationPeriod{}).
		Select(`start_year,
    				end_year`).
		Where("id = ?", evaluationPeriodID).
		Scan(&currentEvaluationPeriod).Error
	if err != nil {
		return err
	}
	return nil
}

func getAcademicPeriodsInEvaluationPeriod(evaluationPeriodID int, academicPeriodIDs *[]int) (err error) {
	err = database.DB.Model(&evaluationAcademicPeriod.EvaluationAcademicPeriod{}).
		Select("academic_period_id").
		Where("evaluation_period_id = ?", evaluationPeriodID).
		Scan(&academicPeriodIDs).Error
	if err != nil {
		return err
	}
	return nil
}
