package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	errFetchFailed    = errors.New("Failed to fetch data")
	errNotImplemented = errors.New("Not Implemented")
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

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return errNotImplemented
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return errNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, errNotImplemented
}
