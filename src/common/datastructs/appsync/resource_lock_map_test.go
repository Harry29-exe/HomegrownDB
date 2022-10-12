package appsync

import (
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestResLockMap_RLockUnlock(t *testing.T) {
	lock := NewResLockMap[uint]()

	lock.RLockRes(1)
	res1, ok := lock.locksMap[1]
	assert.Eq(ok, true, t)
	assert.Eq(res1.accessing, 1, t)

	lock.RLockRes(75)
	res2, ok := lock.locksMap[75]
	assert.Eq(ok, true, t)
	assert.Eq(res2.accessing, 1, t)

	lock.RLockRes(1)
	assert.Eq(res1.accessing, 2, t)

	lock.RUnlockRes(1)
	assert.Eq(res1.accessing, 1, t)
	lock.RUnlockRes(1)
	assert.Eq(res1.accessing, 0, t)
	_, ok = lock.locksMap[1]
	assert.Eq(ok, false, t)

	lock.RUnlockRes(75)
	assert.Eq(res1.accessing, 0, t)
	_, ok = lock.locksMap[75]
	assert.Eq(ok, false, t)
}

func TestResLockMap_WLockUnlock(t *testing.T) {
	lock := NewResLockMap[uint]()

	lock.WLockRes(1)
	res1, ok := lock.locksMap[1]
	assert.Eq(ok, true, t)
	assert.Eq(res1.accessing, 1, t)

	lock.WLockRes(75)
	res2, ok := lock.locksMap[75]
	assert.Eq(ok, true, t)
	assert.Eq(res2.accessing, 1, t)

	lock.WUnlockRes(75)
	_, ok = lock.locksMap[75]
	assert.Eq(ok, false, t)

	lock.WUnlockRes(1)
	_, ok = lock.locksMap[1]
	assert.Eq(ok, false, t)
}
