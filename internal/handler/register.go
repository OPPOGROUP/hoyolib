package handler

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/hoyolib/internal/user"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"github.com/spf13/viper"
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
	_uid, err := createUser(req, req.GetAccountType())
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

func createUser(req *hoyolib_pb.RegisterRequest, server hoyolib_pb.RegisterRequest_AccountType) (int64, error) {
	u := req.GetUserId()
	if u == 0 {
		u = uid
	}
	info := &user.Info{
		AccountId:   req.AccountId,
		CookieToken: req.CookieToken,
		ClientNotes: make(map[hoyolib_pb.RegisterRequest_AccountType]map[hoyolib_pb.GameType]struct{}),
		Clients:     make(map[hoyolib_pb.RegisterRequest_AccountType]map[hoyolib_pb.GameType]client.Client),
	}
	for _, g := range req.GetGames() {
		var (
			c   client.Client
			err error
		)
		switch g {
		case hoyolib_pb.GameType_Genshin:
			c, err = client.NewGenshinClient(server, req.AccountId, req.CookieToken)
			if err != nil {
				return 0, err
			}
		case hoyolib_pb.GameType_StarRail:
			c, err = client.NewStarRailClient(server, req.AccountId, req.CookieToken)
			if err != nil {
				return 0, err
			}
		}
		if info.Clients[server] == nil {
			info.Clients[server] = make(map[hoyolib_pb.GameType]client.Client)
		}
		info.Clients[server][g] = c
		if info.ClientNotes[server] == nil {
			info.ClientNotes[server] = make(map[hoyolib_pb.GameType]struct{})
		}
		info.ClientNotes[server][g] = struct{}{}
	}
	m[u] = info
	uid++
	if viper.GetBool("data.enable") {
		go func() {
			err := saveUser()
			if err != nil {
				log.Error().Err(err).Msg("Save user failed")
				return
			}
			log.Info().Msgf("Save user %d success", u)
		}()
	}
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
	if hoyolib_pb.RegisterRequest_AccountType_name[int32(req.AccountType)] == "" || req.AccountType == hoyolib_pb.RegisterRequest_UNKNOWN {
		return errors.ErrInvalidAccountType
	}
	if len(req.GetGames()) == 0 {
		return errors.ErrEmptyGames
	}
	for _, g := range req.GetGames() {
		if hoyolib_pb.GameType_name[int32(g)] == "" || g == hoyolib_pb.GameType_UNKNOWN_GAME {
			return errors.ErrInvalidGameType
		}
	}
	return nil
}
