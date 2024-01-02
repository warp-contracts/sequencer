package controller

import (
	"fmt"
	"sync"

	"github.com/warp-contracts/syncer/src/utils/task"
)

type Cache struct {
	*task.Task

	// Queue of messages to be processed
	input chan any

	cache []map[string]int64
	mtx   sync.RWMutex
}

func NewCache(numLimiters int) (self *Cache) {
	self = new(Cache)

	self.input = make(chan any, 1000)
	self.cache = make([]map[string]int64, numLimiters)

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

// Try to set the counter, but do not block if the channel is full
// Worse case this operation is skipped, but the cache will be updated upon the next try
func (self *Cache) Set(limiterIndex int, key string, value int64) {
	select {
	case self.input <- &MsgSet{LimiterIndex: limiterIndex, Key: key}:
	default:
	}
}

func (self *Cache) Subtract(limiterIndex int, key string, value int64) {
	select {
	case self.input <- &MsgSubtract{LimiterIndex: limiterIndex, Key: key, Value: value}:
	case <-self.Ctx.Done():
	}
}

// Gets all messages stored in the input channel and applies them to the cache
func (self *Cache) getAllMessages(m any) error {
	self.mtx.Lock()
	defer self.mtx.Unlock()

	for {
		switch msg := m.(type) {
		case *MsgSet:
			self.cache[msg.LimiterIndex][string(msg.Key)] = msg.Value
		case *MsgSubtract:
			newValue := self.cache[msg.LimiterIndex][string(msg.Key)] - msg.Value
			if newValue <= 0 {
				delete(self.cache[msg.LimiterIndex], string(msg.Key))
			} else {
				self.cache[msg.LimiterIndex][string(msg.Key)] = newValue
			}
		default:
			return fmt.Errorf("unknown message type: %T", msg)
		}

		select {
		case <-self.Ctx.Done():
			return nil
		case m = <-self.input:
			continue
		default:
			// Prevents blocking.
			// There are no more msgs in the channel, break the loop
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
