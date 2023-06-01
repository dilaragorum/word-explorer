package vocabulary

import (
	"context"
	"encoding/json"
)

type Service interface {
	ConsumeMessage(ctx context.Context) (Vocabulary, error)
}

type service struct {
	consumer Consumer
}

func NewService(consumer Consumer) Service {
	return &service{consumer: consumer}
}

func (s *service) ConsumeMessage(ctx context.Context) (Vocabulary, error) {
	message, err := s.consumer.Consume(ctx)
	if err != nil {
		return Vocabulary{}, err
	}

	if message.Value == nil {
		return Vocabulary{}, nil
	}

	var vocabulary Vocabulary
	if err = json.Unmarshal(message.Value, &vocabulary); err != nil {
		return Vocabulary{}, err
	}

	return vocabulary, nil
}
