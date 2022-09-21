package _testcontainers

import "sync"

var arLock = &sync.Mutex{}
var arWg sync.WaitGroup

func RunArweaveContainer() {
	arLock.Lock()
	defer arLock.Unlock()

	arWg.Add(1)
}
