package pkg

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	kr *kafka.Reader
}

// TODO: Consumer Group kavramını ögrenelim.
func NewReader(topic, groupID string) *KafkaReader {
	return &KafkaReader{
		kr: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{"localhost:29092"},
			GroupID:     groupID,
			GroupTopics: []string{topic},
		}),
	}
}

func (r *KafkaReader) ReadMessage(ctx context.Context) (kafka.Message, error) {
	message, err := r.kr.ReadMessage(ctx)
	if err != nil {
		return kafka.Message{}, nil
	}
	return message, nil
}

func (r *KafkaReader) Close() error {
	return r.kr.Close()
}
