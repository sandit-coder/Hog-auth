package verification_code_publisher

import (
	"Hog-auth/internal/app/adapters/secondary/kafka-adapter/kafka_client"
	"Hog-auth/internal/app/application/ports/publishers"
	"log/slog"

	"github.com/twmb/franz-go/pkg/kgo"
)

type KafkaAdapterPublisher struct {
	logger *slog.Logger
	config kafka_client.Config
	client *kgo.Client
}

func New(logger *slog.Logger, client *kgo.Client, config kafka_client.Config) publishers.VerificationCode {
	return &KafkaAdapterPublisher{
		logger: logger,
		config: config,
		client: client,
	}
}
