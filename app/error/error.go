package error

import (
	"fmt"
)

type Error struct {
	ErrorCode string
	ErrorMsg  string
	HTTPCode  int
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %s, error msg: %s", e.ErrorCode, e.ErrorMsg)
}

func NewError(errorCode string, errorMsg string, httpCode int) *Error {
	return &Error{
		ErrorCode: errorCode,
		ErrorMsg:  errorMsg,
		HTTPCode:  httpCode,
	}
}
