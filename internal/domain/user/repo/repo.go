package repo

import (
	"context"

	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	"github.com/umefy/go-web-app-template/pkg/pagination"
)

type Repository interface {
	FindUser(ctx context.Context, id int) (*userDomain.User, error)
	FindUsers(ctx context.Context, p pagination.Pagination) ([]*userDomain.User, *pagination.PaginationMetadata, error)
	CreateUser(ctx context.Context, user *userDomain.User) (*userDomain.User, error)
	UpdateUser(ctx context.Context, id int, user *userDomain.User) (*userDomain.User, error)
	IsUserEmailExists(ctx context.Context, email string) (bool, error)
	FindUserWithOrders(ctx context.Context, id int) (*userDomain.UserWithOrder, error)
}
