package handlers

import (
	"go-blog/backend/internal/models"
	"go-blog/backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func CreateUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) InitializeHandler(ginEngine *gin.Engine) {
	ginEngine.GET("/users/*id", h.Get)
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

func (h *UserHandler) Get(c *gin.Context) {
	postID := c.Param("id")

	if postID != "/" {
		id, err := strconv.Atoi(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		post, err := h.Service.GetUserByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		c.JSON(http.StatusOK, post)
		return
	}

	posts, err := h.Service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
