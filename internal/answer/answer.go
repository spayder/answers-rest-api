package answer

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingAnswer = errors.New("failed to fetch answer by id")
	ErrNotImplemented = errors.New("not implemented")
)

type Answer struct {
	ID    string
	Key   string
	Value string
}

type Store interface {
	GetAnswer(ctx context.Context, id string) (Answer, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetAnswer(ctx context.Context, id string) (Answer, error) {
	fmt.Println("retrieving an answer")
	answer, err := s.Store.GetAnswer(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Answer{}, ErrFetchingAnswer
	}
	return answer, nil
}

func (s *Service) UpdateAnswer(ctx context.Context, answer Answer) error {
	return ErrNotImplemented
}

func (s *Service) DeleteAnswer(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateAnswer(ctx context.Context, answer Answer) (Answer, error) {
	return Answer{}, ErrNotImplemented
}
