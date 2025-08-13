package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	loggerMocks "github.com/umefy/go-web-app-template/mocks/infrastructure/logger"
	userSrvMocks "github.com/umefy/go-web-app-template/mocks/service/user"
	"github.com/umefy/go-web-app-template/pkg/pagination"
)

type GetUsersSuite struct {
	suite.Suite
}

func (s *GetUsersSuite) TestGetUsers() {
	userService := userSrvMocks.NewMockService(s.T())
	logger := loggerMocks.NewMockLogger(s.T())

	users := []*userDomain.User{
		{
			ID:    1,
			Email: "john.doe@example.com",
			Age:   20,
		},
	}

	userService.EXPECT().GetUsers(context.Background(), pagination.NewFromQueryParams("0", "25", "false")).Return(users, nil, nil)
	logger.EXPECT().DebugContext(context.Background(), "GetUsers")

	h := NewHandler(userService, logger, nil)

	req := httptest.NewRequest(http.MethodGet, "/openapi/v1/users", nil)
	rec := httptest.NewRecorder()

	err := h.GetUsers(rec, req)
	s.NoError(err)

	s.Equal(http.StatusOK, rec.Code)
}

func TestGetUsersSuite(t *testing.T) {
	suite.Run(t, new(GetUsersSuite))
}
