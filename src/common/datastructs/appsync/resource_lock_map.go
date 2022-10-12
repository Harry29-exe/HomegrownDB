package appsync

import (
	"sync"
)

func NewResLockMap[ID comparable]() *ResLockMap[ID] {
	return &ResLockMap[ID]{
		locksMap: map[ID]*resLock{},
		mapLock:  &sync.Mutex{},
	}
}

type ResLockMap[ID comparable] struct {
	locksMap map[ID]*resLock
	mapLock  *sync.Mutex
}

func (m *ResLockMap[ID]) RLockRes(resId ID) {
	m.mapLock.Lock()

	lock, ok := m.locksMap[resId]
	if ok {
		lock.accessing += 1
		m.mapLock.Unlock()

		lock.resLock.RLock()
	} else {
		newLock := newResLock()
		newLock.resLock.RLock()
		newLock.accessing++

		m.locksMap[resId] = newLock
		m.mapLock.Unlock()
	}
}

func (m *ResLockMap[ID]) RUnlockRes(resId ID) {
	m.mapLock.Lock()

	lock := m.locksMap[resId]
	lock.resLock.RUnlock()
	lock.accessing--
	if lock.accessing == 0 {
		delete(m.locksMap, resId)
	}
	m.mapLock.Unlock()
}

func (m *ResLockMap[ID]) WLockRes(resId ID) {
	m.mapLock.Lock()

	lock, ok := m.locksMap[resId]
	if ok {
		lock.accessing++
		m.mapLock.Unlock()

		lock.resLock.Lock()
	} else {
		lock = newResLock()
		lock.resLock.Lock()
		lock.accessing++

		m.locksMap[resId] = lock
		m.mapLock.Unlock()
	}
}

func (m *ResLockMap[ID]) WUnlockRes(resId ID) {
	m.mapLock.Lock()

	lock := m.locksMap[resId]
	lock.resLock.Unlock()
	lock.accessing--
	if lock.accessing == 0 {
		delete(m.locksMap, resId)
	}
	m.mapLock.Unlock()
}

func newResLock() *resLock {
	return &resLock{
		resLock:   sync.RWMutex{},
		accessing: 0,
	}
}

type resLock struct {
	resLock sync.RWMutex

	accessing int32
}
