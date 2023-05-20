package handler

import (
	"context"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

func (HoyolibServer) Sign(context.Context, *hoyolib_pb.SignRequest) (*hoyolib_pb.SignResponse, error) {
	resp := &hoyolib_pb.SignResponse{}
	return resp, nil
}
