package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"secret-management/internal/domain"
	"secret-management/internal/dto"
)

type SecretHandler struct {
	secretUsecase domain.SecretUsecase
}

func (h SecretHandler) getSecretByUserId(context *gin.Context) {
	userId := context.Param("userId")
	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"data": "userId required",
		})
	}

	result, err := h.secretUsecase.GetSecretByUserId(userId)
	if err != nil {
		log.Print(err.Error())
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (h SecretHandler) validateKey(context *gin.Context) {
	var payload dto.ValidateUserKey
	err := context.ShouldBindJSON(&payload)
	if err != nil {
		// TODO(Rahmat): add validation error handling
		log.Print(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": "valid",
	})
}

func (h SecretHandler) check(context *gin.Context) {
	err := h.secretUsecase.DailySecretCheck()
	if err != nil {
		log.Print(err.Error())
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"data": "OK",
	})
}

func NewSecretHandler(group *gin.RouterGroup, secretUsecase domain.SecretUsecase) *SecretHandler {
	handler := SecretHandler{
		secretUsecase: secretUsecase,
	}

	handlerGroup := group.Group("/secret")

	handlerGroup.GET("/:userId", handler.getSecretByUserId)
	handlerGroup.POST("/validate", handler.validateKey)
	//TODO(Rahmat): Delete this, this only for testing
	handlerGroup.POST("/check", handler.check)

	return &handler
}
