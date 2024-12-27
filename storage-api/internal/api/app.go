package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/uttom-akash/storage/internal/api/handlers"
	"github.com/uttom-akash/storage/internal/api/middlewares"
	"github.com/uttom-akash/storage/internal/application/services"
	"github.com/uttom-akash/storage/internal/infrastructure/db/postgres"
	externalclient "github.com/uttom-akash/storage/internal/infrastructure/external_client"
	"github.com/uttom-akash/storage/internal/infrastructure/repositories"
	"go.uber.org/fx"
)

func HookStartup(lc fx.Lifecycle, server *gin.Engine, document_handler *handlers.FileHandler) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Run(":8081")
			return nil
		},
	})
}

func httpServer() *gin.Engine {
	server := gin.Default()

	server.Use(middlewares.HandleError())

	return server
}

func NewApp() {
	fx.New(
		fx.Provide(
			handlers.DI,
			services.DI,
			repositories.DI,
			externalclient.DI,
		),

		fx.Invoke(postgres.RunMigrations),
		fx.Invoke(HookStartup),
	).Run()
}
