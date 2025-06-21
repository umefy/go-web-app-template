package repository

import (
	"context"
	"errors"
	"log/slog"

	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	userError "github.com/umefy/go-web-app-template/app/user/error"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	"github.com/umefy/go-web-app-template/internal/server/httpserver/middlewares"
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

	if errors.Is(err, db.ErrRecordNotFound) {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		return nil, userError.UserNotFound
	}

	if err != nil {
		r.loggerService.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
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
	tx := ctx.Value(middlewares.Transaction).(*query.QueryTx)

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

	// Use the transaction to get the updated user
	updatedUser, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(id)).Preload(tx.User.Orders).First()
	r.loggerService.DebugContext(ctx, "Get User", slog.Any("user", updatedUser.Name))
	if err != nil {
		r.loggerService.ErrorContext(ctx, "Get User error", slog.String("error", err.Error()))
		return nil, err
	}

	return updatedUser, nil
}
