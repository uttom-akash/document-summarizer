package services

import "go.uber.org/fx"

var DI = fx.Options(
	fx.Provide(NewFileService),
)