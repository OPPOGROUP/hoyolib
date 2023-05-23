package client

import (
	"fmt"
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/utils"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"net/http"
)

type GenshinClient struct {
	client
}

func NewGenshinClient(oversea bool, accountId, cookieToken string) (Client, error) {
	c := &GenshinClient{client{}}
	var err error
	var (
		accountApi     string
		api            string
		actId          string
		signInfoUrl    string
		signUrl        string
		accountInfoUrl string
	)
	if oversea {
		accountApi = "https://bbs-api-os.hoyolab.com"
		api = "https://sg-hk4e-api.hoyolab.com"
		mark := "sol"
		actId = "e202102251931481"
		accountInfoUrl = fmt.Sprintf("%s/game_record/card/api/getGameRecordCard", accountApi)
		signUrl, signInfoUrl = utils.GetSignUrl(api, mark)
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
