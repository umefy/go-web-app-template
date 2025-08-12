package order

import (
	"time"
)

type Order struct {
	ID          int
	UserID      int
	AmountCents int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
