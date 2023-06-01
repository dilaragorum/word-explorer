package vocabulary

import (
	"bufio"
	"context"
	"fmt"
	"github.com/dilaragorum/word-explorer/ingester/pkg"
	"os"
)

const readRange = "Kelimeler!A:C"

type Service interface {
	ProcessMessages(ctx context.Context) error
}

type SheetClient interface {
	GetSpreadsheetByID(spreadsheetId, readRange string) ([][]interface{}, error)
}

type service struct {
	publisher     Publisher
	spreadSheetID string
	sheetClient   SheetClient
}

func NewVocabularyService(spreadsheetId string, publisher Publisher, client SheetClient) Service {
	return &service{
		spreadSheetID: spreadsheetId,
		publisher:     publisher,
		sheetClient:   client,
	}
}

func (gsc *service) ProcessMessages(ctx context.Context) error {
	excelRows, err := gsc.getExcelCells(ctx)
	if err != nil {
		return err
	}

	hashFilePtr, err := os.OpenFile("hash.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error opening hash txt %v", err)
	}

	hashLines := readFileLineByLine(hashFilePtr)
	vocabularyHashMap := pkg.StringSliceToMap(hashLines)

	hashToBeSaved := make([]string, 0, len(excelRows))
	wordsToBePublished := make([]Vocabulary, 0, len(excelRows))

	for _, row := range excelRows {
		vocabulary := rowToVocabulary(row)
		hashedVocabulary := generateHashFromVocabulary(vocabulary)

		if _, ok := vocabularyHashMap[hashedVocabulary]; !ok {
			hashToBeSaved = append(hashToBeSaved, hashedVocabulary+"\n")
			wordsToBePublished = append(wordsToBePublished, vocabulary)
		}
	}

	if err = gsc.publishMessages(ctx, wordsToBePublished); err != nil {
		return err
	}

	for i := range hashToBeSaved {
		hashFilePtr.WriteString(hashToBeSaved[i])
	}

	return nil
}

func (gsc *service) publishMessages(ctx context.Context, wordEvents []Vocabulary) error {
	return gsc.publisher.PublishVocabularies(ctx, wordEvents)
}

func (gsc *service) getExcelCells(_ context.Context) ([][]interface{}, error) {
	rows, err := gsc.sheetClient.GetSpreadsheetByID(gsc.spreadSheetID, readRange)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		fmt.Println("No data found.")
		return nil, nil
	}

	return rows, nil
}

func generateHashFromVocabulary(vocabulary Vocabulary) string {
	return pkg.GetMD5Hash(vocabulary.ID())
}

func readFileLineByLine(file *os.File) []string {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
