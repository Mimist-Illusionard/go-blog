package handlers

import (
	"go-blog/internal/models"
	"go-blog/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Service *services.CommentService
}

func CreateCommentHandler(s *services.CommentService) *CommentHandler {
	return &CommentHandler{Service: s}
}

func (h *CommentHandler) InitializeHandler(ginEngine *gin.Engine) {
	ginEngine.GET("/comments/:postId", h.Get)
	ginEngine.POST("/comments/", h.Create)
	ginEngine.PUT("/comments/:id", h.Put)
	ginEngine.DELETE("/comments/", h.Delete)
}

func (h *CommentHandler) Get(c *gin.Context) {
	postID := c.Param("postId")

	id, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := h.Service.GetAllComments(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req struct {
		UserId uint   `json:"user_id"`
		PostId uint   `json:"post_id"`
		Text   string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	comment := &models.Comment{
		UserID: req.UserId,
		PostID: req.PostId,
		Text:   req.Text,
	}

	err := h.Service.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create comment on post", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment created"})
}

func (h *CommentHandler) Put(c *gin.Context) {
	commentId := c.Param("id")

	id, err := strconv.Atoi(commentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var req struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input", "details": err.Error()})
		return
	}

	editedPost := &models.Comment{
		Text: req.Text,
	}

	if err := h.Service.EditComment(uint(id), editedPost); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	post, err := h.Service.GetCommentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *CommentHandler) Delete(c *gin.Context) {
	var req struct {
		UserId    uint `json:"user_id"`
		CommentId uint `json:"comment_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	comment, err := h.Service.GetCommentByID(req.CommentId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if comment.UserID != req.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	if err := h.Service.DeleteComment(req.CommentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
