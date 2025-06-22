package user

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/internal/infrastructure/http/handler/user/mapping"
	loggerSrvMocks "github.com/umefy/go-web-app-template/mocks/domain/logger/service"
	userSrvMocks "github.com/umefy/go-web-app-template/mocks/domain/user/service"
	api "github.com/umefy/go-web-app-template/openapi/protogen/v1/models"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/sliceskit"
)

type GetUsersSuite struct {
	suite.Suite
}

func (s *GetUsersSuite) TestGetUsers() {
	userService := userSrvMocks.NewMockService(s.T())
	loggerService := loggerSrvMocks.NewMockService(s.T())

	users := []*model.User{
		{
			ID:   1,
			Name: null.ValueFrom("John Doe"),
			Age:  null.ValueFrom(20),
		},
	}

	userService.EXPECT().GetUsers(context.Background()).Return(users, nil)
	loggerService.EXPECT().DebugContext(context.Background(), mock.Anything).Return()

	h := NewHandler(userService, loggerService)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	err := h.GetUsers(rec, req)
	s.NoError(err)

	s.Equal(http.StatusOK, rec.Code)
	expectedJSON, err := jsonkit.MarshalProto(&api.UserGetAllResponse{
		Data: sliceskit.Map(users, mapping.UserModelToApiUser),
	})

	s.NoError(err)
	s.JSONEq(string(expectedJSON), rec.Body.String())
}

func TestGetUsersSuite(t *testing.T) {
	suite.Run(t, new(GetUsersSuite))
}
