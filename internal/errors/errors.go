package errors

import (
	"errors"
	"fmt"
)

var (
	ErrBuildRequest         = errors.New("build request fail")
	ErrSendRequest          = errors.New("error detected in sending http request")
	ErrHttpMethod           = errors.New("wrong http method")
	ErrHttpCode             = errors.New("wrong http code")
	ErrJsonDecode           = errors.New("json decode error")
	ErrInvalidRequest       = errors.New("invalid request")
	ErrInvalidAccountId     = errors.New("invalid account id")
	ErrInvalidCookieToken   = errors.New("invalid cookie token")
	ErrInvalidAccountType   = errors.New("invalid account type")
	ErrEmptyGames           = errors.New("empty games")
	ErrInvalidGameType      = errors.New("invalid game type")
	ErrInvalidUserId        = errors.New("invalid user id")
	ErrUserNotRegistered    = errors.New("user not registered")
	ErrNotImplemented       = errors.New("not implemented")
	ErrGrpcClientInitFailed = errors.New("grpc client init failed")
)

func NewInternalError(code int, msg string) error {
	return errors.New(fmt.Sprintf("code = %d, msg = %s", code, msg))
}
