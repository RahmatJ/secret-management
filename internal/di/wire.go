//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"secret-management/internal/handler"
)

type Handlers struct {
	HealthHandler *handler.HealthHandler
}

func InitializeDependency(router *gin.RouterGroup) (*Handlers, error) {
	wire.Build(
		handler.NewHealthHandler,

		wire.Struct(new(Handlers), "*"),
	)

	return nil, nil
}
