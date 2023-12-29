package controller

import (
	"fmt"
	"sync"

	"github.com/warp-contracts/syncer/src/utils/task"
)

type Cache struct {
	*task.Task

	input chan any

	cache []map[string]int64
	mtx   sync.RWMutex
}

func NewCache() (self *Cache) {
	self = new(Cache)

	self.input = make(chan any, 1000)

	self.Task = task.NewTask(nil, "limiter-cache").
		WithSubtaskFunc(self.run).
		WithOnAfterStop(func() {
			close(self.input)
		})

	return
}

func (self *Cache) GetCount(limiterIndex int, key string) int64 {
	self.mtx.RLock()
	defer self.mtx.RUnlock()

	return self.cache[limiterIndex][key]
}

// Try to increment the counter, but do not block if the channel is full
func (self *Cache) Increment(limiterIndex int, key string) {
	select {
	case self.input <- &MsgIncrement{LimiterIndex: limiterIndex, Key: key}:
	default:
	}
}

// Try to increment the counter, but do not block if the channel is full
func (self *Cache) Delete(limiterIndex int, key string) {
	select {
	case self.input <- &MsgIncrement{LimiterIndex: limiterIndex, Key: key}:
	case <-self.Ctx.Done():
	}
}

// Gets all messages stored in the input channel and applies them to the cache
func (self *Cache) getAllMessages(m any) error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	for {
		switch msg := m.(type) {
		case *MsgIncrement:
			self.cache[msg.LimiterIndex][string(msg.Key)] += 1
		case *MsgDelete:
			delete(self.cache[msg.LimiterIndex], string(msg.Key))
		default:
			return fmt.Errorf("unknown message type: %T", msg)
		}

		select {
		case <-self.Ctx.Done():
			return nil
		case m = <-self.input:
			continue
		default:
			// Prevents blocking. There are no more msgs in the channel
			break
		}
	}
}

func (self *Cache) run() error {
	for {
		select {
		case <-self.Ctx.Done():
			return nil
		case m := <-self.input:
			err := self.getAllMessages(m)
			if err != nil || self.IsStopping.Load() {
				return err
			}
		}
	}

}
