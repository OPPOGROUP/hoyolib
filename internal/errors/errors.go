package errors

import "errors"

var (
	ErrRequestParams = errors.New("not enough params or error params in building http request")
	ErrBuildRequest  = errors.New("build request fail")
	ErrSendRequest   = errors.New("error detected in sending http request")
)
