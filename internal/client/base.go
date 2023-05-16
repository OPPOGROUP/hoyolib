package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/user"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"time"
)

type IClient interface {
	Sign() error
}

type Client struct {
	Api                string
	ActId              string
	SignInfoUrl        string
	SignUrl            string
	userInfo           *user.Info
	accountInfoRequest *request.Request
	signInfoRequest    *request.Request
	signRequest        *request.Request
	cancel             chan struct{}
}

func (c *Client) Loop() {
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

func (c *Client) StopLoop() {
	defer close(c.cancel)
	c.cancel <- struct{}{}
}

func (c *Client) Sign() error {
	return nil
}

func (c *Client) updateAccountInfo() {
	// TODO: update account info
}
