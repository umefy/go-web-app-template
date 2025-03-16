package error

import (
	"fmt"
	"net/http"

	bizError "github.com/umefy/go-web-app-template/app/error"
)

const (
	serviceName = "userService"
)

var (
	UserNotFound = bizError.NewError(fmt.Sprintf("%s_1001", serviceName), "user not found", http.StatusNotFound)
)
