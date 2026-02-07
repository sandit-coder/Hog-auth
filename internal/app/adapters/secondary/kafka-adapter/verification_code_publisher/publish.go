package verification_code_publisher

import (
	"Hog-auth/internal/app/domain/events"
	"context"
	"encoding/json"

	"github.com/twmb/franz-go/pkg/kgo"
)

const (
	Topic = "user_verification.requested"
)

func (p *KafkaAdapterPublisher) Publish(ctx context.Context, event events.UserVerificationRequested) error {
	data, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	bytes := event.EventId[:]

	return p.client.ProduceSync(ctx, &kgo.Record{
		Topic:   Topic,
		Key:     bytes,
		Value:   data,
		Context: ctx,
	}).FirstErr()
}
