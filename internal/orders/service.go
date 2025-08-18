package orders

import (
	"context"
	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

type Repo interface {
	Create(ctx context.Context, o Order) (Order, error)
	Get(ctx context.Context, id uuid.UUID) (Order, error)
}

func (s *Service) CreateOrder(ctx context.Context, req CreatRequest) (Order, error) {

}
