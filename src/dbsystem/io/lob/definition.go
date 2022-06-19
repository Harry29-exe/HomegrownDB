package lob

import "HomegrownDB/datastructs"

type Id = uint64

var IdCounter idCounter

func init() {
	//todo save id counter state to disc
	// and init it from it
	IdCounter = idCounter{
		counter: datastructs.NewUint64SyncCounter(0),
	}
}

type idCounter struct {
	counter datastructs.SyncCounter[Id]
}

func (c idCounter) NextId() Id {
	return c.counter.IncrementAndGet()
}
