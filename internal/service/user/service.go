package user

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/umefy/go-web-app-template/gorm/generated/query"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	userError "github.com/umefy/go-web-app-template/internal/domain/user/error"
	"github.com/umefy/go-web-app-template/internal/domain/user/repo"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*userDomain.User, error)
	GetUser(ctx context.Context, id string) (*userDomain.User, error)
	IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error)
	CreateUser(ctx context.Context, userCreateInput *UserCreateInput, tx *query.QueryTx) (*userDomain.User, error)
	UpdateUser(ctx context.Context, id string, userUpdateInput *UserUpdateInput, tx *query.QueryTx) (*userDomain.User, error)
}

type userService struct {
	logger         logger.Logger
	userRepository repo.Repository
}

var _ Service = (*userService)(nil)

func NewService(logger logger.Logger, userRepository repo.Repository) *userService {
	return &userService{logger: logger, userRepository: userRepository}
}

// GetUsers implements Service.
func (u *userService) GetUsers(ctx context.Context) ([]*userDomain.User, error) {
	usersDb, err := u.userRepository.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	return usersDb, nil
}

func (u *userService) GetUser(ctx context.Context, id string) (*userDomain.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	user, err := u.userRepository.FindUser(ctx, userID)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.GetUser", slog.String("error", err.Error()))
		return nil, err
	}
	return user, nil
}

func (u *userService) IsUserExists(ctx context.Context, email string, tx *query.QueryTx) (bool, error) {
	return u.userRepository.IsUserEmailExists(ctx, email, tx)
}

func (u *userService) CreateUser(ctx context.Context, createUserInput *UserCreateInput, tx *query.QueryTx) (*userDomain.User, error) {

	if err := createUserInput.Validate(); err != nil {
		return nil, err
	}

	if exists, err := u.IsUserExists(ctx, createUserInput.Email, tx); err != nil {
		return nil, err
	} else if exists {
		return nil, userError.UserAlreadyExists
	}

	user := createUserInput.MapToDomainUser()
	userDb, err := u.userRepository.CreateUser(ctx, user, tx)
	if err != nil {
		return nil, err
	}
	return userDb, nil
}

// UpdateUser implements Service.
func (u *userService) UpdateUser(ctx context.Context, id string, updateUserInput *UserUpdateInput, tx *query.QueryTx) (*userDomain.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.logger.ErrorContext(ctx, "UserService.UpdateUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	if err := updateUserInput.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userRepository.FindUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	updatedUser := updateDomainUser(user, updateUserInput)

	userDb, err := u.userRepository.UpdateUser(ctx, userID, updatedUser, tx)
	if err != nil {
		return nil, err
	}

	return userDb, nil
}
