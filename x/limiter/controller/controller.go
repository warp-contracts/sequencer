package controller

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/sirupsen/logrus"

	"github.com/warp-contracts/syncer/src/utils/task"
)

type Controller struct {
	*task.Task

	Cache *Cache
}

func NewController(log log.Logger, homePath string) (self *Controller, err error) {
	self = new(Controller)
	InitLogger(log, logrus.InfoLevel.String())

	// Setup the tasks
	self.Task = task.NewTask(nil, "limiter-controller")

	self.Cache = NewCache()

	self.Task = self.Task.
		WithSubtask(self.Cache.Task)

	// Starts all the tasks
	err = self.Start()
	if err != nil {
		return
	}

	return
}
