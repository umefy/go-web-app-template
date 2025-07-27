package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	api "github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/generated"
	"github.com/umefy/go-web-app-template/internal/delivery/restful/openapi/v1/user/mapping"
	userDomain "github.com/umefy/go-web-app-template/internal/domain/user"
	loggerMocks "github.com/umefy/go-web-app-template/mocks/infrastructure/logger"
	userSrvMocks "github.com/umefy/go-web-app-template/mocks/service/user"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
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

	userService.EXPECT().GetUsers(context.Background()).Return(users, nil)
	logger.EXPECT().DebugContext(context.Background(), "GetUsers")

	h := NewHandler(userService, logger)

	req := httptest.NewRequest(http.MethodGet, "/openapi/v1/users", nil)
	rec := httptest.NewRecorder()

	err := h.GetUsers(rec, req)
	s.NoError(err)

	s.Equal(http.StatusOK, rec.Code)
	expectedJSON, err := jsonkit.Marshal(&api.UserGetAllResponse{
		Data: sliceskit.Map(users, mapping.UserModelToApiUser),
	})

	s.NoError(err)
	s.JSONEq(string(expectedJSON), rec.Body.String())
}

func TestGetUsersSuite(t *testing.T) {
	suite.Run(t, new(GetUsersSuite))
}
