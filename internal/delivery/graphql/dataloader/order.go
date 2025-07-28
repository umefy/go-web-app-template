package dataloader

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/umefy/go-web-app-template/internal/delivery/graphql/mapping"
	"github.com/umefy/go-web-app-template/internal/delivery/graphql/model"
	orderDomain "github.com/umefy/go-web-app-template/internal/domain/order"
	"github.com/umefy/go-web-app-template/internal/infrastructure/logger"
	orderSrv "github.com/umefy/go-web-app-template/internal/service/order"
	"github.com/umefy/godash/sliceskit"
	"github.com/vikstrous/dataloadgen"
)

func createOrderLoader(ctx context.Context, logger logger.Logger, orderService orderSrv.Service) *dataloadgen.Loader[string, []*model.Order] {
	// Dataloader must return a slice with same order of the userIDs input
	return dataloadgen.NewLoader(func(ctx context.Context, userIDs []string) ([][]*model.Order, []error) {
		userIDsInt, err := sliceskit.MapWithFuncErr(userIDs, func(id string) (int, error) {
			return strconv.Atoi(id)
		})
		if err != nil {
			return nil, []error{err}
		}
		orders, err := orderService.GetOrdersByUserIDs(ctx, userIDsInt)
		if err != nil {
			return nil, []error{err}
		}

		userOrderMap := make(map[string][]*orderDomain.Order)
		for _, order := range orders {
			userOrderMap[strconv.Itoa(order.UserID)] = append(userOrderMap[strconv.Itoa(order.UserID)], order)
		}

		result := make([][]*model.Order, len(userIDs))
		for i, userID := range userIDs {
			if orders, ok := userOrderMap[userID]; ok {
				result[i] = sliceskit.Map(orders, mapping.OrderModelToGraphqlOrder)
			} else {
				result[i] = []*model.Order{}
			}
		}
		return result, nil
	})
}

func GetOrdersByUserIDs(ctx context.Context, userIDs []string) ([][]*model.Order, error) {
	loader := ctx.Value(LoaderCtxKey).(*Loaders)
	orders, err := loader.OrderLoader.LoadAll(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrdersByUserID(ctx context.Context, userID string, logger logger.Logger) ([]*model.Order, error) {
	loader := ctx.Value(LoaderCtxKey).(*Loaders)
	orders, err := loader.OrderLoader.Load(ctx, userID)
	if err != nil {
		logger.ErrorContext(ctx, "failed to get orders by user ID", slog.String("error", err.Error()))
		return nil, err
	}
	return orders, nil
}
