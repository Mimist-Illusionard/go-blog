package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	InitializeHandler(ginEngine *gin.Engine)
}
