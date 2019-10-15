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
		Success bool        `json:"success"`
		Message interface{} `json:"message,omitempty"`
		Data    interface{} `json:"file,omitempty"`
	}
)

var (
	ErrNotFound           = NewError(http.StatusNotFound)
	ErrBadRequest         = NewError(http.StatusBadRequest)
	ErrServiceUnavailable = NewError(http.StatusServiceUnavailable)
)

func (err *Error) Error() string {
	return err.Message
}

func NewResult(code int, data interface{}) *Result {
	return ResultOf(code, data)
}

// ResultOf 构造Result
func ResultOf(code int, data ...interface{}) *Result {
	return &Result{
		Code:    code,
		Success: true,
		Data:    data,
	}
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
