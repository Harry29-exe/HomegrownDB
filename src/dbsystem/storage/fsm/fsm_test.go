package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/tx"
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

	pageId := uint32(dbbs.PageSize)*2 + 7
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
	pageId := uint32(dbbs.PageSize)*5 + 7
	helper.testFsmUpdate(pageId, 255)
	helper.clear(pageId)
	helper.assertFind(5, 2)

	helper.clear(5)
	helper.assertNoFound(1)

	helper.clearAll()
}

func newFsmTestHelper(t *testing.T) *fsmTestHelper {
	inMemFile := dbfs.NewInMemoryFile("")
	fsMap, err := fsm.CreateTestFreeSpaceMap(inMemFile, t)
	assert.IsNil(err, t)

	return &fsmTestHelper{
		fsMap:   fsMap,
		t:       t,
		ctx:     nil,
		pageIds: make([]dbbs.PageId, 0, 10),
	}
}

type fsmTestHelper struct {
	fsMap *fsm.FreeSpaceMap
	t     *testing.T
	ctx   *tx.Ctx

	pageIds []dbbs.PageId
}

func (pt *fsmTestHelper) testFsmUpdate(pageId dbbs.PageId, newSize uint8) {
	pt.pageIds = append(pt.pageIds, pageId)

	size := pt.toAbsSize(newSize)

	err := pt.fsMap.UpdatePage(size, pageId)
	assert.IsNil(err, pt.t)
	foundPageId, err := pt.fsMap.FindPage(size, pt.ctx)
	if err != nil {
		pt.t.Error(err.Error())
	}

	assert.Eq(pageId, foundPageId, pt.t)
}

func (pt *fsmTestHelper) assertFind(id dbbs.PageId, size uint8) {
	absSize := pt.toAbsSize(size)
	page, err := pt.fsMap.FindPage(absSize, pt.ctx)
	assert.IsNil(err, pt.t)
	assert.Eq(page, id, pt.t)
}

func (pt *fsmTestHelper) assertNoFound(size uint8) {
	_, err := pt.fsMap.FindPage(pt.toAbsSize(size), pt.ctx)
	assert.NotNil(err, pt.t)
}

func (pt *fsmTestHelper) clear(id dbbs.PageId) {
	pt.fsMap.UpdatePage(0, id)
}

func (pt *fsmTestHelper) clearAll() {
	for _, id := range pt.pageIds {
		pt.fsMap.UpdatePage(0, id)
	}
	pt.pageIds = pt.pageIds[:0]

	pt.assertNoFound(1)
}

func (pt *fsmTestHelper) toAbsSize(compressedSize uint8) uint16 {
	divider := dbbs.PageSize / 256
	return uint16(compressedSize) * divider
}

func (pt *fsmTestHelper) close() {
}
