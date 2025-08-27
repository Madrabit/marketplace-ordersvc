package orders

import (
	"context"
	"github.com/google/uuid"
)

type PriceItem struct {
	ProductId uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	UnitPrice int64     `json:"unit_price"`
}

type Pricing interface {
	PriceBatch(ctx context.Context, items []ItemQty) map[uuid.UUID]PriceItem
}
