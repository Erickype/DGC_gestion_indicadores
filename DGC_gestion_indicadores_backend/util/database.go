package util

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	modelConsts "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	user "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	contractType "github.com/Erickype/DGC_gestion_indicadores_backend/model/contractType"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/model/dedication"
	degree "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	evaluationAcademicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"log"
	"os"
)

func LoadDatabase() {
	database.InitDb()

	database.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + modelConsts.IndicatorsInformationSchema)

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
		&degree.TeachersDegreesAcademicPeriod{},
		&teacher.Teacher{},
		&indicatorsInformation.TeachersList{},
		&academicProduction.AcademicDatabase{},
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
	database.DB.Save(&contractTypes)
	database.DB.Save(&scaledGrades)
	database.DB.Save(&dedications)
	database.DB.Save(&roles)
	database.DB.Save(&adminUser)
}
