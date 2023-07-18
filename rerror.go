package rerror

import (
	stderrors "errors"
	"fmt"

	"github.com/pkg/errors"
)

var errNil = stderrors.New("nil error")

type responseError struct {
	err error // primitive error or primitive error with stack

	httpStatus int // http status code

	Code    string `json:"code,omitempty"`    // application specific error code
	Message string `json:"message,omitempty"` // error message and how to solve this problem
}

func New(httpStatus int, format string, arg ...any) *responseError {
	return &responseError{
		err:        errNil,
		httpStatus: httpStatus,
		Code:       "",
		Message:    fmt.Sprintf(format, arg...),
	}
}

func (e *responseError) E(err error) *responseError {
	if err == nil {
		return e
	}

	if _, ok := err.(*responseError); ok {
		panic("err is already a responseError")
	}

	if _, ok := err.(interface{ StackTrace() errors.StackTrace }); ok {
		e.err = err
		return e
	}

	e.err = errors.WithStack(err)
	return e
}

func (e *responseError) C(code string) *responseError {
	e.Code = code
	return e
}

func (e *responseError) Unwrap() error {
	return e.err
}

func (e *responseError) Error() string {
	return e.err.Error()
}

func (e *responseError) Stack() string {
	if e.err == nil {
		return ""
	}

	if _, ok := e.err.(interface{ StackTrace() errors.StackTrace }); ok {
		return fmt.Sprintf("%+v", e.err)
	}

	return ""
}

func (e *responseError) HTTPStatus() int {
	return e.httpStatus
}
