package client

import (
	"encoding/json"
	"github.com/OPPOGROUP/hoyolib/internal/cte"
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
	userInfo           *gameInfo
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
	gameType := cte.LocalGameTypeToMihoyoGameType[c.game]
	for _, data := range r.Data.List {
		data := data
		if data.GameId == gameType {
			userInfo := &gameInfo{
				GameRoleId: data.GameRoleId,
				Region:     data.Region,
				Nickname:   data.Nickname,
				Level:      data.Level,
				Data:       data.Data,
			}
			c.userInfo = userInfo
			break
		}
	}
	return nil
}
