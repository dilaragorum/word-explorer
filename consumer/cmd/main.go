package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/dilaragorum/word-explorer/consumer/internal/vocabulary"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	consumer := vocabulary.NewConsumer()
	defer consumer.Close()

	consumerService := vocabulary.NewService(consumer)
	for {
		message, err := consumerService.ConsumeMessage(context.Background())
		if err != nil {
			log.Error(err)
			continue
		}

		bytesMsg, _ := json.Marshal(message)
		bufferBytes := bytes.NewBuffer(bytesMsg)

		_, err = http.Post("http://localhost:3000/api/v1/words", "application/json", bufferBytes)
		if err != nil {
			log.Error(err)
		}
	}
}
