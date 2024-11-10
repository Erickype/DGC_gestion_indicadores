package main

import (
	"fmt"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/controller/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/controller/academicProduction"
	author "github.com/Erickype/DGC_gestion_indicadores_backend/controller/academicProduction/author"
	controller "github.com/Erickype/DGC_gestion_indicadores_backend/controller/auth"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/controller/career"
	contractTypes "github.com/Erickype/DGC_gestion_indicadores_backend/controller/contractType"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/controller/dedication"
	degree "github.com/Erickype/DGC_gestion_indicadores_backend/controller/degree"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/controller/evaluationPeriod"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/controller/faculty"
	indicators "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicators"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation"
	indicatorsInformationAcademicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/academicProduction"
	indicatorsInformationArtisticProduction "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/artisticProduction"
	indicatorsInformationBooksOrChaptersProduction "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/booksOrChaptersProduction"
	indicatorsInformationGradeRate "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/gradeRate"
	indicatorsInformationPostgraduate "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/postgraduate"
	indicatorsInformationResearchInnovationProject "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/researchInnovationProject"
	indicatorsInformationSocialProject "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/socialProjectProduction"
	indicatorsInformationTeachers "github.com/Erickype/DGC_gestion_indicadores_backend/controller/indicatorsInformation/teachers"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/controller/knowledgeLevel"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/controller/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/controller/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/controller/teacher"
	"github.com/Erickype/DGC_gestion_indicadores_backend/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	healthcheck "github.com/tavsec/gin-healthcheck"
	healthcheckChecks "github.com/tavsec/gin-healthcheck/checks"
	healthcheckConfig "github.com/tavsec/gin-healthcheck/config"
	"log"
	"os"
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
	appEnv := os.Getenv("APP_ENV")

	// Skip loading .env file if running in production
	if appEnv != "production" {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Println(".env file loaded successfully")
	} else {
		log.Println("Running in production mode, .env file not loaded")
	}
}

