package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"net/http"
)

type GenshinClient struct {
	client
}

func NewGenshinClient(oversea bool) (Client, error) {
	c := &GenshinClient{client{}}
	var err error
	c.signRequest, err = request.NewRequest(
		request.WithMethod(http.MethodPost),
		request.WithHeaders(cte.GetHeaders(oversea)),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *GenshinClient) Init() {

}
