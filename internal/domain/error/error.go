package error

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code     string
	Message  string
	HTTPCode int
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %s, error msg: %s", e.Code, e.Message)
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"code":    e.Code,
		"message": e.Message,
	})
}

func NewError(errorCode string, errorMsg string, httpCode int) *Error {
	return &Error{
		Code:     errorCode,
		Message:  errorMsg,
		HTTPCode: httpCode,
	}
}
