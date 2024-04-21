//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"secret-management/internal/domain"
	"secret-management/internal/handler"
	"secret-management/internal/usecase"
)

type Handlers struct {
	HealthHandler *handler.HealthHandler
	SecretHandler *handler.SecretHandler
}

func InitializeDependency(router *gin.RouterGroup) (*Handlers, error) {
	wire.Build(
		// repository

		// usecase
		usecase.NewSecretUsecase,

		// handler
		handler.NewHealthHandler,
		handler.NewSecretHandler,

		// bind repository

		// bind usecase
		wire.Bind(new(domain.SecretUsecase), new(*usecase.SecretUsecase)),

		// bind handler

		// struct handler
		wire.Struct(new(Handlers), "*"),
	)

	return nil, nil
}
