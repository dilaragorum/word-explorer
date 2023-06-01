package vocabulary

import "fmt"

type Vocabulary struct {
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}

func (v Vocabulary) ID() string {
	return fmt.Sprintf("%s-%s-%s", v.Word, v.Meaning, v.Sentence)
}

func rowToVocabulary(row []interface{}) Vocabulary {
	wordEvent := Vocabulary{
		Word: row[0].(string),
	}

	if len(row) >= 2 {
		wordEvent.Meaning = row[1].(string)
	}

	if len(row) >= 3 {
		wordEvent.Sentence = row[2].(string)
	}

	return wordEvent
}
