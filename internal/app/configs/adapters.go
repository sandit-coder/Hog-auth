package configs

import (
	"Hog-auth/internal/app/adapters/secondary/kafka-adapter/kafka_client"
	"Hog-auth/internal/pkg/fiber"
)

type Adapters struct {
	Primary   Primary   `koanf:"primary"`
	Secondary Secondary `koanf:"secondary"`
}

type Primary struct {
	Fiber fiber.Config `koanf:"fiber"`
}

type Secondary struct {
	KafkaClient kafka_client.Config `koanf:"kafka_client"`
}
