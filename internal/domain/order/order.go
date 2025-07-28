package order

import (
	"time"
)

type Order struct {
	ID        int
	UserID    int
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
