package main

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	dbModel "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"
	"github.com/umefy/go-web-app-template/pkg/null"
)

const userCount = 100

func seedUsers(query *query.Query) ([]*dbModel.User, error) {
	gofakeit.Seed(666)

	users := make([]*dbModel.User, userCount)

	for i := range users {
		users[i] = &dbModel.User{
			Email: null.ValueFrom(gofakeit.Email()),
			Age:   null.ValueFrom(gofakeit.IntRange(0, 60)),
		}
	}

	err := query.User.WithContext(context.Background()).CreateInBatches(
		users,
		userCount,
	)
	if err != nil {
		return nil, err
	}

	return users, nil
}
