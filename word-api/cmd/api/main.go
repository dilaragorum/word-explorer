package main

import (
	"github.com/dilaragorum/word-explorer/word-api/internal/vocabulary"
	"github.com/dilaragorum/word-explorer/word-api/pkg/elastic"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	esClient, err := elastic.NewClient()
	if err != nil {
		log.Fatal("Cannot connect to DB ", err.Error())
	}

	// Vocabulary
	vocabRepo := vocabulary.NewRepository(esClient)
	vocabService := vocabulary.NewService(vocabRepo)
	vocabulary.NewHandler(app, vocabService)

	if err = app.Listen(":3200"); err != nil {
		log.Fatal(err)
	}
}
