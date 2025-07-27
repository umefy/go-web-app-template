package repo

import (
	"context"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
)

type Repository interface {
	FindUser(ctx context.Context, id int) (*userDomain.User, error)
	FindUserTx(ctx context.Context, id int, tx *query.QueryTx) (*userDomain.User, error)
	FindUsers(ctx context.Context) ([]*userDomain.User, error)
	FindUsersTx(ctx context.Context, tx *query.QueryTx) ([]*userDomain.User, error)
	CreateUser(ctx context.Context, user *userDomain.User, tx *query.QueryTx) (*userDomain.User, error)
	UpdateUser(ctx context.Context, id int, user *userDomain.User, tx *query.QueryTx) (*userDomain.User, error)
	IsUserEmailExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error)

	FindUserWithOrders(ctx context.Context, id int) (*userDomain.UserWithOrder, error)
}
