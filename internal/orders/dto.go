package orders

import (
	"github.com/google/uuid"
	"time"
)

type Status string

const (
	New       Status = "new"
	Shipped   Status = "shipped"
	Delivered Status = "delivered"
	Paid      Status = "paid"
	Canceled  Status = "canceled"
)

type Item struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Quantity  int       `db:"quantity"`
	OrderId   uuid.UUID `db:"order_id"`
	UnitPrice int64     `db:"unit_price"`
}

type Order struct {
	Id         uuid.UUID `db:"id"`
	UserId     uuid.UUID `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	Status     Status    `db:"status"`
	GrandTotal int64     `db:"grand_total"`
}

type ItemRequest struct {
	Id       int64 `json:"id" validate:"required"`
	Quantity int   `json:"quantity" validate:"gte=1"`
}

type CreatRequest struct {
	UserId uuid.UUID     `json:"user_id" validate:"required"`
	Items  []ItemRequest `json:"items" validate:"min=1,dive"`
}

type ItemResponse struct {
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

type Response struct {
	UserId     uuid.UUID      `json:"user_id"`
	GrandTotal int64          `json:"grand_total"`
	Created    time.Time      `json:"created"`
	Items      []ItemResponse `json:"items"`
}

func ToOrder(req *CreatRequest) Order {
	items := make([]Item, len(req.Items))
	for _, i := range req.Items {
		item := Item{
			Id:       i.Id,
			Quantity: i.Quantity,
		}
		items = append(items, item)
	}
	return Order{
		UserId:     req.UserId,
		CreatedAt:  time.Now(),
		Status:     New,
		GrandTotal: 0,
	}
}
