package handler

import (
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
)

var (
	_ hoyolib_pb.HoyolibServer = (*HoyolibServer)(nil)
)

type HoyolibServer struct {
	hoyolib_pb.UnimplementedHoyolibServer
}
