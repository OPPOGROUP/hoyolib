package handler

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/client"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

type uInfo struct {
	AccountId   string
	CookieToken string
	Clients     []client.Client
}

var (
	m         = make(map[int64]*uInfo)
	uid int64 = 100000
)

func (HoyolibServer) Register(ctx context.Context, req *hoyolib_pb.RegisterRequest) (*hoyolib_pb.RegisterResponse, error) {
	if err := verifyRegisterRequest(req); err != nil {
		return nil, err
	}
	resp := &hoyolib_pb.RegisterResponse{}
	return resp, nil
}

func createUser(req *hoyolib_pb.RegisterRequest) {
	u := req.GetUserId()
	if u == 0 {
		u = uid
		uid++
	}
	info := &uInfo{
		AccountId:   req.AccountId,
		CookieToken: req.CookieToken,
	}
	for _, g := range req.GetGames() {
		switch g {
		case hoyolib_pb.GameType_Genshin:
		}
	}
	m[u] = info
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
	return nil
}
