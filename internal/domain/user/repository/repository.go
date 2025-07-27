package repository

import (
	"context"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
)

type Repository interface {
	GetUser(ctx context.Context, id int) (*userDomain.User, error)
	GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*userDomain.User, error)
	GetUsers(ctx context.Context) ([]*userDomain.User, error)
	GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*userDomain.User, error)
	CreateUser(ctx context.Context, user *userDomain.User, tx *query.QueryTx) (*userDomain.User, error)
	UpdateUser(ctx context.Context, id int, user *userDomain.User, tx *query.QueryTx) (*userDomain.User, error)
	IsUserEmailExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error)
}
