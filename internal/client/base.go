package client

import "github.com/OPPOGROUP/hoyolib/internal/utils/request"

type IClient interface {
	Sign() error
}

type Client struct {
	Api                string
	ActId              string
	SignInfoUrl        string
	SignUrl            string
	accountInfoRequest *request.Request
	signRequest        *request.Request
}

func (c *Client) Sign() error {
	return nil
}
