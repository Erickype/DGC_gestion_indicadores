package main

import (
	"fmt"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/controller/academicPeriod"
	controller "github.com/Erickype/DGC_gestion_indicadores_backend/controller/auth"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	model2 "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	"github.com/Erickype/DGC_gestion_indicadores_backend/util"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() {
	database.InitDb()
	err := database.DB.AutoMigrate(&model.Role{})
	if err != nil {
		log.Fatal("Error while migration: ", err.Error())
	}
	err = database.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Error while migration: ", err.Error())
	}
	err = database.DB.AutoMigrate(&model2.AcademicPeriod{})
	if err != nil {
		log.Fatal("Error while migrating: ", err.Error())
	}
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []model.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "upe", Description: "Authenticated UPE role"},
		{Name: "authority", Description: "Authenticated authority role"},
	}
	var user = []model.User{
		{Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
			RoleID:   1},
	}
	database.DB.Save(&roles)
	database.DB.Save(&user)
}

func serveApplication() {
	router := gin.Default()

	authRoutes := router.Group("/auth/user")
	authRoutes.POST("/register", controller.Register)
	authRoutes.POST("/login", controller.Login)

	// Admin routes
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(util.JWTAuth())
	adminRoutes.GET("/users", controller.GetUsers)
	adminRoutes.GET("/user/:id", controller.GetUser)
	adminRoutes.PUT("/user/:id", controller.UpdateUser)
	adminRoutes.POST("/user/role", controller.CreateRole)
	adminRoutes.GET("/user/roles", controller.GetRoles)
	adminRoutes.PUT("/user/role/:id", controller.UpdateRole)

	adminRoutes.POST("/academicPeriod", academicPeriod.CreateAcademicPeriod)

	// Public view routes
	academicPeriodRoutes := router.Group("/view")
	academicPeriodRoutes.GET("/academicPeriods", academicPeriod.GetAcademicPeriods)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Error running server")
	}
	fmt.Println("Server running on port 8000")
}
