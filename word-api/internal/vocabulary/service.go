package vocabulary

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, vocabulary Vocabulary) error
	Filter(ctx context.Context, args SearchArgs) (FilterResult, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, vocabulary Vocabulary) error {
	return s.repository.Create(ctx, vocabulary)
}

func (s *service) Filter(ctx context.Context, args SearchArgs) (FilterResult, error) {
	vocabularies, err := s.repository.Filter(ctx, args)
	if err != nil {
		return FilterResult{}, err
	}

	return vocabularies, nil
}
