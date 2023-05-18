package grpc

import (
	"context"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (HoyolibService) Register(context.Context, *hoyolib_pb.RegisterRequest) (*hoyolib_pb.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
