package orders

import (
	"context"
	"github.com/google/uuid"
)

type PriceItem struct {
	ProductId uuid.UUID
	Name      string
	UnitPrice int64
}

type Pricing interface {
	PriceBatch(ctx context.Context, items []ItemQty) map[uuid.UUID]PriceItem
}
