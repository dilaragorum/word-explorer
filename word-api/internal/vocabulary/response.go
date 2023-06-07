package vocabulary

type FilterResult struct {
	Total        int          `json:"total"`
	Vocabularies []Vocabulary `json:"vocabularies"`
}
