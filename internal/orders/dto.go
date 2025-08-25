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

type ItemRow struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Quantity  int       `db:"quantity"`
	OrderId   uuid.UUID `db:"order_id"`
	UnitPrice int64     `db:"unit_price"`
}

type OrderRow struct {
	Id         uuid.UUID `db:"id"`
	UserId     uuid.UUID `db:"user_id"`
	CreatedAt  time.Time `db:"created_at"`
	Status     Status    `db:"status"`
	GrandTotal int64     `db:"grand_total"`
}

type ItemQty struct {
	Id       uuid.UUID `json:"id" validate:"required"`
	Quantity int       `json:"quantity" validate:"gte=1"`
}

type CreatRequest struct {
	UserId uuid.UUID `json:"user_id" validate:"required"`
	Items  []ItemQty `json:"items" validate:"min=1,dive"`
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
