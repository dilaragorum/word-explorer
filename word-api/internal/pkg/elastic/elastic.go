package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/gommon/log"
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
		log.Error(err)
		return nil, err
	}

	return es, nil
}
