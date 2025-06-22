package repository

import (
	"context"

	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
)

type Repository interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User, tx *query.QueryTx) (*model.User, error)
	UpdateUser(ctx context.Context, id int, user *model.User, tx *query.QueryTx) (*model.User, error)
}
