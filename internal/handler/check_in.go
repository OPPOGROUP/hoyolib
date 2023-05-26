package handler

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

func (HoyolibServer) CheckIn(_ context.Context, req *hoyolib_pb.CheckInRequest) (*hoyolib_pb.CheckInResponse, error) {
	resp := &hoyolib_pb.CheckInResponse{
		Header: &hoyolib_pb.ResponseHeader{
			Code:   int32(hoyolib_pb.ErrorCode_OK),
			UserId: req.GetUserId(),
		},
		CheckInInfoCN:      make(map[int32]*hoyolib_pb.CheckInResponse_CheckInStatus),
		CheckInInfoOversea: make(map[int32]*hoyolib_pb.CheckInResponse_CheckInStatus),
	}
	log.Debug().Msgf("Check-in request: %+v", req)
	defer log.Debug().Msgf("Check-in response: %+v", resp)

	if err := verifyCheckInRequest(req); err != nil {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_INVALID_REQUEST_PARAM),
			Message: err.Error(),
		}
		log.Error().Err(err).Msg("Check-in request verification failed")
		return resp, nil
	}
	if u, ok := m[req.GetUserId()]; !ok {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_ERROR_USER_NOT_REGISTER),
			Message: errors.ErrUserNotRegistered.Error(),
		}
		log.Error().Err(errors.ErrUserNotRegistered).Msg("Check-in request verification failed")
		return resp, nil
	} else {
		for gid, c := range u.Clients[hoyolib_pb.RegisterRequest_CN] {
			err := c.CheckIn()
			if err != nil {
				resp.CheckInInfoCN[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: false,
					Msg:     err.Error(),
				}
				log.Error().Err(err).Str("server", "CN").Int32("game_id", int32(gid)).Msg("Check-in failed")
			} else {
				resp.CheckInInfoCN[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: true,
				}
			}
		}
		for gid, c := range u.Clients[hoyolib_pb.RegisterRequest_OVERSEA] {
			err := c.CheckIn()
			if err != nil {
				resp.CheckInInfoOversea[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: false,
					Msg:     err.Error(),
				}
				log.Error().Err(err).Str("server", "Oversea").Int32("game_id", int32(gid)).Msg("Check-in failed")
			} else {
				resp.CheckInInfoOversea[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: true,
				}
			}
		}
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:   int32(hoyolib_pb.ErrorCode_OK),
			UserId: req.GetUserId(),
		}
	}
	return resp, nil
}

func checkInUser(userid int32) error {
	return nil
}

func verifyCheckInRequest(req *hoyolib_pb.CheckInRequest) error {
	if req == nil {
		return errors.ErrInvalidRequest
	}
	if req.GetUserId() == 0 {
		return errors.ErrInvalidUserId
	}
	return nil
}
