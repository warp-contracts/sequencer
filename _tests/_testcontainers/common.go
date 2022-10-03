package _testcontainers

import (
	"github.com/sirupsen/logrus"
	"sync"
)

// Use this struct for close containers in go routines
type countWaitGroup struct {
	wg    sync.WaitGroup
	count int
}

func (g *countWaitGroup) Add(count int) int {
	g.wg.Add(count)
	g.count += count
	return g.count
}

func (g *countWaitGroup) Done() bool {
	if g.count == 0 {
		logrus.Panic("Can't call Done when count is 0.")
	}
	g.wg.Done()
	g.count--
	return g.count == 0
}

func (g *countWaitGroup) Wait() {
	g.wg.Wait()
}
