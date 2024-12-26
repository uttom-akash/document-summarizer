package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/uttom-akash/storage/internal/api/middlewares"
	"github.com/uttom-akash/storage/internal/application/services"
	"go.uber.org/fx"
)

func HookStartup(lc fx.Lifecycle, server *gin.Engine, document_handler *services.DocumentHandler) {

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
			httpServer,
			services.NewDocumentHandler,
		),
		fx.Invoke(HookStartup),
	).Run()
}
