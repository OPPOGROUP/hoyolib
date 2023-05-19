package errors

import (
	"errors"
	"fmt"
)

var (
	ErrRequestParams = errors.New("not enough params or error params in building http request")
	ErrBuildRequest  = errors.New("build request fail")
	ErrSendRequest   = errors.New("error detected in sending http request")
	ErrHttpMethod    = errors.New("wrong http method")
	ErrHttpCode      = errors.New("wrong http code")
	ErrJsonDecode    = errors.New("json decode error")
)

func NewInternalError(code int, msg string) error {
	return errors.New(fmt.Sprintf("code = %d, msg = %s", code, msg))
}
