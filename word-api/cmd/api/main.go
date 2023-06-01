package main

import (
	"github.com/dilaragorum/word-explorer/word-api/internal/pkg/elastic"
	"github.com/dilaragorum/word-explorer/word-api/internal/vocabulary"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	conn, err := elastic.NewClient()
	if err != nil {
		panic("Cannot connect to DB")
	}

	// Vocabulary
	vocabRepo := vocabulary.NewRepository(conn)
	vocabService := vocabulary.NewService(vocabRepo)
	vocabulary.NewHandler(app, vocabService)

	if err = app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
