package repo

import (
	"context"
	"errors"
	"log/slog"
	"time"

	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repo"
	dbContext "github.com/umefy/go-web-app-template/internal/infrastructure/database/ctx"
	dbModel "github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/infrastructure/database/gorm/repo/mapping"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/go-web-app-template/pkg/null"
	"github.com/umefy/go-web-app-template/pkg/pagination"
	"github.com/umefy/godash/sliceskit"
	"gorm.io/datatypes"
	"gorm.io/gen/field"

	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

type UserRepo struct {
	Logger  logger.Logger
	dbQuery *query.Query
}

var _ userRepo.Repository = (*UserRepo)(nil)

func NewUserRepository(dbQuery *query.Query, logger logger.Logger) *UserRepo {
	return &UserRepo{Logger: logger, dbQuery: dbQuery}
}

func (r *UserRepo) FindUser(ctx context.Context, id int) (*userDomain.User, error) {
	userQuery := r.dbQuery.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).First()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}

	return mapping.DbModelToDomainUser(user), nil
}

func (r *UserRepo) FindUserTx(ctx context.Context, id int) (*userDomain.User, error) {
	tx := ctx.Value(dbContext.TransactionCtxKey).(*query.QueryTx)
	userQuery := tx.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).First()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}

	return mapping.DbModelToDomainUser(user), nil
}

func (r *UserRepo) FindUsers(ctx context.Context, p pagination.Pagination) ([]*userDomain.User, *pagination.PaginationMetadata, error) {
	userQuery := r.dbQuery.User
	users, err := userQuery.WithContext(ctx).Order(userQuery.ID.Asc()).Offset(p.Offset).Limit(p.PageSize + 1).Find()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, nil, userError.UserNotFound
		}
		return nil, nil, err
	}

	hasMore := len(users) > p.PageSize

	if hasMore {
		users = users[:p.PageSize]
	}

	var total *int64 = nil
	metadata := pagination.NewPaginationMetadata(p.Offset, p.PageSize, len(users), hasMore, total)
	if p.IncludeTotal {
		totalCount, err := userQuery.WithContext(ctx).Count()
		if err != nil {
			return nil, nil, err
		}
		metadata.Total = &totalCount
	}

	return sliceskit.Map(users, func(user *dbModel.User) *userDomain.User {
		return mapping.DbModelToDomainUser(user)
	}), &metadata, nil
}

func (r *UserRepo) FindUsersTx(ctx context.Context) ([]*userDomain.User, error) {
	tx := ctx.Value(dbContext.TransactionCtxKey).(*query.QueryTx)
	userQuery := tx.User
	users, err := userQuery.WithContext(ctx).Order(userQuery.ID.Asc()).Find()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return sliceskit.Map(users, func(user *dbModel.User) *userDomain.User {
		return mapping.DbModelToDomainUser(user)
	}), nil
}

func (r *UserRepo) CreateUser(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
	tx := ctx.Value(dbContext.TransactionCtxKey).(*query.QueryTx)
	userQuery := tx.User
	dbModel := mapping.DomainUserToDbModel(user)
	err := userQuery.WithContext(ctx).Create(dbModel)

	if errors.Is(err, db.ErrRecordNotFound) {
		r.Logger.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, userError.UserNotFound
	}

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, err
	}

	return mapping.DbModelToDomainUser(dbModel), nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id int, user *userDomain.User) (*userDomain.User, error) {

	tx := ctx.Value(dbContext.TransactionCtxKey).(*query.QueryTx)
	userQuery := tx.User

	dbModel := mapping.DomainUserToDbModel(user)
	info, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id), userQuery.Version.Eq(user.Version)).Updates(dbModel)

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.UpdateUser", slog.String("error", err.Error()))
		return nil, err
	}

	if info.RowsAffected == 0 {
		// When RowsAffected is 0 with optimistic locking, it means either:
		// 1. User doesn't exist, or 2. Version mismatch (optimistic lock conflict)
		// Since we're using optimistic locking, assume it's a version conflict
		// But service level should already handle 1st case, so we don't need to return user not found error
		r.Logger.ErrorContext(ctx, "UserRepository.UpdateUser", slog.String("error", "user update conflict - version mismatch"))
		return nil, userError.UserUpdateConflict
	}

	return mapping.DbModelToDomainUser(dbModel), nil
}

func (r *UserRepo) IsUserEmailExists(ctx context.Context, email string) (bool, error) {
	userQuery := r.dbQuery.User
	count, err := userQuery.WithContext(ctx).Where(userQuery.Email.Eq(null.ValueFrom(email))).Count()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.IsUserEmailExists", slog.String("error", err.Error()))
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepo) FindUserWithOrders(ctx context.Context, id int) (*userDomain.UserWithOrder, error) {

	type Order struct {
		OrderID     int       `json:"order_id"`
		AmountCents int64     `json:"amount_cents"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	type UserOrderRow struct {
		ID        int
		Email     string
		Age       int
		CreatedAt time.Time
		UpdatedAt time.Time
		Orders    datatypes.JSONType[[]Order]
	}

	userQuery := r.dbQuery.User
	orderQuery := r.dbQuery.Order.As("o")

	orderAgg := field.NewUnsafeFieldRaw(
		`
			COALESCE(
					jsonb_agg(
						jsonb_build_object(
							'order_id', o.id,
							'amount_cents', o.amount_cents,
							'created_at', o.created_at,
							'updated_at', o.updated_at
						)
						order by o.created_at
					)
					,
					'[]'::jsonb
			)
		`,
	).As("orders")

	var userOrderRow UserOrderRow
	err := userQuery.WithContext(ctx).Select(userQuery.ID, userQuery.Email, userQuery.Age, userQuery.CreatedAt, userQuery.UpdatedAt, orderAgg).Join(orderQuery, userQuery.ID.EqCol(orderQuery.UserID)).Where(userQuery.ID.Eq(id)).Group(userQuery.ID).Scan(&userOrderRow)

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.FindUserWithOrders", slog.String("error", err.Error()))
		return nil, err
	}

	if userOrderRow.ID == 0 {
		return nil, userError.UserNotFound
	}

	user := &userDomain.UserWithOrder{
		User: userDomain.User{
			ID:        userOrderRow.ID,
			Email:     userOrderRow.Email,
			Age:       userOrderRow.Age,
			CreatedAt: userOrderRow.CreatedAt,
			UpdatedAt: userOrderRow.UpdatedAt,
		},
		Orders: sliceskit.Map(userOrderRow.Orders.Data(), func(order Order) orderDomain.Order {
			return orderDomain.Order{
				ID:          order.OrderID,
				AmountCents: order.AmountCents,
				CreatedAt:   order.CreatedAt,
				UpdatedAt:   order.UpdatedAt,
			}
		}),
	}

	return user, nil
}
