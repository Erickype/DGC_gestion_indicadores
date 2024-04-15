package util

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	user "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
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
	err = database.DB.AutoMigrate(&academicPeriod.AcademicPeriod{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	err = database.DB.AutoMigrate(&person.Person{})
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
