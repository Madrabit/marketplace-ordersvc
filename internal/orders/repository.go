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
	mu    sync.RWMutex
	m     map[uuid.UUID]OrderRow
	items map[uuid.UUID]ItemRow
}

func NewRepository() *Repository {
	return &Repository{m: make(map[uuid.UUID]OrderRow),
		items: make(map[uuid.UUID]ItemRow)}
}

func (r *Repository) Create(ctx context.Context, o OrderRow, items []ItemRow) (OrderRow, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.m[o.Id] = o
	for _, item := range items {
		r.items[o.Id] = item
	}
	return o, nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (OrderRow, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o, ok := r.m[id]
	if !ok {
		return OrderRow{}, ErrorNotFound
	}
	return o, nil
}
