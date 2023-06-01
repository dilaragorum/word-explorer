package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	kw *kafka.Writer
}

func NewWriter(topic string) *KafkaWriter {
	return &KafkaWriter{
		kw: &kafka.Writer{
			Topic:                  topic,
			Addr:                   kafka.TCP("localhost:29092"),
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
	}
}

func (w *KafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return w.kw.WriteMessages(ctx, msgs...)
}

func (w *KafkaWriter) Close() error {
	return w.kw.Close()
}
