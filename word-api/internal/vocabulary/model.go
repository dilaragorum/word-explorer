package vocabulary

type Vocabulary struct {
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Sentence string `json:"sentence"`
}

type SearchModel struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
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

func (sm *SearchModel) ToDTO() *[]Vocabulary {
	var vocabs []Vocabulary

	results := sm.Hits.Hits
	for k := range results {
		v := Vocabulary{
			Word:     results[k].Source.Word,
			Meaning:  results[k].Source.Meaning,
			Sentence: results[k].Source.Sentence,
		}

		vocabs = append(vocabs, v)
	}

	return &vocabs
}

type SearchArgs struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	SubWord string `json:"sub_word"`
}
