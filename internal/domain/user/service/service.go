package service

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	dbModel "github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/gorm/generated/query"
	loggerSrv "github.com/umefy/go-web-app-template/internal/domain/logger/service"
	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	"github.com/umefy/go-web-app-template/internal/domain/user/model"
	"github.com/umefy/go-web-app-template/internal/domain/user/repository"
	"github.com/umefy/godash/sliceskit"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error)
	CreateUser(ctx context.Context, userCreateInput *model.UserCreateInput, tx *query.QueryTx) (*model.User, error)
	UpdateUser(ctx context.Context, id string, userUpdateInput *model.UserUpdateInput, tx *query.QueryTx) (*model.User, error)
}

type userService struct {
	loggerService  loggerSrv.Service
	userRepository repository.Repository
}

var _ Service = (*userService)(nil)

func NewService(loggerService loggerSrv.Service, userRepository repository.Repository) *userService {
	return &userService{loggerService: loggerService, userRepository: userRepository}
}

// GetUsers implements Service.
func (u *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	usersDb, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	return sliceskit.Map(usersDb, func(userDb *dbModel.User) *model.User {
		return model.User{}.CreateFromDbModel(userDb)
	}), nil
}

func (u *userService) GetUser(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.loggerService.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	user, err := u.userRepository.GetUser(ctx, userID)
	if err != nil {
		u.loggerService.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, err
	}
	return model.User{}.CreateFromDbModel(user), nil
}

func (u *userService) IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error) {
	return u.userRepository.IsUserEmailExists(ctx, email, tx)
}

func (u *userService) CreateUser(ctx context.Context, createUserInput *model.UserCreateInput, tx *query.QueryTx) (*model.User, error) {

	if err := createUserInput.Validate(); err != nil {
		return nil, err
	}

	exists, err := u.IsUserExists(ctx, createUserInput.Email, tx)
	if exists {
		return nil, userError.UserAlreadyExists
	}

	userDb, err := u.userRepository.CreateUser(ctx, createUserInput.MapToDbModel(), tx)
	if err != nil {
		return nil, err
	}
	return model.User{}.CreateFromDbModel(userDb), nil
}

// UpdateUser implements Service.
func (u *userService) UpdateUser(ctx context.Context, id string, updateUserInput *model.UserUpdateInput, tx *query.QueryTx) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.loggerService.ErrorContext(ctx, "UserService.UpdateUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	if err := updateUserInput.Validate(); err != nil {
		return nil, err
	}

	userDb, err := u.userRepository.UpdateUser(ctx, userID, updateUserInput.MapToDbModel(), tx)
	if err != nil {
		return nil, err
	}

	return model.User{}.CreateFromDbModel(userDb), nil
}
