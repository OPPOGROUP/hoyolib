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
	}
	log.Debug().Msgf("Check-in request: %+v", req)
	defer log.Debug().Msgf("Check-in response: %+v", resp)

	if err := verifyCheckInRequest(req); err != nil {
		log.Error().Err(err).Msg("Check-in request verification failed")
		return &hoyolib_pb.CheckInResponse{
			Header: &hoyolib_pb.ResponseHeader{
				Code:    int32(hoyolib_pb.ErrorCode_INVALID_REQUEST_PARAM),
				UserId:  req.GetUserId(),
				Message: err.Error(),
			},
		}, nil
	}
	return CheckInUser(req.GetUserId())
}

func CheckInUser(userid int64) (*hoyolib_pb.CheckInResponse, error) {
	resp := &hoyolib_pb.CheckInResponse{
		Header: &hoyolib_pb.ResponseHeader{
			Code:   int32(hoyolib_pb.ErrorCode_OK),
			UserId: userid,
		},
		ClientsInfo: make(map[int32]*hoyolib_pb.CheckInResponse_CheckInInfo),
	}
	if u, ok := m[userid]; !ok {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_ERROR_USER_NOT_REGISTER),
			Message: errors.ErrUserNotRegistered.Error(),
		}
		log.Error().Err(errors.ErrUserNotRegistered).Msg("Check-in request verification failed")
		return resp, errors.ErrUserNotRegistered
	} else {
		for accountType, clients := range (*u).Clients {
			clients := clients
			info := make(map[int32]*hoyolib_pb.CheckInResponse_CheckInStatus)
			for gid, c := range clients {
				err := c.CheckIn()
				if err != nil {
					info[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
						Success: false,
						Msg:     err.Error(),
					}
					log.Error().Err(err).Int32("server", int32(accountType)).Int32("game_id", int32(gid)).Msg("Check-in failed")
				} else {
					info[int32(gid)] = &hoyolib_pb.CheckInResponse_CheckInStatus{
						Success: true,
					}
				}
			}
			resp.ClientsInfo[int32(accountType)] = &hoyolib_pb.CheckInResponse_CheckInInfo{
				Info: info,
			}
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
	return nil
}
