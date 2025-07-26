package service

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	"github.com/umefy/go-web-app-template/internal/domain/user/repository"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	"github.com/umefy/godash/sliceskit"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*userDomain.User, error)
	GetUser(ctx context.Context, id string) (*userDomain.User, error)
	IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error)
	CreateUser(ctx context.Context, userCreateInput *userDomain.UserCreateInput, tx *query.QueryTx) (*userDomain.User, error)
	UpdateUser(ctx context.Context, id string, userUpdateInput *userDomain.UserUpdateInput, tx *query.QueryTx) (*userDomain.User, error)
}

type userService struct {
	logger         logger.Logger
	userRepository repository.Repository
}

var _ Service = (*userService)(nil)

func NewService(logger logger.Logger, userRepository repository.Repository) *userService {
	return &userService{logger: logger, userRepository: userRepository}
}

// GetUsers implements Service.
func (u *userService) GetUsers(ctx context.Context) ([]*userDomain.User, error) {
	usersDb, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return sliceskit.Map(usersDb, func(userDb *dbModel.User) *userDomain.User {
		return userDomain.User{}.CreateFromDbModel(userDb)
	}), nil
}

func (u *userService) GetUser(ctx context.Context, id string) (*userDomain.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	user, err := u.userRepository.GetUser(ctx, userID)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, err
	}
	return userDomain.User{}.CreateFromDbModel(user), nil
}

func (u *userService) IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error) {
	return u.userRepository.IsUserEmailExists(ctx, email, tx)
}

func (u *userService) CreateUser(ctx context.Context, createUserInput *userDomain.UserCreateInput, tx *query.QueryTx) (*userDomain.User, error) {

	if err := createUserInput.Validate(); err != nil {
		return nil, err
	}

	if exists, err := u.IsUserExists(ctx, createUserInput.Email, tx); err != nil {
		return nil, err
	} else if exists {
		return nil, userError.UserAlreadyExists
	}

	userDb, err := u.userRepository.CreateUser(ctx, createUserInput.MapToDbModel(), tx)
	if err != nil {
		return nil, err
	}
	return userDomain.User{}.CreateFromDbModel(userDb), nil
}

// UpdateUser implements Service.
func (u *userService) UpdateUser(ctx context.Context, id string, updateUserInput *userDomain.UserUpdateInput, tx *query.QueryTx) (*userDomain.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.UpdateUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	if err := updateUserInput.Validate(); err != nil {
		return nil, err
	}

	userDb, err := u.userRepository.UpdateUser(ctx, userID, updateUserInput.MapToDbModel(), tx)
	if err != nil {
		return nil, err
	}

	return userDomain.User{}.CreateFromDbModel(userDb), nil
}
