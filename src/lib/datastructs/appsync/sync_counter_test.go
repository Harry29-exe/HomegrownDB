package appsync_test

import (
	"HomegrownDB/lib/datastructs/appsync"
	"HomegrownDB/lib/tests/assert"
	"sync"
	"testing"
)

func TestUint32SyncCounter(t *testing.T) {
	counter := appsync.NewUint32SyncCounter(0)
	waitGroup := sync.WaitGroup{}

	increments := 100_000
	processes := 8
	var incrementFunc = func() {
		for i := 0; i < increments; i++ {
			counter.IncrAndGet()
		}
		waitGroup.Done()
	}

	waitGroup.Add(processes)
	for i := 0; i < processes; i++ {
		go incrementFunc()
	}
	waitGroup.Wait()

	counterValue := counter.Get()
	assert.Eq(counterValue, uint32(increments*processes), t)
}
