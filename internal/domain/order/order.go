package model

import (
	"time"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
)

type Order struct {
	ID        int
	UserID    int
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o Order) CreateFromDbModel(dbModel *dbModel.Order) *Order {
	return &Order{
		ID:        dbModel.ID,
		UserID:    dbModel.UserID.ValueOrZero(),
		Amount:    dbModel.Amount.ValueOrZero(),
		CreatedAt: dbModel.CreatedAt,
		UpdatedAt: dbModel.UpdatedAt,
	}
}
