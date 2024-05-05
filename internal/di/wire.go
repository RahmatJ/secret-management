//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
	"secret-management/config"
	"secret-management/config/cronConfig"
	"secret-management/internal/domain"
	"secret-management/internal/handler"
	"secret-management/internal/helpers"
	"secret-management/internal/repository"
	"secret-management/internal/usecase"
)

type Handlers struct {
	HealthHandler *handler.HealthHandler
	SecretHandler *handler.SecretHandler
	CronSetup     *cronConfig.CronSetup
}

func InitializeDependency(router *gin.RouterGroup, db *gorm.DB, config config.EnvConfig) (*Handlers, error) {
	wire.Build(
		// repository
		repository.NewSecretRepository,

		// usecase
		usecase.NewSecretUsecase,

		// helpers
		helpers.NewSecretHelpers,

		// handler
		handler.NewHealthHandler,
		handler.NewSecretHandler,

		// cron
		cronConfig.NewCronSetup,

		// bind repository
		wire.Bind(new(domain.SecretRepository), new(*repository.SecretRepository)),

		// bind usecase
		wire.Bind(new(domain.SecretUsecase), new(*usecase.SecretUsecase)),

		// bind helpers
		wire.Bind(new(domain.SecretHelpers), new(*helpers.SecretHelpers)),

		// bind handler

		// struct handler
		wire.Struct(new(Handlers), "*"),
	)

	return &Handlers{}, nil
}
