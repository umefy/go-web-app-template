package user

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guregu/null/v6"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/umefy/go-web-app-template/gorm/generated/model"
	"github.com/umefy/go-web-app-template/internal/server/httpserver/user/mapping"
	loggerSrvMocks "github.com/umefy/go-web-app-template/mocks/app/logger/service"
	userSrvMocks "github.com/umefy/go-web-app-template/mocks/app/user/service"
	"github.com/umefy/godash/sliceskit"
)

type GetUsersSuite struct {
	suite.Suite
}

func (s *GetUsersSuite) TestGetUsers() {
	userService := userSrvMocks.NewService(s.T())
	loggerService := loggerSrvMocks.NewService(s.T())

	users := []*model.User{
		{
			ID:   1,
			Name: null.StringFrom("John Doe"),
			Age:  null.IntFrom(20),
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
	expectedJSON, err := json.Marshal(sliceskit.Map(users, mapping.UserModelToApiUser))
	s.NoError(err)
	s.JSONEq(string(expectedJSON), rec.Body.String())
}

func TestGetUsersSuite(t *testing.T) {
	suite.Run(t, new(GetUsersSuite))
}
