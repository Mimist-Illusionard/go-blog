package handlers

import (
	"go-blog/backend/internal/models"
	"go-blog/backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Service *services.PostService
}

func CreatePostHandler(s *services.PostService) *PostHandler {
	return &PostHandler{Service: s}
}

func (h *PostHandler) InitializeHandler(ginEngine *gin.Engine) {
	ginEngine.GET("/posts/*id", h.Get)
	ginEngine.POST("/posts/", h.Create)
	ginEngine.PUT("/posts/:id", h.Put)
	ginEngine.DELETE("/posts/:id", h.Delete)
}

func (h *PostHandler) Get(c *gin.Context) {
	postID := c.Param("id")

	if postID != "/" {
		id, err := strconv.Atoi(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		post, err := h.Service.GetPostByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}

		c.JSON(http.StatusOK, post)
		return
	}

	posts, err := h.Service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not get posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) Create(c *gin.Context) {
	var req struct {
		UserID uint   `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	post := &models.Post{
		Text: req.Text,
	}

	err := h.Service.CreatePost(req.UserID, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create post", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post created"})
}

func (h *PostHandler) Put(c *gin.Context) {
	postID := c.Param("id")

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var req struct {
		UserID uint   `json:"user_id"`
		Text   string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	editedPost := &models.Post{
		Text: req.Text,
	}

	if err := h.Service.EditPost(uint(id), editedPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	post, err := h.Service.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) Delete(c *gin.Context) {
	postID := c.Param("id")

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = h.Service.DeletePost(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
