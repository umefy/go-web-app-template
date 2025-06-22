package error

import (
	"fmt"
	"net/http"

	appError "github.com/umefy/go-web-app-template/internal/domain/error"
)

const (
	serviceName = "userService"
)

var (
	UserNotFound = appError.NewError(fmt.Sprintf("%s_1001", serviceName), "user not found", http.StatusNotFound)
)
