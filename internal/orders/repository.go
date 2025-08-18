package orders

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
)

var (
	ErrorNotFound = errors.New("not found")
)

type Repository struct {
	mu sync.RWMutex
	m  map[uuid.UUID]Order
}

func NewRepository() *Repository {
	return &Repository{m: make(map[uuid.UUID]Order)}
}

func (r *Repository) Create(ctx context.Context, o Order) (Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[o.Id] = o
	return o, nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o, ok := r.m[id]
	if !ok {
		return Order{}, ErrorNotFound
	}
	return o, nil
}
