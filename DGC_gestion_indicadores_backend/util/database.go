package util

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	user "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/model/dedication"
	evaluationAcademicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"log"
	"os"
)

func LoadDatabase() {
	database.InitDb()
	err := database.DB.AutoMigrate(&user.Role{})
	if err != nil {
		log.Fatal("Error while migration: ", err.Error())
	}
	err = database.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("Error while migration: ", err.Error())
	}
	err = database.DB.AutoMigrate(&evaluationPeriod.EvaluationPeriod{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&academicPeriod.AcademicPeriod{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&evaluationAcademicPeriod.EvaluationAcademicPeriod{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&person.Person{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&faculty.Faculty{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&career.Career{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&dedication.Dedication{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&scaledGrade.ScaledGrade{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&teacher.Teacher{})
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
	database.DB.Save(&roles)
	database.DB.Save(&adminUser)
}
