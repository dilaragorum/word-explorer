package pkg

import (
	"github.com/go-co-op/gocron"
	"github.com/labstack/gommon/log"
	"time"
)

type CronClient interface {
	Start()
	Schedule(interval string, fn func())
	Stop()
}

type cronClient struct {
	sch *gocron.Scheduler
}

func NewCronClient() CronClient {
	location, _ := time.LoadLocation("Europe/Istanbul")
	return &cronClient{
		sch: gocron.NewScheduler(location),
	}
}

func (c *cronClient) Start() {
	c.sch.StartAsync()
}

func (c *cronClient) Schedule(interval string, fn func()) {
	_, err := c.sch.Every(interval).Do(func() {
		fn()
	})
	if err != nil {
		log.Error("Error scheduling cron " + err.Error())
	}
}

func (c *cronClient) Stop() {
	c.sch.Stop()
}
