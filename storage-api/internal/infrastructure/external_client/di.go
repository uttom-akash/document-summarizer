package externalclient

import "go.uber.org/fx"

var DI = fx.Options(
	fx.Provide(NewRabbitMQClient),
	fx.Provide(NewS3Client),
	fx.Provide(NewExternalClientConfig),
)