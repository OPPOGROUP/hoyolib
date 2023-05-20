package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"net/http"
)

type StarRail struct {
	client
}

func NewStarRailClient(oversea bool) (Client, error) {
	c := &StarRail{client{
		Api:         "",
		ActId:       "",
		SignInfoUrl: "",
		SignUrl:     "",
	}}
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

func (c *StarRail) Init() {

}
