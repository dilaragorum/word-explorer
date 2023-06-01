package vocabulary

import (
	"context"
	"github.com/dilaragorum/word-explorer/consumer/pkg"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	Consume(ctx context.Context) (kafka.Message, error)
	Close() error
}

type consumer struct {
	reader *pkg.KafkaReader
}

func NewConsumer() Consumer {
	reader := pkg.NewReader("words", "words-consumer-group")
	return &consumer{reader: reader}
}

func (c *consumer) Consume(ctx context.Context) (kafka.Message, error) {
	message, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return kafka.Message{}, err
	}

	return message, nil
}

func (c *consumer) Close() error {
	return c.reader.Close()
}
