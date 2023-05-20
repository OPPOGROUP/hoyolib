package errors

import (
	"errors"
	"fmt"
)

var (
	ErrRequestParams      = errors.New("not enough params or error params in building http request")
	ErrBuildRequest       = errors.New("build request fail")
	ErrSendRequest        = errors.New("error detected in sending http request")
	ErrHttpMethod         = errors.New("wrong http method")
	ErrHttpCode           = errors.New("wrong http code")
	ErrJsonDecode         = errors.New("json decode error")
	ErrInvalidRequest     = errors.New("invalid request")
	ErrInvalidAccountId   = errors.New("invalid account id")
	ErrInvalidCookieToken = errors.New("invalid cookie token")
	ErrInvalidAccountType = errors.New("invalid account type")
	ErrEmptyGames         = errors.New("empty games")
	ErrInvalidGameType    = errors.New("invalid game type")
)

func NewInternalError(code int, msg string) error {
	return errors.New(fmt.Sprintf("code = %d, msg = %s", code, msg))
}
