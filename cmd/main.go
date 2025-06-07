package main

import (
	"fmt"
	"go-blog/config"
	"go-blog/internal/handlers"
	"go-blog/internal/models"
	"go-blog/internal/repository"
	"go-blog/internal/services"
	"go-blog/utils/log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"

	//"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.LoadDatabaseConfig().GetDSN()), &gorm.Config{})
	if err != nil {
		log.Error("Failde to connect to database", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	if err != nil {
		fmt.Errorf("Failed to migrate: %v", err)
	}

	userRepository := &repository.UserGORMRepository{DB: db}
	userService := services.NewUserSevice(userRepository)
	userHandler := handlers.CreateUserHandler(userService)

	postRepository := &repository.PostGORMRepository{DB: db}
	postService := services.NewPostService(postRepository)
	postHandler := handlers.CreatePostHandler(postService)

	commentRepository := &repository.CommentGORMRepository{DB: db}
	commentService := services.NewCommentService(commentRepository)
	commentHandler := handlers.CreateCommentHandler(commentService)

	var ginEngine = gin.New()
	ginEngine.Use(CORSMiddleware())

	userHandler.InitializeHandler(ginEngine)
	postHandler.InitializeHandler(ginEngine)
	commentHandler.InitializeHandler(ginEngine)

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
