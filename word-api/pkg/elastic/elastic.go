package elastic

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"net/http"
)

func NewClient() (*elasticsearch.Client, error) {
	config := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: &http.Transport{
			MaxConnsPerHost: 100,
		},
	}

	es, err := elasticsearch.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("error initializing elastic client %v", err)
	}

	return es, nil
}
