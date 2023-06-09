package client

import (
	"encoding/json"
	"fmt"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"io"
)

type Client interface {
	CheckIn() error
}

type client struct {
	userInfo           *accountInfo
	server             hoyolib_pb.RegisterRequest_AccountType
	game               hoyolib_pb.GameType
	accountInfoRequest *request.Request
	signInfoRequest    *request.Request
	signRequest        *request.Request
}

func (c *client) CheckIn() error {
	resp, err := c.signRequest.Do()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Error().Msgf("http code = %d, body = %s", resp.StatusCode, string(body))
		return errors.ErrHttpCode
	}
	r := new(SignResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.ErrJsonDecode
	}
	if r.Retcode != 0 {
		return errors.NewInternalError(r.Retcode, r.Message)
	}
	return nil
}

func (c *client) updateAccountInfo() error {
	resp, err := c.accountInfoRequest.Do()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Error().Msgf("http code = %d, body = %s", resp.StatusCode, string(body))
		return errors.ErrHttpCode
	}
	r := new(AccountInfoResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.ErrJsonDecode
	}
	if r.Retcode != 0 {
		return errors.NewInternalError(r.Retcode, r.Message)
	}
	c.userInfo = &r.Data.List[0]

	return nil
}

func (c *client) generateApi(api, mark string) (accountInfoUrl string, signUrl string) {
	accountInfoUrl = fmt.Sprintf("%s/binding/api/getUserGameRolesByCookie", api)
	signUrl = fmt.Sprintf("%s/event/%s/sign", api, mark)
	return
}
