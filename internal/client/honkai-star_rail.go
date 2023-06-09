package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"net/http"
)

type StarRail struct {
	client
}

func NewStarRailClient(server hoyolib_pb.RegisterRequest_AccountType, accountId, cookieToken string) (Client, error) {
	c := &StarRail{client{
		server: server,
		game:   hoyolib_pb.GameType_StarRail,
	}}
	var err error
	var (
		//accountApi     string
		api   string
		actId string
		//signInfoUrl    string
		signUrl        string
		accountInfoUrl string
		oversea        bool
		cookies        map[string]string
		gameBiz        string
	)
	if server == hoyolib_pb.RegisterRequest_OVERSEA {
		//accountApi = "https://bbs-api-os.hoyolab.com"
		api = "https://sg-public-api.hoyolab.com"
		mark := "luna"
		actId = "e202303301540311"
		accountInfoUrl, signUrl = c.generateApi(api, mark)
		oversea = true
		cookies = map[string]string{
			"account_id":   accountId,
			"cookie_token": cookieToken,
			"ltoken":       cookieToken,
			"ltuid":        accountId,
		}
		gameBiz = "hkrpg_global"
	} else {
		// todo: add mainland china api
		return nil, errors.ErrNotImplemented
	}
	c.accountInfoRequest, err = request.NewRequest(
		request.WithMethod(http.MethodGet),
		request.WithUrl(accountInfoUrl),
		request.WithHeaders(cte.GetHeaders(oversea)),
		request.WithCookies(cookies),
		request.WithParams(map[string]string{
			"game_biz": gameBiz,
		}),
	)
	if err != nil {
		return nil, err
	}
	err = c.updateAccountInfo()
	if err != nil {
		return nil, err
	}
	c.signRequest, err = request.NewRequest(
		request.WithMethod(http.MethodPost),
		request.WithUrl(signUrl),
		request.WithHeaders(cte.GetHeaders(oversea)),
		request.WithCookies(cookies),
		request.WithPayloads(map[string]interface{}{
			"act_id": actId,
			"lang":   "zh-cn",
		}),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
