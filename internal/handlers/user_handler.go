package handlers

import (
	"go-blog/internal/models"
	service "go-blog/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *service.UserService
}

func CreateUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) InitializeHandler(ginEngine *gin.Engine) {
	ginEngine.POST("/user/", h.Create)
	ginEngine.POST("/login/", h.Login)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	user := &models.User{
		Login:    req.Login,
		Password: req.Password,
	}

	err := h.Service.CreateUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Couldn't create user", "details": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User created"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	err := h.Service.Login(req.Login, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Incorrect password or login", "details": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "Login succesfull"})
}
