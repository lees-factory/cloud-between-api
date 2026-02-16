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
	"io.lees.cloud-between/core/core-domain/premiumcard"
	"io.lees.cloud-between/core/core-domain/translation"
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
	translationRepository := repository.NewTranslationCoreRepository(db)
	premiumCardRepository := repository.NewPremiumCardCoreRepository(db)

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

	translationFinder := translation.NewTranslationFinder(translationRepository)
	translationService := translation.NewTranslationService(translationFinder)

	premiumCardFinder := premiumcard.NewPremiumCardFinder(premiumCardRepository)
	premiumCardService := premiumcard.NewPremiumCardService(premiumCardFinder)

	// Presentation Layer
	userController := v1.NewUserController(userService)
	diagnosisController := v1.NewDiagnosisController(diagnosisService)
	chemistryController := v1.NewChemistryController(chemistryService)
	personaProfileController := v1.NewPersonaProfileController(personaProfileService)
	translationController := v1.NewTranslationController(translationService)
	premiumCardController := v1.NewPremiumCardController(premiumCardService)

	// Router setup
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		allowed := map[string]bool{
			"http://localhost:5173":                  true,
			"https://www.cloudbetweenus.duckdns.org": true,
		}
		if allowed[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

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

		translationGroup := apiV1.Group("/translations")
		{
			translationGroup.GET("", translationController.GetAll)
			translationGroup.GET("/:namespace", translationController.GetByNamespace)
		}

		premiumCardGroup := apiV1.Group("/premium-cards")
		{
			premiumCardGroup.GET("", premiumCardController.GetAll)
			premiumCardGroup.GET("/:category", premiumCardController.GetByCategory)
		}
	}

	r.Run(":8081")
}
