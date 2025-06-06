package service

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	loggerSrv "github.com/umefy/go-web-app-template/app/logger/service"
	"github.com/umefy/go-web-app-template/app/user/repository"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
)

type Service interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error)
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
	users, err := u.userRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
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
	return user, nil
}

func (u *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements Service.
func (u *userService) UpdateUser(ctx context.Context, id string, user *model.User) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		u.loggerService.ErrorContext(ctx, "UserService.UpdateUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("invalid user id")
	}

	user, err = u.userRepository.UpdateUser(ctx, userID, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
