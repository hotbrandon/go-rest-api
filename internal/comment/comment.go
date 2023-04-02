package comment

import (
	"context"
	"fmt"
)

type Comment struct {
	ID     int
	Slug   string
	Body   string
	Author string
}

type DataStore interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store DataStore
}

// returns a new pointer to a Service
func NewService(store DataStore) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("comment received")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		return Comment{}, err
	}
	return cmt, nil
}
