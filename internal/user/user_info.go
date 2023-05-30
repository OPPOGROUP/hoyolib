package user

import (
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

type Info struct {
	Uid   int64                                                  `json:"uid"`
	Infos map[hoyolib_pb.RegisterRequest_AccountType]*ServerInfo `json:"infos"`
}

type ServerInfo struct {
	Active      bool                                  `json:"-"`
	Msg         string                                `json:"-"`
	AccountId   string                                `json:"account_id"`
	CookieToken string                                `json:"cookie_token"`
	ClientNotes []hoyolib_pb.GameType                 `json:"client_notes"`
	Clients     map[hoyolib_pb.GameType]client.Client `json:"-"`
}

func (i *Info) CreateClients(server hoyolib_pb.RegisterRequest_AccountType) error {
	info := i.Infos[server]
	info.Clients = make(map[hoyolib_pb.GameType]client.Client)
	for _, g := range info.ClientNotes {
		var (
			c   client.Client
			err error
		)
		switch g {
		case hoyolib_pb.GameType_Genshin:
			c, err = client.NewGenshinClient(server, info.AccountId, info.CookieToken)
			if err != nil {
				info.Active = false
				info.Msg = err.Error()
				return err
			}
		case hoyolib_pb.GameType_StarRail:
			c, err = client.NewStarRailClient(server, info.AccountId, info.CookieToken)
			if err != nil {
				info.Active = false
				info.Msg = err.Error()
				return err
			}
		}
		info.Clients[g] = c
	}
	info.Active = true
	return nil
}
