package repositories

import (
	"github.com/uttom-akash/storage/internal/infrastructure/db/postgres"
	"go.uber.org/fx"
)

var DI = fx.Options(
	fx.Provide(postgres.NewDB),
	fx.Provide(NewFileRepository),
)