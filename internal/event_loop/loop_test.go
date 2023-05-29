package event_loop

import (
	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCron(t *testing.T) {
	c := cron.New()
	_, err := c.AddFunc("@every 1s", func() {
		t.Log("1s")
	})
	assert.Nil(t, err)
	_, err = c.AddFunc("@every 5s", func() {
		t.Log("5s")
	})
	_, err = c.AddFunc("@every 1m", func() {
		t.Log("1m")
	})
	_, err = c.AddFunc("@hourly", func() {
		t.Log("hourly")
	})
	_, err = c.AddFunc("@daily", func() {
		t.Log("daily")
	})
	c.Start()
	for _, e := range c.Entries() {
		t.Log(e)
	}
	c.Stop()
}
