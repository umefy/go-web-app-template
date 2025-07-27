package error

import (
	"fmt"
	"net/http"

	appError "github.com/umefy/go-web-app-template/internal/domain/error"
)

const (
	serviceName = "orderService"
)

var (
	OrderNotFound = appError.NewError(fmt.Sprintf("%s_1001", serviceName), "order not found", http.StatusNotFound)
)
