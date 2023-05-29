package user

import (
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

type Info struct {
	AccountId   string                                                                           `json:"account_id"`
	CookieToken string                                                                           `json:"cookie_token"`
	Clients     map[hoyolib_pb.RegisterRequest_AccountType]map[hoyolib_pb.GameType]client.Client `json:"clients"`
}
