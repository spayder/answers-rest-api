package answer

import (
	"context"
	"fmt"
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
		return Answer{}, err
	}
	return answer, nil
}
