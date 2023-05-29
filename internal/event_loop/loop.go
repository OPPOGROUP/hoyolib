package event_loop

import (
	"github.com/OPPOGROUP/hoyolib/internal/handler"
	"github.com/OPPOGROUP/protocol/hoyolib_pb"
	"github.com/robfig/cron"
)

var (
	c = cron.New()
)

func Start() error {
	err := c.AddFunc("0 0 0 * * *", autoCheckIn)
	if err != nil {
		return err
	}
	err = c.AddFunc("0 0 0 * * *", updateAccountInfo)
	if err != nil {
		return err
	}
	c.Start()
	defer c.Stop()
	return nil
}

func autoCheckIn() {
	resMap := make(map[int64]*hoyolib_pb.CheckInResponse)
	for uid := range handler.GetUserData() {
		resp, _ := handler.CheckInUser(uid)
		resMap[uid] = resp
	}
}

func updateAccountInfo() {

}
