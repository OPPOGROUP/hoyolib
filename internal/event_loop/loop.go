package event_loop

import (
	"context"
	"github.com/OPPOGROUP/hoyolib/internal/cte"
	"github.com/OPPOGROUP/hoyolib/internal/errors"
	"github.com/OPPOGROUP/hoyolib/internal/handler"
	"github.com/OPPOGROUP/hoyolib/internal/log"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	c      = cron.New()
	client hoyolib_pb.OpwxClient
	enable = false
)

func init() {
	conn, err := grpc.Dial(cte.WxApi, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msg("grpc dial failed")
	}
	client = hoyolib_pb.NewOpwxClient(conn)
	enable = true
}

func Start() error {
	if !enable {
		return errors.ErrGrpcClientInitFailed
	}
	_, _ = c.AddFunc("@every 1m", autoCheckIn)
	//_, _ = c.AddFunc("@hourly", updateAccountInfo)

	c.Start()
	log.Info().Msg("event loop start")
	log.Debug().Msgf("event loop entries: %v", c.Entries())
	return nil
}

func autoCheckIn() {
	if len(handler.GetUserData()) == 0 {
		log.Info().Msg("no user data, skip auto checkin")
		return
	}
	checkInResults := make([]*hoyolib_pb.CheckInResponse, 0, len(handler.GetUserData()))
	log.Debug().Any("user_data", handler.GetUserData()).Msg("autoCheckIn user data")
	for uid := range handler.GetUserData() {
		resp, _ := handler.CheckInUser(uid)
		checkInResults = append(checkInResults, resp)
	}
	req := &hoyolib_pb.CheckinResults{
		Results: checkInResults,
	}
	log.Debug().Msgf("[grpc] checkin results request: %v", req)
	resp, err := client.PushCheckinResults(context.Background(), req)
	log.Debug().Msgf("[grpc] checkin results response: %v", resp)
	if err != nil {
		log.Error().Err(err).Msg("[grpc] push checkin results failed")
		return
	}
	if resp.GetStatus() != hoyolib_pb.PushResponse_OK {
		log.Error().Msgf("[grpc] push checkin results failed: %s", resp.GetMsg())
	}
}

func updateAccountInfo() {

}
