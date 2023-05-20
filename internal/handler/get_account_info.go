package handler

import (
	"context"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (HoyolibServer) GetAccountInfo(context.Context, *hoyolib_pb.AccountInfoRequest) (*hoyolib_pb.AccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
