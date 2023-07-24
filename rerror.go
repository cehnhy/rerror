// rerror wrap error with stack and http response info
package rerror

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var ErrNil = stderrors.New("rerror: err nil")

type ResponseError struct {
	err error // primitive error or primitive error with stack, default: errNil, non-nil

	httpStatus int // http status code

	Code    string `json:"code,omitempty"`    // application specific error code
	Message string `json:"message,omitempty"` // error message and how to solve this problem
}

func New(httpStatus int, format string, arg ...any) *ResponseError {
	return &ResponseError{
		err:        ErrNil,
		httpStatus: httpStatus,
		Code:       "",
		Message:    fmt.Sprintf(format, arg...),
	}
}

func (e *ResponseError) E(err error) *ResponseError {
	if err == nil {
		return e
	}

	if _, ok := err.(*ResponseError); ok {
		panic("err is already a responseError")
	}

	if _, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		e.err = err
		return e
	}

	e.err = errors.WithStack(err)
	return e
}

func (e *ResponseError) C(code string) *ResponseError {
	e.Code = code
	return e
}

func (e *ResponseError) Unwrap() error {
	return e.err
}

func (e *ResponseError) Error() string {
	return e.err.Error()
}

func (e *ResponseError) Stack() string {
	if _, ok := e.err.(interface{ StackTrace() errors.StackTrace }); ok {
		return fmt.Sprintf("%+v", e.err)
	}

	return ""
}

func (e *ResponseError) HTTPStatus() int {
	return e.httpStatus
}
