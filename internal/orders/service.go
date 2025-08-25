package orders

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Service struct {
	repo   *Repository
	pricer Pricing
}

func NewService(repo *Repository, pricer Pricing) *Service {
	return &Service{repo, pricer}
}

type Repo interface {
	Create(ctx context.Context, o OrderRow, req CreatRequest) (OrderRow, error)
	Get(ctx context.Context, id uuid.UUID) (OrderRow, error)
}

func (s *Service) CreateOrder(ctx context.Context, req CreatRequest) (Response, error) {
	items := req.Items
	if len(items) == 0 {
		return Response{}, fmt.Errorf("no items provided")
	}
	pricedItems := s.pricer.PriceBatch(ctx, items)
	orderId := uuid.New()
	createdTime := time.Now()
	var grandTotal int64
	rowsOfItems := make([]ItemRow, 0, len(items))
	itemsResp := make([]ItemResponse, 0, len(items))
	for _, i := range items {
		p := pricedItems[i.Id]
		grandTotal += int64(i.Quantity) * p.UnitPrice
		item := ItemRow{
			i.Id,
			p.Name,
			i.Quantity,
			orderId,
			p.UnitPrice,
		}
		iResp := ItemResponse{
			p.Name,
			i.Quantity,
			p.UnitPrice,
		}
		itemsResp = append(itemsResp, iResp)
		rowsOfItems = append(rowsOfItems, item)
	}
	order := OrderRow{
		orderId,
		req.UserId,
		createdTime,
		New,
		grandTotal,
	}
	_, err := s.repo.Create(ctx, order, rowsOfItems)
	if err != nil {
		return Response{}, err
	}

	response := Response{
		req.UserId,
		grandTotal,
		createdTime,
		itemsResp,
	}
	return response, nil
}
