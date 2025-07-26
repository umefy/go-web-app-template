package database

import (
	"context"
	"errors"
	"log/slog"

	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userRepo "github.com/umefy/go-web-app-template/internal/domain/user/repository"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/go-web-app-template/pkg/null"

	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	db "github.com/umefy/go-web-app-template/pkg/db/gormdb"
)

type UserRepository struct {
	Logger  logger.Logger
	dbQuery *query.Query
}

var _ userRepo.Repository = (*UserRepository)(nil)

func NewUserRepository(dbQuery *query.Query, logger logger.Logger) *UserRepository {
	return &UserRepository{Logger: logger, dbQuery: dbQuery}
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	userQuery := r.dbQuery.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Preload(userQuery.Orders).First()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserTx(ctx context.Context, id int, tx *query.QueryTx) (*model.User, error) {
	userQuery := tx.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Preload(userQuery.Orders).First()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUser", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*model.User, error) {
	userQuery := r.dbQuery.User
	users, err := userQuery.WithContext(ctx).Preload(userQuery.Orders).Where(userQuery.Age.Gt(null.ValueFrom(1))).Order(userQuery.ID.Asc()).Find()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUsersTx(ctx context.Context, tx *query.QueryTx) ([]*model.User, error) {
	userQuery := tx.User
	users, err := userQuery.WithContext(ctx).Preload(userQuery.Orders).Where(userQuery.Age.Gt(null.ValueFrom(1))).Order(userQuery.ID.Asc()).Find()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.GetUsers", slog.String("error", err.Error()))
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, userError.UserNotFound
		}
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User, tx *query.QueryTx) (*model.User, error) {
	userQuery := tx.User
	err := userQuery.WithContext(ctx).Create(user)

	if errors.Is(err, db.ErrRecordNotFound) {
		r.Logger.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, userError.UserNotFound
	}

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.CreateUser", slog.String("error", err.Error()))
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user *model.User, tx *query.QueryTx) (*model.User, error) {

	userQuery := tx.User
	info, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(id)).Updates(user)

	if err != nil {
		r.Logger.ErrorContext(ctx, "User update error", slog.String("error", err.Error()))
		return nil, err
	}

	if info.RowsAffected == 0 {
		r.Logger.ErrorContext(ctx, "User update error", slog.String("error", "user not found"))
		return nil, userError.UserNotFound
	}

	updatedUser, err := r.GetUserTx(ctx, id, tx)
	if err != nil {
		r.Logger.ErrorContext(ctx, "Get User error", slog.String("error", err.Error()))
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepository) IsUserEmailExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error) {
	userQuery := tx.User
	count, err := userQuery.WithContext(ctx).Where(userQuery.Email.Eq(null.ValueFrom(email))).Count()

	if err != nil {
		r.Logger.ErrorContext(ctx, "UserRepository.IsUserEmailExists", slog.String("error", err.Error()))
		return false, err
	}
	return count > 0, nil
}
