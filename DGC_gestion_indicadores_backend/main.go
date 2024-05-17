package main

import (
	"fmt"
	"log"

	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/controller/evaluationPeriod"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/controller/academicPeriod"
	controller "github.com/Erickype/DGC_gestion_indicadores_backend/controller/auth"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/controller/career"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/controller/dedication"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/controller/faculty"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/controller/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/controller/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/controller/teacher"
	"github.com/Erickype/DGC_gestion_indicadores_backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	util.LoadDatabase()
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

func serveApplication() {
	router := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

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

	adminRoutes.POST("/person", person.CreatePerson)

	adminRoutes.POST("/faculty", faculty.CreateFaculty)

	adminRoutes.POST("/career", career.CreateCareer)

	adminRoutes.POST("/dedication", dedication.CreateDedication)

	adminRoutes.POST("/scaledGrade", scaledGrade.CreateScaledGrade)

	adminRoutes.POST("/evaluationPeriod", evaluationPeriod.CreateEvaluationPeriod)

	// UPE routes
	upeRoutes := router.Group("/api")
	upeRoutes.Use(util.JWTAuth(), util.JWTAuthUPE())
	upeRoutes.GET("/people", person.GetPersons)

	upeRoutes.GET("/faculties", faculty.GetFaculties)

	upeRoutes.GET("/careers", career.GetFaculties)

	upeRoutes.GET("/dedications", dedication.GetDedications)

	upeRoutes.GET("/scaledGrades", scaledGrade.GetScaledGrades)

	upeRoutes.GET("/teachers/byAcademicPeriod/:academicPeriodID", teacher.GetTeachersByAcademicPeriod)
	upeRoutes.POST("/teacher", teacher.CreateTeacher)
	upeRoutes.DELETE("/teacher/:id", teacher.DeleteTeacher)
	upeRoutes.PUT("/teacher/:id", teacher.UpdateTeacher)

	// Public view routes
	academicPeriodRoutes := router.Group("/view")
	academicPeriodRoutes.GET("/academicPeriods", academicPeriod.GetAcademicPeriods)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Error running server")
	}
	fmt.Println("Server running on port 8000")
}
