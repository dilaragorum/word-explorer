package vocabulary

type Vocabulary struct {
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}

type SearchModel struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Index  string  `json:"_index"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Word     string `json:"word"`
				Meaning  string `json:"meaning"`
				Sentence string `json:"sentence"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (sm *SearchModel) ToDTO() FilterResult {
	var responseModel FilterResult

	responseModel.Total = sm.Hits.Total.Value

	vocabs := sm.toSearchModelToVocab()

	responseModel.Vocabularies = vocabs

	return responseModel
}

func (sm *SearchModel) toSearchModelToVocab() []Vocabulary {
	results := sm.Hits.Hits

	vocabs := make([]Vocabulary, 0, len(results))

	for k := range results {
		v := Vocabulary{
			Word:     results[k].Source.Word,
			Meaning:  results[k].Source.Meaning,
			Sentence: results[k].Source.Sentence,
		}

		vocabs = append(vocabs, v)
	}

	return vocabs
}
