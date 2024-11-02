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
	indicatorsInformationBooksOrChaptersProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/booksOrChaptersProduction"
	indicatorInformationGradeRate "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/gradeRate"
	researchInnovationProject "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/researchInnovationProjectLists"
	indicatorInformationSocialProject "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/socialProjectProduction"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
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
		&knowledgeField.WideField{},
		&knowledgeField.SpecificField{},
		&knowledgeField.DetailedField{},
		&indicatorInformationGradeRate.GradeRateList{},
		&indicatorsInformation.AcademicPeriodAuthorCareer{},
		&indicatorInformationSocialProject.SocialProjectList{},
		&researchInnovationProject.ResearchInnovationProjectList{},
		&indicatorsInformationAcademicProduction.AcademicProductionList{},
		&indicatorsInformationAcademicProduction.AcademicProductionListsAuthor{},
		&indicatorsInformationBooksOrChaptersProduction.BooksOrChaptersProductionList{},
		&indicatorsInformationBooksOrChaptersProduction.BooksOrChaptersProductionListAuthor{},
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
		{ID: 0, Name: "N/A", Abbreviation: "N/A", Description: "N/A"},
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
		{
			ID:           19,
			Name:         "Indicador 19: Tasa de Deserción Institucional de Segundo Año – Oferta Académica de Grado",
			Abbreviation: "Tasa de Deserción Institucional de Segundo Año – Oferta Académica de Grado",
			Description:  "La institución cuenta con una tasa promedio de deserción de estudiantes de grado al segundo año de máximo el 14%.",
		},
		{
			ID:           21,
			Name:         "Indicador 21: Tasa de Titulación Institucional - Oferta Académica de Grado",
			Abbreviation: "Tasa de Titulación Institucional - Oferta Académica de Grado",
			Description:  "La institución cuenta con una tasa promedio de titulación institucional de la oferta académica de grado de al menos el 50%.",
		},
		{
			ID:           25,
			Name:         "Indicador 25: Proyectos de Investigación e Innovación con Financiamiento Externo o en Red",
			Abbreviation: "Proyectos de Investigación e Innovación con Financiamiento Externo o en Red",
			Description:  "Se espera que al menos el 40% de los proyectos de investigación e innovación concluidos o en ejecución cuenten con financiamiento externo o con participación en redes internacionales o nacionales.",
		},
		{
			ID:           26,
			Name:         "Indicador 26: Producción Académica",
			Abbreviation: "Producción Académica",
			Description:  "La institución cuenta con producción académica como resultado de sus procesos de investigación. Se espera que el índice de producción académica per cápita sea de al menos 1,5 en 3 años.",
		},
		{
			ID:           29,
			Name:         "Indicador 29: Proyectos de vinculación con la sociedad",
			Abbreviation: "Proyectos de vinculación con la sociedad",
			Description:  "Se espera un mínimo de 1,5 proyectos de vinculación por carrera y programa con resultados verificables totales o parciales.",
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
		{ID: 9, AcademicDatabaseID: 5, CompensationFactorID: 7},
	}

	var wideField = []knowledgeField.WideField{
		{ID: 1, Name: "Educación", Abbreviation: "Educación"},
		{ID: 2, Name: "Artes y Humanidades", Abbreviation: "Artes y Humanidades"},
		{ID: 3, Name: "Ciencias sociales, Periodismo, información y derecho", Abbreviation: "Ciencias sociales, Periodismo, información y derecho"},
		{ID: 4, Name: "Administración", Abbreviation: "Administración"},
		{ID: 5, Name: "Ciencias naturales, matemáticas y estadística", Abbreviation: "Ciencias naturales, matemáticas y estadística"},
		{ID: 6, Name: "Tecnologías de la información y la comunicación (TIC)", Abbreviation: "Tecnologías de la información y la comunicación (TIC)"},
		{ID: 7, Name: "Ingeniería, industria y construcción", Abbreviation: "Ingeniería, industria y construcción"},
		{ID: 8, Name: "Agricultura, silvicultura, pesca y veterinaria", Abbreviation: "Agricultura, silvicultura, pesca y veterinaria"},
		{ID: 9, Name: "Salud y bienestar", Abbreviation: "Salud y bienestar"},
		{ID: 10, Name: "Servicios", Abbreviation: "Servicios"},
	}

	var specificField = []knowledgeField.SpecificField{
		{ID: 1, Name: "Educación", Description: "Educación", WideFieldID: 1},
		{ID: 2, Name: "Artes", Description: "Artes", WideFieldID: 2},
		{ID: 3, Name: "Humanidades", Description: "Humanidades", WideFieldID: 2},
		{ID: 4, Name: "Idiomas", Description: "Idiomas", WideFieldID: 2},
		{ID: 5, Name: "Ciencias sociales y comportamiento", Description: "Ciencias sociales y comportamiento", WideFieldID: 3},
		{ID: 6, Name: "Periodismo e información", Description: "Periodismo e información", WideFieldID: 3},
		{ID: 7, Name: "Derecho", Description: "Derecho", WideFieldID: 3},
		{ID: 8, Name: "Educación comercial y administración", Description: "Educación comercial y administración", WideFieldID: 4},
		{ID: 9, Name: "Ciencias biológicas y afines", Description: "Ciencias biológicas y afines", WideFieldID: 5},
		{ID: 10, Name: "Medio ambiente", Description: "Medio ambiente", WideFieldID: 5},
		{ID: 11, Name: "Ciencias físicas", Description: "Ciencias físicas", WideFieldID: 5},
		{ID: 12, Name: "Matemáticas y estadística", Description: "Matemáticas y estadística", WideFieldID: 5},
		{ID: 13, Name: "Tecnologías de la información y la comunicación (TIC)", Description: "Tecnologías de la información y la comunicación (TIC)", WideFieldID: 6},
		{ID: 14, Name: "Ingeniería y profesiones afines", Description: "Ingeniería y profesiones afines", WideFieldID: 7},
		{ID: 15, Name: "Industria y producción", Description: "Industria y producción", WideFieldID: 7},
		{ID: 16, Name: "Arquitectura y construcción", Description: "Arquitectura y construcción", WideFieldID: 7},
		{ID: 17, Name: "Agricultura", Description: "Agricultura", WideFieldID: 8},
		{ID: 18, Name: "Silvicultura", Description: "Silvicultura", WideFieldID: 8},
		{ID: 19, Name: "Pesca", Description: "Pesca", WideFieldID: 8},
		{ID: 20, Name: "Veterinaria", Description: "Veterinaria", WideFieldID: 8},
		{ID: 21, Name: "Salud", Description: "Salud", WideFieldID: 9},
		{ID: 22, Name: "Bienestar", Description: "Bienestar", WideFieldID: 9},
		{ID: 23, Name: "Servicio personales", Description: "Servicio personales", WideFieldID: 10},
		{ID: 24, Name: "Servicios de protección", Description: "Servicios de protección", WideFieldID: 10},
		{ID: 25, Name: "Servicios de seguridad", Description: "Servicios de seguridad", WideFieldID: 10},
		{ID: 26, Name: "Servicios de transporte", Description: "Servicios de transporte", WideFieldID: 10},
	}

	var detailedField = []knowledgeField.DetailedField{
		{Name: "Educación", Abbreviation: "Educación", SpecificFieldID: 1},
		{Name: "Psicopedagogía", Abbreviation: "Psicopedagogía", SpecificFieldID: 1},
		{Name: "Formación para docentes de educación preprimaria", Abbreviation: "Formación para docentes de educación preprimaria", SpecificFieldID: 1},
		{Name: "Formación para docentes sin asignaturas de especialización", Abbreviation: "Formación para docentes sin asignaturas de especialización", SpecificFieldID: 1},
		{Name: "Formación para docentes con asignaturas de especialización", Abbreviation: "Formación para docentes con asignaturas de especialización", SpecificFieldID: 1},
		{Name: "Técnicas audiovisuales y producción para medios de comunicación", Abbreviation: "Técnicas audiovisuales y producción para medios de comunicación", SpecificFieldID: 2},
		{Name: "Diseño", Abbreviation: "Diseño", SpecificFieldID: 2},
		{Name: "Artes", Abbreviation: "Artes", SpecificFieldID: 2},
		{Name: "Música y artes escénicas", Abbreviation: "Música y artes escénicas", SpecificFieldID: 2},
		{Name: "Religión y Teología", Abbreviation: "Religión y Teología", SpecificFieldID: 3},
		{Name: "Historia y arqueolgia", Abbreviation: "Historia y arqueolgia", SpecificFieldID: 3},
		{Name: "Filosofía", Abbreviation: "Filosofía", SpecificFieldID: 3},
		{Name: "Idiomas", Abbreviation: "Idiomas", SpecificFieldID: 4},
		{Name: "Literatura y lingüistica", Abbreviation: "Literatura y lingüistica", SpecificFieldID: 4},
		{Name: "Economía", Abbreviation: "Economía", SpecificFieldID: 5},
		{Name: "Economía Matemática", Abbreviation: "Economía Matemática", SpecificFieldID: 5},
		{Name: "Ciencias políticas", Abbreviation: "Ciencias políticas", SpecificFieldID: 5},
		{Name: "Psicología", Abbreviation: "Psicología", SpecificFieldID: 5},
		{Name: "Estudios Sociales y Culturales", Abbreviation: "Estudios Sociales y Culturales", SpecificFieldID: 5},
		{Name: "Estudios de Género", Abbreviation: "Estudios de Género", SpecificFieldID: 5},
		{Name: "Geografía y Territorio", Abbreviation: "Geografía y Territorio", SpecificFieldID: 5},
		{Name: "Periodismo y comunicación", Abbreviation: "Periodismo y comunicación", SpecificFieldID: 6},
		{Name: "Bibliotecología, documentación y archivología", Abbreviation: "Bibliotecología, documentación y archivología", SpecificFieldID: 6},
		{Name: "Derecho", Abbreviation: "Derecho", SpecificFieldID: 7},
		{Name: "Contabilidad y auditoría", Abbreviation: "Contabilidad y auditoría", SpecificFieldID: 8},
		{Name: "Gestión Financiera", Abbreviation: "Gestión Financiera", SpecificFieldID: 8},
		{Name: "Administración", Abbreviation: "Administración", SpecificFieldID: 8},
		{Name: "Mercadotecnia y publicidad", Abbreviation: "Mercadotecnia y publicidad", SpecificFieldID: 8},
		{Name: "Información gerencial", Abbreviation: "Información gerencial", SpecificFieldID: 8},
		{Name: "Comercio", Abbreviation: "Comercio", SpecificFieldID: 8},
		{Name: "Competencias laborales", Abbreviation: "Competencias laborales", SpecificFieldID: 8},
		{Name: "Biología", Abbreviation: "Biología", SpecificFieldID: 9},
		{Name: "Biofísica", Abbreviation: "Biofísica", SpecificFieldID: 9},
		{Name: "Biofarmacéutica", Abbreviation: "Biofarmacéutica", SpecificFieldID: 9},
		{Name: "Biomedicina", Abbreviation: "Biomedicina", SpecificFieldID: 9},
		{Name: "Bioquímica", Abbreviation: "Bioquímica", SpecificFieldID: 9},
		{Name: "Genética", Abbreviation: "Genética", SpecificFieldID: 9},
		{Name: "Biodiversidad", Abbreviation: "Biodiversidad", SpecificFieldID: 9},
		{Name: "Neurociencias", Abbreviation: "Neurociencias", SpecificFieldID: 9},
		{Name: "Medio ambiente", Abbreviation: "Medio ambiente", SpecificFieldID: 10},
		{Name: "Recursos Naturales Renovables", Abbreviation: "Recursos Naturales Renovables", SpecificFieldID: 10},
		{Name: "Química", Abbreviation: "Química", SpecificFieldID: 11},
		{Name: "Ciencias de la Tierra", Abbreviation: "Ciencias de la Tierra", SpecificFieldID: 11},
		{Name: "Física", Abbreviation: "Física", SpecificFieldID: 11},
		{Name: "Matemáticas", Abbreviation: "Matemáticas", SpecificFieldID: 12},
		{Name: "Estadísticas", Abbreviation: "Estadísticas", SpecificFieldID: 12},
		{Name: "Logística y trasnporte", Abbreviation: "Logística y trasnporte", SpecificFieldID: 12},
		{Name: "Computación", Abbreviation: "Computación", SpecificFieldID: 13},
		{Name: "Diseño y administración de redes y bases de datos", Abbreviation: "Diseño y administración de redes y bases de datos", SpecificFieldID: 13},
		{Name: "Desarollo y análisis de software y aplicaciones", Abbreviation: "Desarollo y análisis de software y aplicaciones", SpecificFieldID: 13},
		{Name: "Sistemas de Información", Abbreviation: "Sistemas de Información", SpecificFieldID: 13},
		{Name: "Química Aplicada", Abbreviation: "Química Aplicada", SpecificFieldID: 14},
		{Name: "Tecnología de protección del medio ambiente", Abbreviation: "Tecnología de protección del medio ambiente", SpecificFieldID: 14},
		{Name: "Electricidad y energía", Abbreviation: "Electricidad y energía", SpecificFieldID: 14},
		{Name: "Electrónica, automatización y sonido", Abbreviation: "Electrónica, automatización y sonido", SpecificFieldID: 14},
		{Name: "Mecánica y profesiones afines a la metalistería", Abbreviation: "Mecánica y profesiones afines a la metalistería", SpecificFieldID: 14},
		{Name: "Diseño y construcción de vehículos, barcos y aeronaves motorizadas", Abbreviation: "Diseño y construcción de vehículos, barcos y aeronaves motorizadas", SpecificFieldID: 14},
		{Name: "Tecnologías Nucleares y Energéticas", Abbreviation: "Tecnologías Nucleares y Energéticas", SpecificFieldID: 14},
		{Name: "Mecatrónica", Abbreviation: "Mecatrónica", SpecificFieldID: 14},
		{Name: "Hidráulica", Abbreviation: "Hidráulica", SpecificFieldID: 14},
		{Name: "Telecomunicaciones", Abbreviation: "Telecomunicaciones", SpecificFieldID: 14},
		{Name: "Nanotecnología", Abbreviation: "Nanotecnología", SpecificFieldID: 14},
		{Name: "Procesamiento de alimentos", Abbreviation: "Procesamiento de alimentos", SpecificFieldID: 15},
		{Name: "Materiales", Abbreviation: "Materiales", SpecificFieldID: 15},
		{Name: "Productos textiles", Abbreviation: "Productos textiles", SpecificFieldID: 15},
		{Name: "Minería y extracción", Abbreviation: "Minería y extracción", SpecificFieldID: 15},
		{Name: "Producción Industrial", Abbreviation: "Producción Industrial", SpecificFieldID: 15},
		{Name: "Seguridad Industrial", Abbreviation: "Seguridad Industrial", SpecificFieldID: 15},
		{Name: "Diseño Industrial y de Procesos", Abbreviation: "Diseño Industrial y de Procesos", SpecificFieldID: 15},
		{Name: "Mantenimiento Industrial", Abbreviation: "Mantenimiento Industrial", SpecificFieldID: 15},
		{Name: "Arquitectura, urbanismo y restauración", Abbreviation: "Arquitectura, urbanismo y restauración	", SpecificFieldID: 16},
		{Name: "Construcción e ingenierfa civil", Abbreviation: "Construcción e ingenierfa civil", SpecificFieldID: 16},
		{Name: "Producción agricola y ganadera", Abbreviation: "Producción agricola y ganadera", SpecificFieldID: 17},
		{Name: "Silvicultura", Abbreviation: "Silvicultura", SpecificFieldID: 18},
		{Name: "Pesca", Abbreviation: "Pesca", SpecificFieldID: 19},
		{Name: "Veterinaria", Abbreviation: "Veterinaria", SpecificFieldID: 20},
		{Name: "Odontología", Abbreviation: "Odontología", SpecificFieldID: 21},
		{Name: "Medicina", Abbreviation: "Medicina", SpecificFieldID: 21},
		{Name: "Enfermería y obstetricia", Abbreviation: "Enfermería y obstetricia", SpecificFieldID: 21},
		{Name: "Tecnología de diagnóstico y tratamiento médico", Abbreviation: "Tecnología de diagnóstico y tratamiento médico", SpecificFieldID: 21},
		{Name: "Terapia y Rehabilitación", Abbreviation: "Terapia y Rehabilitación", SpecificFieldID: 21},
		{Name: "Farmacia", Abbreviation: "Farmacia", SpecificFieldID: 21},
		{Name: "Terapias alternativas y complementarias", Abbreviation: "Terapias alternativas y complementarias", SpecificFieldID: 21},
		{Name: "Bioquímica y Farmacia", Abbreviation: "Bioquímica y Farmacia", SpecificFieldID: 21},
		{Name: "Asistencia a adultos mayores y discapacitados", Abbreviation: "Asistencia a adultos mayores y discapacitados", SpecificFieldID: 22},
		{Name: "Asistencia a la infancia y servicios para jóvenes", Abbreviation: "Asistencia a la infancia y servicios para jóvenes", SpecificFieldID: 22},
		{Name: "Peluquería y tratamientos de belleza", Abbreviation: "Peluquería y tratamientos de belleza", SpecificFieldID: 23},
		{Name: "Hoteleria y Gastronomía", Abbreviation: "Hoteleria y Gastronomía", SpecificFieldID: 23},
		{Name: "Actividad Física", Abbreviation: "Actividad Física", SpecificFieldID: 23},
		{Name: "Turismo", Abbreviation: "Turismo", SpecificFieldID: 23},
		{Name: "Prevención y Gestión de Riesgos", Abbreviation: "Prevención y Gestión de Riesgos", SpecificFieldID: 24},
		{Name: "Salud y seguridad ocupacional", Abbreviation: "Salud y seguridad ocupacional", SpecificFieldID: 24},
		{Name: "Educación policial, militar y defensa", Abbreviation: "Educación policial, militar y defensa", SpecificFieldID: 25},
		{Name: "Seguridad Ciudadana", Abbreviation: "Seguridad Ciudadana", SpecificFieldID: 25},
		{Name: "Gestión del Transporte", Abbreviation: "Gestión del Transporte", SpecificFieldID: 26},
	}

	database.DB.Save(&wideField)
	database.DB.Save(&specificField)
	database.DB.Save(&detailedField)
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
