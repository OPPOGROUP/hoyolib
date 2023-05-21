package handler

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

func (HoyolibServer) Sign(ctx context.Context, req *hoyolib_pb.SignRequest) (*hoyolib_pb.SignResponse, error) {
	resp := &hoyolib_pb.SignResponse{}
	log.Debug().Msgf("Sign request: %+v", req)
	defer log.Debug().Msgf("Sign response: %+v", resp)

	if err := verifySignRequest(req); err != nil {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_INVALID_REQUEST_PARAM),
			Message: err.Error(),
		}
		return resp, nil
	}
	if u, ok := m[req.GetUserId()]; !ok {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_ERROR_USER_NOT_REGISTER),
			Message: errors.ErrInvalidUserNotRegistered.Error(),
		}
		return resp, nil
	} else {

	}
	return resp, nil
}

func verifySignRequest(req *hoyolib_pb.SignRequest) error {
	if req == nil {
		return errors.ErrInvalidRequest
	}
	if req.GetUserId() == 0 {
		return errors.ErrInvalidUserId
	}
	if len(req.GetGames()) == 0 {
		return errors.ErrEmptyGames
	}
	for _, g := range req.GetGames() {
		if hoyolib_pb.GameType_name[int32(g)] == "" {
			return errors.ErrInvalidGameType
		}
	}
	return nil
}
