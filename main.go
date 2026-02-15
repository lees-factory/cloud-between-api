package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	v1 "io.lees.cloud-between/core/core-api/controller/v1"
	"io.lees.cloud-between/core/core-domain/chemistry"
	"io.lees.cloud-between/core/core-domain/diagnosis"
	"io.lees.cloud-between/core/core-domain/persona"
	"io.lees.cloud-between/core/core-domain/user"
	"io.lees.cloud-between/storage/db-core/repository"
)

func main() {
	_ = godotenv.Load()

	// Database setup
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_SCHEMA")
	if dbSchema == "" {
		dbSchema = "public"
	}

	dsn := fmt.Sprintf("host=aws-1-ap-northeast-2.pooler.supabase.com user=postgres.cnlbftaurppzamivofhb password='%s' dbname=postgres port=6543 sslmode=require search_path=%s TimeZone=Asia/Seoul", dbPassword, dbSchema)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal("failed to connect database: ", err) // 에러 상세 내용 출력
	}

	// Data Access Layer
	userRepository := repository.NewUserCoreRepository(db)
	diagnosisRepository := repository.NewDiagnosisCoreRepository(db)
	chemistryRepository := repository.NewChemistryCoreRepository(db)
	personaProfileRepository := repository.NewPersonaProfileCoreRepository(db)

	// Domains
	userAppender := user.NewUserAppender(userRepository)
	userFinder := user.NewUserFinder(userRepository)
	userUpdater := user.NewUserUpdater(userRepository)
	userService := user.NewUserService(userAppender, userFinder, userUpdater)

	diagnosisService := diagnosis.NewDiagnosisService(diagnosisRepository)

	chemistryFinder := chemistry.NewChemistryFinder(chemistryRepository)
	chemistryService := chemistry.NewChemistryService(chemistryFinder)

	personaProfileFinder := persona.NewPersonaProfileFinder(personaProfileRepository)
	personaProfileService := persona.NewPersonaProfileService(personaProfileFinder)

	// Presentation Layer
	userController := v1.NewUserController(userService)
	diagnosisController := v1.NewDiagnosisController(diagnosisService)
	chemistryController := v1.NewChemistryController(chemistryService)
	personaProfileController := v1.NewPersonaProfileController(personaProfileService)

	// Router setup
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger UI
	r.StaticFile("/docs/openapi.yaml", "./core/core-api/docs/openapi.yaml")
	r.GET("/docs", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!DOCTYPE html>
<html><head>
<title>Cloud Between API</title>
<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css">
</head><body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
<script>SwaggerUIBundle({url:"/docs/openapi.yaml",dom_id:"#swagger-ui"})</script>
</body></html>`))
	})

	apiV1 := r.Group("/api/v1")
	{
		userGroup := apiV1.Group("/users")
		{
			userGroup.POST("/signup", userController.Signup)
			userGroup.POST("/login", userController.Login)
			userGroup.POST("/login/social", userController.SocialLogin)
		}

		diagnosisGroup := apiV1.Group("/diagnosis")
		{
			diagnosisGroup.GET("/questions", diagnosisController.GetQuestions)
			diagnosisGroup.POST("/analyze", diagnosisController.Analyze)
		}

		chemistryGroup := apiV1.Group("/chemistries")
		{
			chemistryGroup.GET("", chemistryController.GetAllChemistries)
			chemistryGroup.GET("/match", chemistryController.GetChemistry)
		}

		personaGroup := apiV1.Group("/personas")
		{
			personaGroup.GET("/profiles", personaProfileController.GetProfiles)
			personaGroup.GET("/profiles/:typeKey", personaProfileController.GetProfile)
		}
	}

	r.Run(":8081")
}
