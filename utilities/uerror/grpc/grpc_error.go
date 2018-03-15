package grpc

import (
	"errors"
	"utilities/uerror"

	"google.golang.org/grpc/status"
)

const (
	errorTypeBadParam          = "bad_param"
	errorTypeInternalError     = "internal_error"
	errorTypeAuthorizationFail = "authorization_fail"
	errorTypePermissionDenied  = "permission_denied"
)

type customError struct {
	ErrType string
	Message string
	Trace   string
}

func GetStatusFromError(err error) (retry bool, status int, retError error) {
	if err == nil {
		return false, 200, nil
	}
	if ok, retError := IsNotFound(err); ok {
		return false, 404, retError
	}
	if ok, retError := IsBadRequest(err); ok {
		return false, 400, retError
	}
	return true, 500, uerror.StackTrace(err)
}

func NotFound(errString string) error {
	return status.Error(404, errString)
}

func BadRequest(errString string) error {
	return status.Error(400, errString)
}

func IsNotFound(err error) (bool, error) {
	stt, ok := status.FromError(err)
	if ok {
		code := stt.Code()
		desc := stt.Message()
		if code == 404 {
			return true, errors.New(desc)
		}
	}
	return false, err
}

func IsBadRequest(err error) (bool, error) {
	stt, ok := status.FromError(err)
	if ok {
		code := stt.Code()
		desc := stt.Message()
		if code == 400 {
			return true, errors.New(desc)
		}
	}
	return false, err
}

func (m customError) Error() string {
	return m.Message
}

func StackTrace(err error) error {
	return customError{
		ErrType: errorTypeInternalError,
		Message: err.Error(),
		Trace:   uerror.StackTraceStr(err),
	}
}

func ErrAuthorizationFail(err error) error {
	return customError{
		ErrType: errorTypeAuthorizationFail,
		Message: err.Error(),
	}
}

func ErrPermissionDenied(err error) error {
	return customError{
		ErrType: errorTypePermissionDenied,
		Message: err.Error(),
	}
}

func ErrBadParam(err error) error {
	return customError{
		ErrType: errorTypeBadParam,
		Message: err.Error(),
	}
}
