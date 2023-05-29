package user

import (
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

type Info struct {
	AccountId   string                                                                           `json:"account_id"`
	CookieToken string                                                                           `json:"cookie_token"`
	ClientNotes map[hoyolib_pb.RegisterRequest_AccountType]map[hoyolib_pb.GameType]struct{}      `json:"client_note"`
	Clients     map[hoyolib_pb.RegisterRequest_AccountType]map[hoyolib_pb.GameType]client.Client `json:"-"`
}
