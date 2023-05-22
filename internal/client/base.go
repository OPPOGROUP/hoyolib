package client

import (
	"encoding/json"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/user"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"io"
	"time"
)

type Client interface {
	CheckIn() error
}

type client struct {
	userInfo           *user.Info
	accountInfoRequest *request.Request
	signInfoRequest    *request.Request
	signRequest        *request.Request
	cancel             chan struct{}
}

func (c *client) Loop() {
	cancel := make(chan struct{}, 1)
	c.updateAccountInfo()

	tick := time.NewTicker(1 * time.Hour)
	go func(cancel chan struct{}) {
		defer tick.Stop()
		for range tick.C {
			select {
			case <-cancel:
				return
			default:
				c.updateAccountInfo()
			}
		}
	}(cancel)
	c.cancel = cancel
}

func (c *client) StopLoop() {
	defer close(c.cancel)
	c.cancel <- struct{}{}
}

func (c *client) CheckIn() error {
	resp, err := c.signRequest.Do()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.ErrHttpCode
	}
	body, _ := io.ReadAll(resp.Body)
	r := new(SignResponse)
	err = json.Unmarshal(body, r)
	if err != nil {
		return errors.ErrJsonDecode
	}
	switch r.Retcode {
	case 0:
	case -5003:
	default:
		return errors.NewInternalError(r.Retcode, r.Message)
	}
	return nil
}

func (c *client) updateAccountInfo() {
	// TODO: update account info
}

func (c *client) updateSignInfo(isSign bool) {
	c.userInfo.SetSign(isSign)
}