func serveApplication() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

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
	adminRoutes.PUT("/academicPeriod/:id", academicPeriod.UpdateAcademicPeriod)
	adminRoutes.DELETE("/academicPeriod/:id", academicPeriod.DeleteAcademicPeriod)

	adminRoutes.POST("/faculty", faculty.CreateFaculty)
	adminRoutes.PUT("/faculty/:id", faculty.UpdateFaculty)
	adminRoutes.DELETE("/faculty/:id", faculty.DeleteFaculty)

	adminRoutes.POST("/career", career.CreateCareer)
	adminRoutes.PUT("/career/:id", career.UpdateCareer)

	adminRoutes.POST("/dedication", dedication.CreateDedication)

	adminRoutes.POST("/scaledGrade", scaledGrade.CreateScaledGrade)

	adminRoutes.POST("/evaluationPeriod", evaluationPeriod.CreateEvaluationPeriod)
	adminRoutes.PUT("/evaluationPeriod/:id", evaluationPeriod.UpdateEvaluationPeriod)
	adminRoutes.DELETE("/evaluationPeriod/:id", evaluationPeriod.DeleteEvaluationPeriod)

	// UPE routes
	upeRoutes := router.Group("/api")
	upeRoutes.Use(util.JWTAuthUPE())
	upeRoutes.GET("/people", person.GetPersons)
	upeRoutes.POST("/people/filter", person.FilterPeople)
	upeRoutes.POST("/person", person.CreatePerson)
	upeRoutes.PUT("/person/:id", person.UpdatePerson)
	upeRoutes.DELETE("/person/:id", person.DeletePerson)

	upeRoutes.POST("/person/withRoles", person.PostPersonWithRoles)
	upeRoutes.PUT("/person/withRoles", person.UpdatePersonWithRoles)

	upeRoutes.GET("/faculties", faculty.GetFaculties)

	upeRoutes.GET("/careers", career.GetCareers)

	upeRoutes.GET("/dedications", dedication.GetDedications)

	upeRoutes.GET("/contractTypes", contractTypes.GetContractTypes)

	upeRoutes.GET("/scaledGrades", scaledGrade.GetScaledGrades)

	upeRoutes.GET("/scienceMagazines", academicProduction.GetScienceMagazines)

	upeRoutes.GET("/impactCoefficients", academicProduction.GetImpactCoefficients)

	upeRoutes.POST("/teachers/filter", teacher.FilterTeachers)
	upeRoutes.GET("/teachers/filter/teacherPersonJoinedByTeacherID/:id", teacher.GetTeacherPersonJoinedByTeacherID)
	upeRoutes.POST("/teacher", teacher.CreateTeacher)
	upeRoutes.DELETE("/teacher/:id", teacher.DeleteTeacher)
	upeRoutes.PUT("/teacher/:id", teacher.UpdateTeacher)

	upeRoutes.POST("/teachers/degree", degree.PostTeachersDegree)
	upeRoutes.PATCH("/teachers/degree/:id", degree.PatchTeachersDegreeByTeachersDegreeID)

	upeRoutes.GET("/degreeLevels", degree.GetDegreeLevels)

	upeRoutes.POST("/author", author.PostAuthor)
	upeRoutes.GET("/author/joined/:id", author.GetAuthorPersonJoinedByAuthorID)
	upeRoutes.POST("/authors/filter", author.FilterAuthors)

	upeRoutes.POST("/detailedFields/filter", knowledgeField.FilterDetailedFields)
	upeRoutes.GET("/detailedField/joined/:id", knowledgeField.GetDetailedFilterJoinedByDetailedFieldID)

	// Indicators information routes
	indicatorsInformationRoutes := router.Group("/api/indicators/information")
	indicatorsInformationRoutes.Use(util.JWTAuthUPE())

	indicatorsInformationRoutes.GET("/academicPeriodAuthorPreviousCareers/:authorID",
		indicatorsInformation.GetAcademicPeriodAuthorPreviousCareers)

	indicatorsInformationRoutes.POST("/teachersLists/degree", indicatorsInformationTeachers.PostTeachersListsDegree)
	indicatorsInformationRoutes.POST("/teachersLists/filter", indicatorsInformationTeachers.FilterTeachersLists)
	indicatorsInformationRoutes.POST("/teachersLists/AddDegreeAndTeachersListsDegree", indicatorsInformationTeachers.AddDegreeAndTeachersListsDegree)
	indicatorsInformationRoutes.GET("/teachersLists/degrees/:academicPeriodID/:teacherID", indicatorsInformationTeachers.GetTeachersListsDegreesJoined)
	indicatorsInformationRoutes.GET("/teachersLists/degrees/notAssigned/:academicPeriodID/:teacherID", indicatorsInformationTeachers.GetDegreesNotAssigned)

	indicatorsInformationRoutes.POST("/teachersList", indicatorsInformationTeachers.CreateTeachersList)
	indicatorsInformationRoutes.PATCH("/teachersList/:academicPeriodID/:teacherID", indicatorsInformationTeachers.PatchTeachersList)

	indicatorsInformationRoutes.POST("/academicProductionLists/filter",
		indicatorsInformationAcademicProduction.FilterAcademicProductionListsByAcademicPeriod)
	indicatorsInformationRoutes.POST("/academicProductionList", indicatorsInformationAcademicProduction.PostAcademicProductionList)
	indicatorsInformationRoutes.PATCH("/academicProductionList/:id", indicatorsInformationAcademicProduction.PatchAcademicProductionList)

	indicatorsInformationRoutes.POST("/academicProductionListsAuthor/authorCareers",
		indicatorsInformationAcademicProduction.PostAcademicProductionListsAuthorCareers)
	indicatorsInformationRoutes.PUT("/academicProductionListsAuthor/authorCareers",
		indicatorsInformationAcademicProduction.UpdateAcademicProductionListsAuthorCareersByAcademicPeriodID)
	indicatorsInformationRoutes.GET("/academicProductionListsAuthors/JoinedByAcademicProductionListID/:academicProductionListID",
		indicatorsInformationAcademicProduction.GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID)

	indicatorsInformationRoutes.POST("/booksOrChaptersProductionList",
		indicatorsInformationBooksOrChaptersProduction.PostBooksOrChaptersProductionList)
	indicatorsInformationRoutes.POST("/booksOrChaptersProductionLists/filter",
		indicatorsInformationBooksOrChaptersProduction.FilterBooksOrChaptersProductionListsByAcademicPeriod)
	indicatorsInformationRoutes.PUT("/booksOrChaptersProductionList/:id",
		indicatorsInformationBooksOrChaptersProduction.UpdateBooksOrChaptersProductionList)

	indicatorsInformationRoutes.POST("/booksOrChaptersProductionListsAuthor",
		indicatorsInformationBooksOrChaptersProduction.PostBooksOrChaptersProductionListsAuthorCareers)
	indicatorsInformationRoutes.PUT("/booksOrChaptersProductionListsAuthor/authorCareers",
		indicatorsInformationBooksOrChaptersProduction.UpdateBooksOrChaptersProductionListsAuthorCareers)
	indicatorsInformationRoutes.GET(
		"/booksOrChaptersProductionListsAuthors/JoinedByBooksOrChaptersProductionListID/:booksOrChaptersProductionListID",
		indicatorsInformationBooksOrChaptersProduction.GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID)

	indicatorsInformationRoutes.POST("/socialProjectList", indicatorsInformationSocialProject.PostSocialProjectList)
	indicatorsInformationRoutes.PUT("/socialProjectList/:id", indicatorsInformationSocialProject.PutSocialProjectList)
	indicatorsInformationRoutes.POST("/socialProjectLists/filter",
		indicatorsInformationSocialProject.FilterSocialProjectListsByAcademicPeriod)

	indicatorsInformationRoutes.GET("/gradeRateLists/:id", indicatorsInformationGradeRate.GetGradeRateListsByAcademicPeriod)
	indicatorsInformationRoutes.POST("/gradeRateList", indicatorsInformationGradeRate.PostGradeRateList)
	indicatorsInformationRoutes.PUT("/gradeRateList/:academicPeriodID/:careerID",
		indicatorsInformationGradeRate.UpdateGradeRateList)

	indicatorsInformationRoutes.GET("/researchInnovationProjectLists",
		indicatorsInformationResearchInnovationProject.GetResearchInnovationProjectListsByAcademicPeriod)
	indicatorsInformationRoutes.POST("/researchInnovationProjectList",
		indicatorsInformationResearchInnovationProject.PostResearchInnovationProjectList)
	indicatorsInformationRoutes.PUT("/researchInnovationProjectList/:academicPeriodID",
		indicatorsInformationResearchInnovationProject.UpdateResearchInnovationProjectList)

	indicatorsInformationRoutes.GET("/artisticProductionLists",
		indicatorsInformationArtisticProduction.GetArtisticProductionLists)
	indicatorsInformationRoutes.POST("/artisticProductionList",
		indicatorsInformationArtisticProduction.PostArtisticProductionList)
	indicatorsInformationRoutes.PUT("/artisticProductionList/:academicPeriodID",
		indicatorsInformationArtisticProduction.UpdateArtisticProjectList)

	indicatorsInformationRoutes.POST("/postgraduate/programs/filter",
		indicatorsInformationPostgraduate.FilterPostGraduatePrograms)
	indicatorsInformationRoutes.GET("/postgraduate/program/:programID",
		indicatorsInformationPostgraduate.GetPostgraduateProgramByID)
	indicatorsInformationRoutes.GET("/postgraduate/program/cohort/missingYears/:programID",
		indicatorsInformationPostgraduate.GetPostgraduateProgramMissingCohortYearsByProgramID)
	indicatorsInformationRoutes.POST("/postgraduate/program",
		indicatorsInformationPostgraduate.PostPostgraduateProgram)
	indicatorsInformationRoutes.PUT("/postgraduate/program/:programID",
		indicatorsInformationPostgraduate.UpdatePostGraduateProgram)

	indicatorsInformationRoutes.GET("/postgraduate/cohortLists/:programID",
		indicatorsInformationPostgraduate.GetPostgraduateCohortListsByProgramID)
	indicatorsInformationRoutes.POST("/postgraduate/cohortList",
		indicatorsInformationPostgraduate.PostPostgraduateCohortList)
	indicatorsInformationRoutes.PUT("/postgraduate/cohortList/:programID/:year",
		indicatorsInformationPostgraduate.UpdatePostgraduateCohortList)

	// Indicators routes
	indicatorsRoutes := router.Group("/api/indicators/")

	indicatorsRoutes.GET("/evaluationPeriod/calculate/:evaluationPeriodID/:indicatorTypeID",
		indicators.GetCalculateIndicatorByTypeIDAndEvaluationPeriod)
	indicatorsRoutes.GET("/evaluationPeriod/:evaluationPeriodID/:indicatorTypeID",
		indicators.GetIndicatorByTypeIDAndEvaluationPeriodJoined)
	indicatorsRoutes.GET("/evaluationPeriod/calculate/:evaluationPeriodID", indicators.GetCalculateIndicatorsByEvaluationPeriod)
	indicatorsRoutes.GET("/evaluationPeriod/:evaluationPeriodID", indicators.GetIndicatorsByEvaluationPeriod)

	indicatorsRoutes.GET("/academicPeriod/calculate/:academicPeriodID/:indicatorTypeID",
		indicators.GetCalculateIndicatorByTypeIDAndAcademicPeriod)
	indicatorsRoutes.GET("/academicPeriod/calculate/:academicPeriodID", indicators.GetCalculateIndicatorsByAcademicPeriod)
	indicatorsRoutes.GET("/academicPeriod/:academicPeriodID", indicators.GetIndicatorsByAcademicPeriod)

	// Public view routes
	academicPeriodRoutes := router.Group("/view")

	academicPeriodRoutes.GET("/academicPeriods", academicPeriod.GetAcademicPeriods)
	academicPeriodRoutes.POST("/academicPeriods/filter", academicPeriod.FilterAcademicPeriods)
	academicPeriodRoutes.GET("/academicPeriod/:id", academicPeriod.GetAcademicPeriodByID)

	academicPeriodRoutes.GET("/evaluationPeriods", evaluationPeriod.GetEvaluationPeriods)

	err := healthcheck.New(router, healthcheckConfig.DefaultConfig(), []healthcheckChecks.Check{})
	if err != nil {
		log.Fatal("Error creating healthcheck")
	}

	err = router.Run(":8000")
	if err != nil {
		log.Fatal("Error running server")
	}
	fmt.Println("Server running on port 8000")
}
