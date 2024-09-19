package util

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	modelConsts "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	author "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	user "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	contractType "github.com/Erickype/DGC_gestion_indicadores_backend/model/contractType"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/model/dedication"
	degree "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	evaluationAcademicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	indicators "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicators"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation"
	indicatorsInformationAcademicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"log"
	"os"
)

func LoadDatabase() {
	database.InitDb()

	database.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + modelConsts.IndicatorsInformationSchema)
	database.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + modelConsts.IndicatorsSchema)

	err := database.DB.AutoMigrate(
		&user.Role{},
		&user.User{},
		&evaluationPeriod.EvaluationPeriod{},
		&academicPeriod.AcademicPeriod{},
		&evaluationAcademicPeriod.EvaluationAcademicPeriod{},
		&person.Person{},
		&faculty.Faculty{},
		&career.Career{},
		&dedication.Dedication{},
		&contractType.ContractType{},
		&scaledGrade.ScaledGrade{},
		&degree.DegreeLevel{},
		&degree.TeachersDegree{},
		&model.TeachersListsDegree{},
		&teacher.Teacher{},
		&model.TeachersList{},
		&academicProduction.AcademicDatabase{},
		&academicProduction.ScienceMagazine{},
		&academicProduction.CompensationFactor{},
		&academicProduction.ImpactCoefficient{},
		&author.Author{},
		&academicProduction.PublicationType{},
		&indicatorsInformationAcademicProduction.AcademicProductionList{},
		&indicatorsInformationAcademicProduction.AcademicProductionListsAuthor{},
		&indicatorsInformation.BooksOrChaptersProductionList{},
		&indicatorsInformation.BooksOrChaptersProductionListAuthor{},
		&indicators.IndicatorType{},
		&indicators.IndicatorsEvaluationPeriod{},
		&indicators.IndicatorsAcademicPeriod{},
	)
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []user.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "upe", Description: "Authenticated UPE role"},
		{Name: "authority", Description: "Authenticated authority role"},
	}
	var adminUser = []user.User{
		{Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
			RoleID:   1},
	}
	var dedications = []dedication.Dedication{
		{Name: "Tiempo completo", Abbreviation: "TC", Description: "Personal 8 horas diarias"},
		{Name: "Medio tiempo", Abbreviation: "MT", Description: "Personal 4 horas diarias"},
		{Name: "Tiempo parcial", Abbreviation: "TP", Description: "Personal horas variadas"},
		{Name: "Tiempo completo 10 horas permiso", Abbreviation: "TCP", Description: "Personal 8 horas diarias y horas de permiso"},
	}
	var scaledGrades = []scaledGrade.ScaledGrade{
		{Name: "Profesor titular agregado 1", Abbreviation: "Titular agregado 1", Description: "Profesor agregado"},
		{Name: "Profesor titular agregado 2", Abbreviation: "Titular agregado 2", Description: "Profesor agregado"},
		{Name: "Profesor titular agregado 3", Abbreviation: "Titular agregado 3", Description: "Profesor agregado"},
		{Name: "Profesor titular auxiliar 1", Abbreviation: "Titular auxiliar 1", Description: "Profesor auxiliar"},
		{Name: "Profesor titular auxiliar 2", Abbreviation: "Titular auxiliar 2", Description: "Profesor auxiliar"},
	}
	var contractTypes = []contractType.ContractType{
		{Name: "Docente a nombramiento", Abbreviation: "Nombramiento", Description: "Docente nombramiento"},
		{Name: "Docente a contrato ocasional", Abbreviation: "Contrato ocasional", Description: "Docente contrato ocasional"},
		{Name: "Docente de apoyo académico", Abbreviation: "Apoyo académico", Description: "Docente de apoyo académico"},
	}

	var degreeLevels = []degree.DegreeLevel{
		{Name: "Tercer nivel", Abbreviation: "3rd Nivel", Description: "Títulos universitarios tercer nivel"},
		{Name: "Cuarto nivel Posgrado/Master", Abbreviation: "4to Nivel", Description: "Títulos posgrado/master cuarto nivel"},
		{Name: "Cuarto nivel Doctorado", Abbreviation: "PhD", Description: "Títulos doctorado cuarto nivel"},
	}

	var indicatorTypes = []indicators.IndicatorType{
		{
			ID:           16,
			Name:         "Indicador 16: Personal académico con formacion doctoral",
			Abbreviation: "Tasa personal académico con formacion doctoral",
			Description:  "Tasa de formación doctoral de al menos el 20%",
		},
		{
			ID:           17,
			Name:         "Indicador 17: Personal Académico con Dedicación a Tiempo Completo",
			Abbreviation: "Tasa personal académico con dedicación tiempo completo",
			Description:  "Tasa del personal académico con dedicación a tiempo completo al menos el 50%",
		},
	}

	var compensationFactors = []academicProduction.CompensationFactor{
		{ID: 1, Name: "Cuartil 1", Abbreviation: "Q1", Description: "Factor para cualtil 1", Weight: 1},
		{ID: 2, Name: "Cuartil 2", Abbreviation: "Q2", Description: "Factor para cualtil 2", Weight: 0.9},
		{ID: 3, Name: "Cuartil 3", Abbreviation: "Q3", Description: "Factor para cualtil 3", Weight: 0.8},
		{ID: 4, Name: "Cuartil 4", Abbreviation: "Q4", Description: "Factor para cualtil 4", Weight: 0.7},
		{ID: 5, Name: "ACI", Abbreviation: "ACI", Description: "Factor para ACI", Weight: 0.6},
		{ID: 6, Name: "BR", Abbreviation: "BR", Description: "Factor para BR", Weight: 0.5},
		{ID: 7, Name: "LA", Abbreviation: "LA", Description: "Factor para LA", Weight: 0.2},
	}

	var academicDatabases = []academicProduction.AcademicDatabase{
		{ID: 1, Name: "SCOPUS", Description: "Base datos principal", Abbreviation: "SCOPUS"},
		{ID: 2, Name: "WebOfScience", Description: "Base datos WebOfScience", Abbreviation: "WebOfScience"},
		{ID: 3, Name: "Regionales", Description: "Base datos Regionales", Abbreviation: "Regionales"},
		{ID: 4, Name: "Anexo1", Description: "Base datos Anexo1", Abbreviation: "Anexo1"},
		{ID: 5, Name: "Latindex", Description: "Base datos Latindex", Abbreviation: "Latindex"},
	}

	var impactCoefficients = []academicProduction.ImpactCoefficient{
		{ID: 1, AcademicDatabaseID: 1, CompensationFactorID: 1},
		{ID: 2, AcademicDatabaseID: 1, CompensationFactorID: 2},
		{ID: 3, AcademicDatabaseID: 1, CompensationFactorID: 3},
		{ID: 4, AcademicDatabaseID: 1, CompensationFactorID: 4},
		{ID: 5, AcademicDatabaseID: 1, CompensationFactorID: 5},
		{ID: 6, AcademicDatabaseID: 2, CompensationFactorID: 5},
		{ID: 7, AcademicDatabaseID: 3, CompensationFactorID: 6},
		{ID: 8, AcademicDatabaseID: 4, CompensationFactorID: 6},
		{ID: 9, AcademicDatabaseID: 5, CompensationFactorID: 7},
	}

	database.DB.Save(&academicDatabases)
	database.DB.Save(&compensationFactors)
	database.DB.Save(&impactCoefficients)
	database.DB.Save(&indicatorTypes)
	database.DB.Save(&degreeLevels)
	database.DB.Save(&contractTypes)
	database.DB.Save(&scaledGrades)
	database.DB.Save(&dedications)
	database.DB.Save(&roles)
	database.DB.Save(&adminUser)
}
