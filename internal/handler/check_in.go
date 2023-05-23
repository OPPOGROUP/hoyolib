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
		for gid, c := range u.ClientsCN {
			err := c.CheckIn()
			if err != nil {
				resp.CheckInInfoCN[gid] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: false,
					Msg:     err.Error(),
				}
				log.Error().Err(err).Str("server", "CN").Int32("game_id", gid).Msg("Check-in failed")
			} else {
				resp.CheckInInfoCN[gid] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: true,
				}
			}
		}
		for gid, c := range u.ClientsOversea {
			err := c.CheckIn()
			if err != nil {
				resp.CheckInInfoOversea[gid] = &hoyolib_pb.CheckInResponse_CheckInStatus{
					Success: false,
					Msg:     err.Error(),
				}
				log.Error().Err(err).Str("server", "Oversea").Int32("game_id", gid).Msg("Check-in failed")
			} else {
				resp.CheckInInfoOversea[gid] = &hoyolib_pb.CheckInResponse_CheckInStatus{
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

func verifyCheckInRequest(req *hoyolib_pb.CheckInRequest) error {
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
