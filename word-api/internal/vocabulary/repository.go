package vocabulary

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"io"
	"net/http"
	"strings"
	"time"
)

const indexName = "vocabulary"

type repository struct {
	esc *elasticsearch.Client
}

func NewRepository(esc *elasticsearch.Client) Repository {
	return &repository{esc: esc}
}

func (r *repository) Create(_ context.Context, vocabulary Vocabulary) error {
	jsonBytes, _ := json.Marshal(vocabulary)

	res, err := r.esc.Index(indexName,
		bytes.NewBuffer(jsonBytes),
		r.esc.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("error creating vocabulary on es %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("error creating vocabulary %d", res.StatusCode)
	}

	return nil
}

func (r *repository) Filter(ctx context.Context, args SearchArgs) (FilterResult, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := fmt.Sprintf(`{
			"size": %d,
			"from": %d
		}`, args.Size, args.Page)

	if args.SearchWord != "" {
		query = fmt.Sprintf(`{
			"size": %d,
			"from": %d,
			"query": {
				"query_string": {
					"fields": ["word", "meaning", "sentence"],
					"query": %q,
					"minimum_should_match": "20%%"
				}
			}
		}`, args.Size, args.Page, args.SearchWord)
	}

	var sb strings.Builder
	sb.WriteString(query)
	reader := strings.NewReader(sb.String())

	res, err := r.esc.Search(
		r.esc.Search.WithContext(timeoutCtx),
		r.esc.Search.WithIndex(indexName),
		r.esc.Search.WithBody(reader),
		r.esc.Search.WithTrackTotalHits(true),
		r.esc.Search.WithFrom(args.Page),
		r.esc.Search.WithSize(args.Size),
	)

	if err != nil {
		return FilterResult{}, fmt.Errorf("error when filtering vocabularies: %v", err.Error())
	}

	defer res.Body.Close()

	responseByte, _ := io.ReadAll(res.Body)

	var searchModels SearchModel
	if err = json.Unmarshal(responseByte, &searchModels); err != nil {
		return FilterResult{}, fmt.Errorf("error when unmarshall response: %v", err.Error())
	}

	vocabularies := searchModels.ToDTO()

	return vocabularies, nil
}
