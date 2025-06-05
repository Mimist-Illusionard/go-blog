package main

import (
	"go-blog/config"
	"go-blog/internal/handlers"
	"go-blog/internal/models"
	"go-blog/internal/repository"
	"go-blog/internal/services"
	"go-blog/utils/log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.LoadDatabaseConfig().GetDSN()), &gorm.Config{})
	if err != nil {
		log.Error("Failde to connect to database", err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})

	userRepository := &repository.UserGORMRepository{DB: db}
	userService := services.NewUserSevice(userRepository)
	userHandler := handlers.CreateUserHandler(userService)

	var ginEngine = gin.New()
	ginEngine.Use(CORSMiddleware())

	userHandler.InitializeHandler(ginEngine)

	ginEngine.Run(":9090")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
