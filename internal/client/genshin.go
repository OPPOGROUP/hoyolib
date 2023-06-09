package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"net/http"
)

type GenshinClient struct {
	client
}

func NewGenshinClient(server hoyolib_pb.RegisterRequest_AccountType, accountId, cookieToken string) (Client, error) {
	c := &GenshinClient{client{
		server: server,
		game:   hoyolib_pb.GameType_Genshin,
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
		mark := "sol"
		actId = "e202102251931481"
		accountInfoUrl, signUrl = c.generateApi(api, mark)
		oversea = true
		cookies = map[string]string{
			"account_id":   accountId,
			"cookie_token": cookieToken,
			"ltoken":       cookieToken,
			"ltuid":        accountId,
		}
		gameBiz = "hk4e_global"
	} else { // cn api
		//accountApi = "https://bbs-api.mihoyo.com"
		api = "https://api-takumi.mihoyo.com"
		mark := "bbs_sign_reward"
		actId = "e202009291139501"
		accountInfoUrl, signUrl = c.generateApi(api, mark)
		oversea = false
		cookies = map[string]string{
			"account_id_v2":   accountId,
			"cookie_token_v2": cookieToken,
			"ltuid_v2":        accountId,
			"ltoken_v2":       cookieToken,
		}
		gameBiz = "hk4e_cn"
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
			"uid":    c.userInfo.GameUid,
			"region": c.userInfo.Region,
		}),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
