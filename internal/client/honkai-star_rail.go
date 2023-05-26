package client

import (
	"fmt"
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/utils"
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
		accountApi     string
		api            string
		actId          string
		signInfoUrl    string
		signUrl        string
		accountInfoUrl string
		oversea        bool
	)
	if server == hoyolib_pb.RegisterRequest_OVERSEA {
		accountApi = "https://bbs-api-os.hoyolab.com"
		api = "https://sg-public-api.hoyolab.com"
		mark := "luna"
		actId = "e202303301540311"
		accountInfoUrl = fmt.Sprintf("%s/game_record/card/api/getGameRecordCard", accountApi)
		signUrl, signInfoUrl = utils.GetSignUrl(api, mark)
		oversea = true
	} else {
		// todo: add mainland china api
		return nil, errors.ErrNotImplemented
	}
	c.accountInfoRequest, err = request.NewRequest(
		request.WithMethod(http.MethodGet),
		request.WithUrl(accountInfoUrl),
		request.WithHeaders(cte.GetHeaders(oversea)),
		request.WithCookies(map[string]string{
			"account_id":   accountId,
			"cookie_token": cookieToken,
		}),
		request.WithParams(map[string]string{
			"uid": accountId,
		}),
	)
	if err != nil {
		return nil, err
	}
	c.signInfoRequest, err = request.NewRequest(
		request.WithMethod(http.MethodGet),
		request.WithUrl(signInfoUrl),
		request.WithHeaders(cte.GetHeaders(oversea)),
		request.WithCookies(map[string]string{
			"account_id":   accountId,
			"cookie_token": cookieToken,
		}),
		request.WithParams(map[string]string{
			"act_id": actId,
			"lang":   "zh-cn",
		}),
	)
	if err != nil {
		return nil, err
	}
	c.signRequest, err = request.NewRequest(
		request.WithMethod(http.MethodPost),
		request.WithUrl(signUrl),
		request.WithHeaders(cte.GetHeaders(oversea)),
		request.WithCookies(map[string]string{
			"account_id":   accountId,
			"cookie_token": cookieToken,
		}),
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
