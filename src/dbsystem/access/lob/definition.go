package lob

import (
	"HomegrownDB/datastructs/appsync"
)

type Id = uint64

var IdCounter idCounter

func init() {
	//todo save id counter state to disc
	// and init it from it
	IdCounter = idCounter{
		counter: appsync.NewUint64SyncCounter(0),
	}
}

type idCounter struct {
	counter appsync.SyncCounter[Id]
}

func (c idCounter) NextId() Id {
	return c.counter.IncrementAndGet()
}
