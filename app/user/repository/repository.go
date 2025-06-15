package repository

import (
	"context"
	"errors"
	"log/slog"

	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	userError "github.com/umefy/go-web-app-template/app/user/error"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
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
	query         *query.Query
}

var _ Repository = (*userRepository)(nil)

func NewUserRepository(db *db.DB, loggerService loggerSrv.Service) *userRepository {
	return &userRepository{loggerService: loggerService, query: query.Use(db)}
}

func (r *userRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	userQuery := r.query.User
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
	userQuery := r.query.User
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
	userQuery := r.query.User
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
	userQuery := r.query.User
	info, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Updates(user)

	if err != nil {
		r.loggerService.ErrorContext(ctx, "User update error", slog.String("error", err.Error()))
		return nil, err
	}

	if info.RowsAffected == 0 {
		r.loggerService.ErrorContext(ctx, "User update error", slog.String("error", "user not found"))
		return nil, userError.UserNotFound
	}

	user, err = r.GetUser(ctx, id)
	if err != nil {
		r.loggerService.ErrorContext(ctx, "Get User error", slog.String("error", err.Error()))
		return nil, err
	}

	return user, nil
}
