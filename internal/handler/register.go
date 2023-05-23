package handler

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

type uInfo struct {
	AccountId      string
	CookieToken    string
	ClientsCN      map[int32]client.Client
	ClientsOversea map[int32]client.Client
}

var (
	m         = make(map[int64]*uInfo)
	uid int64 = 100000
)

func (HoyolibServer) Register(_ context.Context, req *hoyolib_pb.RegisterRequest) (*hoyolib_pb.RegisterResponse, error) {
	resp := &hoyolib_pb.RegisterResponse{
		Header: &hoyolib_pb.ResponseHeader{
			Code:   int32(hoyolib_pb.ErrorCode_OK),
			UserId: req.GetUserId(),
		},
	}
	log.Debug().Msgf("Register request: %+v", req)
	defer log.Debug().Msgf("Register response: %+v", resp)

	if err := verifyRegisterRequest(req); err != nil {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_INVALID_REQUEST_PARAM),
			Message: err.Error(),
		}
		log.Error().Err(err).Msg("Register request verification failed")
		return resp, nil
	}
	var oversea = false
	if req.GetAccountType() == hoyolib_pb.RegisterRequest_OVERSEA {
		oversea = true
	}
	_uid, err := createUser(req, oversea)
	if err != nil {
		resp.Header = &hoyolib_pb.ResponseHeader{
			Code:    int32(hoyolib_pb.ErrorCode_ERROR_CREATE_USER),
			Message: err.Error(),
		}
		log.Error().Err(err).Msg("Create user failed")
		return resp, nil
	}

	resp.Header = &hoyolib_pb.ResponseHeader{
		Code:   int32(hoyolib_pb.ErrorCode_OK),
		UserId: _uid,
	}
	return resp, nil
}

func createUser(req *hoyolib_pb.RegisterRequest, oversea bool) (int64, error) {
	u := req.GetUserId()
	if u == 0 {
		u = uid
		uid++
	}
	info := &uInfo{
		AccountId:      req.AccountId,
		CookieToken:    req.CookieToken,
		ClientsCN:      make(map[int32]client.Client),
		ClientsOversea: make(map[int32]client.Client),
	}
	for _, g := range req.GetGames() {
		var (
			c   client.Client
			err error
		)
		switch g {
		case hoyolib_pb.GameType_Genshin:
			c, err = client.NewGenshinClient(oversea, req.AccountId, req.CookieToken)
			if err != nil {
				return 0, err
			}
		case hoyolib_pb.GameType_StarRail:
			c, err = client.NewStarRailClient(oversea, req.AccountId, req.CookieToken)
			if err != nil {
				return 0, err
			}
		}
		if oversea {
			info.ClientsOversea[int32(g)] = c
		} else {
			info.ClientsCN[int32(g)] = c
		}
	}
	m[u] = info
	return u, nil
}

func verifyRegisterRequest(req *hoyolib_pb.RegisterRequest) error {
	if req == nil {
		return errors.ErrInvalidRequest
	}
	if req.AccountId == "" {
		return errors.ErrInvalidAccountId
	}
	if req.CookieToken == "" {
		return errors.ErrInvalidCookieToken
	}
	if hoyolib_pb.RegisterRequest_AccountType_name[int32(req.AccountType)] == "" {
		return errors.ErrInvalidAccountType
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
