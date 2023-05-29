package main

import (
	"fmt"
	"github.com/OPPOGROUP/hoyolib/internal/config"
	"github.com/OPPOGROUP/hoyolib/internal/event_loop"
	"github.com/OPPOGROUP/hoyolib/internal/handler"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("config init failed")
	}
	err = log.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("log init failed")
	}
	if viper.GetBool("data.enable") {
		handler.LoadSavedUsers()
	}
	err = event_loop.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("event loop start failed")
	}
	if err = startServer(); err != nil {
		log.Fatal().Err(err).Msg("start server failed")
	}
}

func startServer() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	hoyolib_pb.RegisterHoyolibServer(server, &handler.HoyolibServer{})
	log.Info().Int("port", viper.GetInt("port")).Msg("hoyolib server start")
	if err = server.Serve(listen); err != nil {
		return err
	}
	return nil
}
