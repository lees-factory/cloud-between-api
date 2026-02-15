package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"io.lees.cloud-between/core/core-api/user/handler"
	"io.lees.cloud-between/core/core-domain/user/business"
	"io.lees.cloud-between/core/core-domain/user/implement"
	userrepo "io.lees.cloud-between/storage/db-core/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database setup (mocking DSN for now)
	dsn := "host=localhost user=postgres password=postgres dbname=cloud_between port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Data Access Layer
	userRepository := userrepo.NewUserRepository(db)

	// Implement Layer
	userAppender := implement.NewUserAppender(userRepository)
	userFinder := implement.NewUserFinder(userRepository)
	userUpdater := implement.NewUserUpdater(userRepository)

	// Business Layer
	userService := business.NewUserService(userAppender, userFinder, userUpdater)

	// Presentation Layer
	userHandler := handler.NewUserHandler(userService)

	// Router setup
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/signup", userHandler.Signup)
		userGroup.POST("/login", userHandler.Login)
	}

	r.Run(":8080")
}
