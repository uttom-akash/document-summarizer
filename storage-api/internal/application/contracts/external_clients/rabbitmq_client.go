package externalclients

type IRabbitMQClient interface {
	PublishMessage(filename string)
}
