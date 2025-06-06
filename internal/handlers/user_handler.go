package handlers

import (
	"go-blog/internal/models"
	"go-blog/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func CreateUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) InitializeHandler(ginEngine *gin.Engine) {
	ginEngine.POST("/users/", h.Create)
	ginEngine.POST("/auth/", h.Authorization)
}

func (h *UserHandler) Create(c *gin.Context) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	user := &models.User{
		Login:    req.Login,
		Password: req.Password,
	}

	err := h.Service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (h *UserHandler) Authorization(c *gin.Context) {
	var req struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	err := h.Service.Login(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Incorrect password or login", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login succesfull"})
}
