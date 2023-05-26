package event_loop

import "github.com/robfig/cron"

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

}

func updateAccountInfo() {

}
