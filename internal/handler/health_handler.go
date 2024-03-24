package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler(group *gin.RouterGroup) *HealthHandler {
	handler := HealthHandler{}

	handlerGroup := group.Group("/ping")

	handlerGroup.GET("", handler.ping)

	return &handler
}

func (h *HealthHandler) ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "pong",
	})
}
