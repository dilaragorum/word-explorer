package vocabulary

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/gommon/log"
	"io"
	"strings"
	"time"
)

type repository struct {
	esc *elasticsearch.Client
}

func NewRepository(esc *elasticsearch.Client) Repository {
	return &repository{esc: esc}
}

func (r *repository) Create(_ context.Context, vocabulary *Vocabulary) error {
	str := fmt.Sprintf(`{"word" : %q, "meaning" : %q, "sentence" : %q}`, vocabulary.Word, vocabulary.Meaning, vocabulary.Sentence)

	res, err := r.esc.Index("vocabulary",
		strings.NewReader(str),
		r.esc.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Error(err)
		return err
	}

	defer res.Body.Close()

	return nil
}

func (r *repository) Filter(ctx context.Context, args SearchArgs) (*[]Vocabulary, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Hour)
	defer cancel()

	query := fmt.Sprintf(`{
		"size": %d,
		"from": %d,
		"query": {
			"query_string": {
				"fields": ["word", "meaning", "sentence"],
				"query": "%s",
				"minimum_should_match": "20%%"
			}
		}
	}`, args.Size, args.Page, args.SubWord)

	var sb strings.Builder
	sb.WriteString(query)
	reader := strings.NewReader(sb.String())

	res, err := r.esc.Search(
		r.esc.Search.WithContext(timeoutCtx),
		r.esc.Search.WithIndex("vocabulary"),
		r.esc.Search.WithBody(reader),
		r.esc.Search.WithTrackTotalHits(true),
		r.esc.Search.WithPretty(),
		r.esc.Search.WithFrom(args.Page),
		r.esc.Search.WithSize(args.Size),
	)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer res.Body.Close()

	responseByte, _ := io.ReadAll(res.Body)

	var searchModels SearchModel
	if err = json.Unmarshal(responseByte, &searchModels); err != nil {
		log.Error(err)
		return nil, err
	}

	vocabularies := searchModels.ToDTO()

	return vocabularies, nil
}
