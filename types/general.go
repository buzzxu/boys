package types

import (
	"net/http"
)

type (
	// Error 异常结构
	Error struct {
		Code     int    `json:"code"`
		Key      string `json:"error,omitempty"`
		Success  bool   `json:"success"`
		Message  string `json:"message"`
		Internal error  `json:"-"`
	}
	// Result 返回结果
	Result struct {
		Code    int         `json:"code"`
		Success bool        `json:"success,omitempty"`
		Message interface{} `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
)

var (
	ErrNotFound           = NewError(http.StatusNotFound)
	ErrBadRequest         = NewError(http.StatusBadRequest)
	ErrServiceUnavailable = NewError(http.StatusServiceUnavailable)
)

// Result 构造Result
func ResultOf(code int, data interface{}) *Result {
	return &Result{
		Code:    code,
		Success: true,
		Data:    data,
	}
}

//Result 无Data
func ResultNilData(code int) *Result {
	return &Result{
		Code:    code,
		Success: true,
	}
}

func (err Error) Error() string {
	return err.Message
}
func ErrorOf(err error) *Error {
	return &Error{Code: http.StatusInternalServerError, Success: false, Message: err.Error(), Internal: err}
}
func NewError(code int, message ...string) *Error {
	he := &Error{Code: code, Success: false}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}
func NewHttpError(code int, message ...string) *Error {
	he := &Error{Code: code, Success: false, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}
func IsError(err error) bool {
	_, ok := err.(Error)
	return ok
}
