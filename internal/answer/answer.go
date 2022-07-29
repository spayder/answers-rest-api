package answer

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingAnswer = errors.New("failed to fetch answer by id")
	ErrPostingAnswer  = errors.New("failed to post answer")
	ErrUpdatingAnswer = errors.New("failed to update answer")
)

type Answer struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Store interface {
	GetAnswer(context.Context, string) (Answer, error)
	PostAnswer(context.Context, Answer) (Answer, error)
	UpdateAnswer(context.Context, string, Answer) (Answer, error)
	DeleteAnswer(context.Context, string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetAnswer(ctx context.Context, uuid string) (Answer, error) {
	fmt.Println("retrieving an answer")
	answer, err := s.Store.GetAnswer(ctx, uuid)
	if err != nil {
		fmt.Println(err)
		return Answer{}, ErrFetchingAnswer
	}
	return answer, nil
}

func (s *Service) UpdateAnswer(ctx context.Context, uuid string, answer Answer) (Answer, error) {
	ans, err := s.Store.UpdateAnswer(ctx, uuid, answer)
	if err != nil {
		fmt.Println(err)
		return Answer{}, ErrUpdatingAnswer
	}

	return ans, nil
}

func (s *Service) DeleteAnswer(ctx context.Context, uuid string) error {
	return s.Store.DeleteAnswer(ctx, uuid)
}

func (s *Service) PostAnswer(ctx context.Context, answer Answer) (Answer, error) {
	insertedAnswer, err := s.Store.PostAnswer(ctx, answer)

	if err != nil {
		fmt.Println(err)
		return Answer{}, ErrPostingAnswer
	}

	return insertedAnswer, nil
}
