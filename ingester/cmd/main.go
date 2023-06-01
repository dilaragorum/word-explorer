package main

import (
	"context"
	"github.com/dilaragorum/word-explorer/ingester/internal/vocabulary"
	"github.com/dilaragorum/word-explorer/ingester/pkg"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("../ingester/.env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	publisher := vocabulary.NewPublisher()
	defer publisher.Close()

	spreadsheetId := viper.GetString("SPREADSHEET_ID")

	googleSheetClient, err := vocabulary.NewGoogleSheetClient()
	if err != nil {
		log.Fatalf("error initializing google sheet client %v", err)
	}

	vocabularyService := vocabulary.NewVocabularyService(spreadsheetId, publisher, googleSheetClient)
	cronClient := pkg.NewCronClient()
	cronClient.Schedule("30m", func() {
		if err = vocabularyService.ProcessMessages(context.Background()); err != nil {
			log.Error("Error cron function " + err.Error())
		}
	})
	cronClient.Start()
	defer cronClient.Stop()

	select {}
}
