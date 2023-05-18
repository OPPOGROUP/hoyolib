package grpc

import (
	"context"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (HoyolibService) Sign(context.Context, *hoyolib_pb.SignRequest) (*hoyolib_pb.SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
