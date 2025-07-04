package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	userModel "github.com/umefy/go-web-app-template/internal/domain/user/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/openapi/v1/handler/user/mapping"
	loggerSrvMocks "github.com/umefy/go-web-app-template/mocks/domain/logger/service"
	userSrvMocks "github.com/umefy/go-web-app-template/mocks/domain/user/service"
	api "github.com/umefy/go-web-app-template/openapi/generated/go/openapi"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
)

type GetUsersSuite struct {
	suite.Suite
}

func (s *GetUsersSuite) TestGetUsers() {
	userService := userSrvMocks.NewMockService(s.T())
	loggerService := loggerSrvMocks.NewMockService(s.T())

	users := []*userModel.User{
		{
			ID:    1,
			Email: "john.doe@example.com",
			Age:   20,
		},
	}

	userService.EXPECT().GetUsers(context.Background()).Return(users, nil)
	loggerService.EXPECT().DebugContext(context.Background(), mock.Anything).Return()

	h := NewHandler(userService, loggerService)

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
