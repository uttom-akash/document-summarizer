package handlers

import "go.uber.org/fx"

var DI = fx.Options(
	fx.Provide(NewFileHandler),
)