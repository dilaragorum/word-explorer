package vocabulary

import (
	"context"
	"encoding/json"
	"github.com/dilaragorum/word-explorer/ingester/pkg"
	"github.com/segmentio/kafka-go"
)

type Publisher interface {
	PublishVocabularies(context.Context, []Vocabulary) error
	Close() error
}

type publisher struct {
	publish *pkg.KafkaWriter
}

func NewPublisher() Publisher {
	writer := pkg.NewWriter("words")
	return &publisher{publish: writer}
}

func (p *publisher) PublishVocabularies(ctx context.Context, vocabularies []Vocabulary) error {
	messages := make([]kafka.Message, 0, len(vocabularies))

	for i := range vocabularies {
		vocabularyBytes, _ := json.Marshal(vocabularies[i])
		messages = append(messages,
			kafka.Message{
				Value: vocabularyBytes,
			})
	}

	return p.publish.WriteMessages(ctx, messages...)
}

func (p *publisher) Close() error {
	return p.publish.Close()
}
