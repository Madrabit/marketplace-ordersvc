package pricing

import "github.com/google/uuid"

type Pricing struct {
	ProductId uuid.UUID
	Name      string
	UnitPrice int64
}
