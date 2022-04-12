package locker

import "sync"

type PageId struct {
	tableId    uint32
	pageNumber uint32
	nodeNumber uint32
}

var lockedPages = map[PageId]*sync.Mutex{}

func LockPage(id PageId) {
	mutex, ok := lockedPages[id]
	if !ok {
		panic("page does not exist")
	}

	mutex.Lock()
}

func UnlockPage(id PageId) {
	mutex, ok := lockedPages[id]
	if !ok {
		panic("page does not exist")
	}

	mutex.Unlock()
}
