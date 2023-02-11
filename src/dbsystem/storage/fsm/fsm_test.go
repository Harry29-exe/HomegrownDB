package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"testing"
)

func TestFreeSpaceMap_UpdatePage(t *testing.T) {
	helper := newFsmTestHelper(t)
	defer helper.close()

	helper.testFsmUpdate(1, 1)
	helper.testFsmUpdate(2, 2)
	helper.testFsmUpdate(0, 3)
	helper.clear(0)
	helper.assertFind(2, 2)

	pageId := uint32(page.Size)*2 + 7
	helper.testFsmUpdate(pageId, 8)
	helper.clear(pageId)

	helper.assertFind(2, 2)
	helper.testFsmUpdate(1, 5)
	helper.clear(1)
	helper.assertFind(2, 2)

	helper.clearAll()
}

func TestFreeSpaceMap_UpdatePage2(t *testing.T) {
	helper := newFsmTestHelper(t)
	defer helper.close()

	helper.testFsmUpdate(5, 2)
	pageId := uint32(page.Size)*5 + 7
	helper.testFsmUpdate(pageId, 255)
	helper.clear(pageId)
	helper.assertFind(5, 2)

	helper.clear(5)
	helper.assertNoFound(1)

	helper.clearAll()
}

func newFsmTestHelper(t *testing.T) *fsmTestHelper {
	dbUtils := hgtest.CreateAndLoadDBWith(nil, t).
		WithUsersTable().
		Build()

	table := dbUtils.TableByName(tt_user.TableName)
	tableFsm := fsm.NewFSM(table.FsmOID(), dbUtils.DB.AccessModule().SharedBuffer())

	return &fsmTestHelper{
		fsMap:   tableFsm,
		t:       t,
		tx:      nil,
		pageIds: make([]page.Id, 0, 10),
	}
}

type fsmTestHelper struct {
	fsMap *fsm.FSM
	t     *testing.T
	tx    tx.Tx

	pageIds []page.Id
}

func (pt *fsmTestHelper) testFsmUpdate(pageId page.Id, newSize uint8) {
	pt.pageIds = append(pt.pageIds, pageId)

	size := pt.toAbsSize(newSize)

	err := pt.fsMap.UpdatePage(size, pageId)
	assert.IsNil(err, pt.t)
	foundPageId, err := pt.fsMap.FindPage(size, pt.tx)
	assert.ErrIsNil(err, pt.t)

	assert.Eq(pageId, foundPageId, pt.t)
}

func (pt *fsmTestHelper) assertFind(id page.Id, size uint8) {
	absSize := pt.toAbsSize(size)
	page, err := pt.fsMap.FindPage(absSize, pt.tx)
	assert.IsNil(err, pt.t)
	assert.Eq(page, id, pt.t)
}

func (pt *fsmTestHelper) assertNoFound(size uint8) {
	pageId, err := pt.fsMap.FindPage(pt.toAbsSize(size), pt.tx)
	assert.Eq(page.InvalidId, pageId, pt.t)
	assert.IsNil(err, pt.t)
}

func (pt *fsmTestHelper) clear(id page.Id) {
	err := pt.fsMap.UpdatePage(0, id)
	assert.IsNil(err, pt.t)
}

func (pt *fsmTestHelper) clearAll() {
	for _, id := range pt.pageIds {
		err := pt.fsMap.UpdatePage(0, id)
		assert.IsNil(err, pt.t)
	}
	pt.pageIds = pt.pageIds[:0]

	pt.assertNoFound(1)
}

func (pt *fsmTestHelper) toAbsSize(compressedSize uint8) uint16 {
	divider := page.Size / 256
	return uint16(compressedSize) * divider
}

func (pt *fsmTestHelper) close() {
}
