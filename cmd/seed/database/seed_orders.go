package main

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	dbModel "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/godash/sliceskit"
)

const orderCount = 1000

func seedOrders(query *query.Query, users []*dbModel.User) ([]*dbModel.Order, error) {
	gofakeit.Seed(666)

	orders := make([]*dbModel.Order, orderCount)

	userIds := sliceskit.Map(users, func(user *dbModel.User) int {
		return user.ID
	})

	for i := range orders {
		orders[i] = &dbModel.Order{
			UserID:      userIds[gofakeit.IntRange(0, len(userIds)-1)],
			AmountCents: null.ValueFrom(int64(gofakeit.IntRange(1, 100_000))),
		}
	}

	err := query.Order.WithContext(context.Background()).CreateInBatches(
		orders,
		orderCount,
	)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
