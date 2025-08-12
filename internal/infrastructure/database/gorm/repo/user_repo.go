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
	"github.com/umefy/godash/sliceskit"

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

func (r *UserRepo) FindUsers(ctx context.Context) ([]*userDomain.User, error) {
	userQuery := r.dbQuery.User
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
	type UserOrderRow struct {
		ID               int
		Email            string
		Age              int
		CreatedAt        time.Time
		UpdatedAt        time.Time
		OrderID          int
		OrderAmountCents int64
		OrderCreatedAt   time.Time
		OrderUpdatedAt   time.Time
	}

	userQuery := r.dbQuery.User
	orderQuery := r.dbQuery.Order

	var userOrderRows []UserOrderRow
	err := userQuery.WithContext(ctx).Select(
		userQuery.ID,
		userQuery.Email,
		userQuery.Age,
		userQuery.CreatedAt,
		userQuery.UpdatedAt,
		orderQuery.ID.As("order_id"),
		orderQuery.AmountCents.As("order_amount_cents"),
		orderQuery.CreatedAt.As("order_created_at"),
		orderQuery.UpdatedAt.As("order_updated_at"),
	).LeftJoin(orderQuery, userQuery.ID.EqCol(orderQuery.UserID)).Where(userQuery.ID.Eq(id)).Scan(&userOrderRows)
	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.FindUserWithOrders", slog.String("error", err.Error()))
		return nil, err
	}

	if len(userOrderRows) == 0 {
		return nil, userError.UserNotFound
	}

	user := &userDomain.UserWithOrder{
		User: (userDomain.User{
			ID:        userOrderRows[0].ID,
			Email:     userOrderRows[0].Email,
			Age:       userOrderRows[0].Age,
			CreatedAt: userOrderRows[0].CreatedAt,
			UpdatedAt: userOrderRows[0].UpdatedAt,
		}),
		Orders: []orderDomain.Order{},
	}

	for _, userOrderRow := range userOrderRows {
		if userOrderRow.OrderID == 0 {
			continue
		}

		user.Orders = append(user.Orders, orderDomain.Order{
			ID:          userOrderRow.OrderID,
			AmountCents: userOrderRow.OrderAmountCents,
			CreatedAt:   userOrderRow.OrderCreatedAt,
			UpdatedAt:   userOrderRow.OrderUpdatedAt,
		})
	}

	return user, nil
}
