package repository

import (
	"context"
	"errors"
	"log/slog"

	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/middleware"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
	"github.com/umefy/go-web-app-template/pkg/null"
)

type Repository interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error)
}

type userRepository struct {
	loggerService loggerSrv.Service
	dbQuery       *query.Query
}

var _ Repository = (*userRepository)(nil)

func NewUserRepository(dbQuery *query.Query, loggerService loggerSrv.Service) *userRepository {
	return &userRepository{loggerService: loggerService, dbQuery: dbQuery}
}

func (r *userRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	userQuery := r.dbQuery.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Preload(userQuery.Orders).First()

	if err != nil {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserTx(ctx context.Context, id int) (*model.User, error) {
	tx := ctx.Value(middleware.TransactionCtxKey).(*query.QueryTx)
	userQuery := tx.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Preload(userQuery.Orders).First()

	if err != nil {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	userQuery := r.dbQuery.User
	users, err := userQuery.WithContext(ctx).Preload(userQuery.Orders).Where(userQuery.Age.Gt(null.ValueFrom(1))).Order(userQuery.ID.Asc()).Find()

	if errors.Is(err, db.ErrRecordNotFound) {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		return nil, userError.UserNotFound
	}
	if err != nil {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	userQuery := r.dbQuery.User
	err := userQuery.WithContext(ctx).Create(user)

	if errors.Is(err, db.ErrRecordNotFound) {
		r.loggerService.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, userError.UserNotFound
	}

	if err != nil {
		r.loggerService.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	tx, err := middleware.GetTransaction(ctx)
	if err != nil {
		return nil, err
	}

	userQuery := tx.User
	info, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Updates(user)

	if err != nil {
		r.loggerService.ErrorContext(ctx, "User update error", slog.String("error", err.Error()))
		return nil, err
	}

	if info.RowsAffected == 0 {
		r.loggerService.ErrorContext(ctx, "User update error", slog.String("error", "user not found"))
		return nil, userError.UserNotFound
	}

	updatedUser, err := r.GetUserTx(ctx, id)
	if err != nil {
		r.loggerService.ErrorContext(ctx, "Get User error", slog.String("error", err.Error()))
		return nil, err
	}

	return updatedUser, nil
}
