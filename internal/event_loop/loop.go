package event_loop

import "github.com/robfig/cron"

var (
	c = cron.New()
)

func Start() {
	c.Start()

}

func checkIn() {

}
