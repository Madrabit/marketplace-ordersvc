package orders

import (
	"context"

	"github.com/google/uuid"
)

type RepoPrice struct {
	base map[uuid.UUID]PriceItem
}

func NewInmemPricing() *RepoPrice {
	base := map[uuid.UUID]PriceItem{
		uuid.MustParse("11111111-1111-1111-1111-111111111111"): {
			ProductId: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:      "Tea",
			UnitPrice: 1299,
		},
		uuid.MustParse("22222222-2222-2222-2222-222222222222"): {
			ProductId: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:      "Coffee",
			UnitPrice: 1599,
		},
	}
	return &RepoPrice{
		base: base,
	}
}

func (r *RepoPrice) PriceBatch(ctx context.Context, items []ItemQty) map[uuid.UUID]PriceItem {
	pricedItems := make(map[uuid.UUID]PriceItem, len(items))
	for _, i := range items {
		product := r.base[i.Id]
		p := PriceItem{
			product.ProductId,
			product.Name,
			product.UnitPrice,
		}
		pricedItems[i.Id] = p
	}
	return pricedItems
}
