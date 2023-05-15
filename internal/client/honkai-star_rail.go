package client

import (
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/utils/request"
	"net/http"
)

type StarRail struct {
	Client
}

func NewStarRailClient() (IClient, error) {
	c := &StarRail{Client{
		Api:         "",
		ActId:       "",
		SignInfoUrl: "",
		SignUrl:     "",
	}}
	var err error
	c.signRequest, err = request.NewRequest(request.WithMethod(http.MethodPost),
		request.WithHeaders(cte.HoyolabHeaders),
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}
