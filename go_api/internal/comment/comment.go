package comment

import (
	"context"
	"fmt"
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}
type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("fetch comments")
	return Comment{}, nil
}
func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, nil
}
