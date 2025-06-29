package repository

import (
	"context"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
)

type Repository interface {
	GetUser(ctx context.Context, id int) (*dbModel.User, error)
	GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*dbModel.User, error)
	GetUsers(ctx context.Context) ([]*dbModel.User, error)
	GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*dbModel.User, error)
	CreateUser(ctx context.Context, user *dbModel.User, tx *query.QueryTx) (*dbModel.User, error)
	UpdateUser(ctx context.Context, id int, user *dbModel.User, tx *query.QueryTx) (*dbModel.User, error)
}
