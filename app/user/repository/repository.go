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
	userQuery := r.query.User.WithContext(ctx).Preload(r.query.User.Orders)
	user, err := userQuery.Where(r.query.User.ID.Eq(id)).First()

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
	userQuery := r.query.User.WithContext(ctx).Preload(r.query.User.Orders)
	users, err := userQuery.Find()
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
	userQuery := r.query.User.WithContext(ctx)
	err := userQuery.Create(user)

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
	userQuery := r.query.User.WithContext(ctx)
	info, err := userQuery.Where(r.query.User.ID.Eq(id)).Updates(user)

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
