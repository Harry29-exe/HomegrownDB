package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/access/fsm"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/tx"
	"github.com/google/uuid"
	"os"
	"testing"
)

//todo optimize to not write 32MB to disc every time
func TestFreeSpaceMap_UpdatePage(t *testing.T) {
	fsmTests := newUpdatePageTests(t)
	defer fsmTests.close()

	fsmTests.testFsmUpdate(1, 1)
	fsmTests.testFsmUpdate(2, 2)
	fsmTests.testFsmUpdate(0, 3)
	fsmTests.clear(0)
	fsmTests.assertFind(2, 2)
	pageId := uint32(dbbs.PageSize)*2 + 7
	fsmTests.testFsmUpdate(pageId, 8)
	fsmTests.clear(pageId)
	fsmTests.assertFind(2, 2)
	fsmTests.testFsmUpdate(1, 5)
	fsmTests.clear(1)
	fsmTests.assertFind(2, 2)

	fsmTests.clearAll()
}

func newUpdatePageTests(t *testing.T) *updatePageTests {
	newUUID, _ := uuid.NewUUID()
	filename := os.TempDir() + "/" + newUUID.String()

	return &updatePageTests{
		fsMap:    fsm.CreateFreeSpaceMap(filename),
		t:        t,
		ctx:      nil,
		filename: filename,
		pageIds:  make([]dbbs.PageId, 0, 10),
	}
}

type updatePageTests struct {
	fsMap *fsm.FreeSpaceMap
	t     *testing.T
	ctx   *tx.Ctx

	filename string
	pageIds  []dbbs.PageId
}

func (pt *updatePageTests) testFsmUpdate(pageId dbbs.PageId, newSize uint8) {
	pt.pageIds = append(pt.pageIds, pageId)

	size := pt.toAbsSize(newSize)

	pt.fsMap.UpdatePage(size, pageId)
	foundPageId, err := pt.fsMap.FindPage(size, pt.ctx)
	if err != nil {
		pt.t.Error(err.Error())
	}

	assert.Eq(pageId, foundPageId, pt.t)
}

func (pt *updatePageTests) assertFind(id dbbs.PageId, size uint8) {
	absSize := pt.toAbsSize(size)
	page, err := pt.fsMap.FindPage(absSize, pt.ctx)
	assert.IsNil(err, pt.t)
	assert.Eq(page, id, pt.t)
}

func (pt *updatePageTests) clear(id dbbs.PageId) {
	pt.fsMap.UpdatePage(0, id)
}

func (pt *updatePageTests) clearAll() {
	for _, id := range pt.pageIds {
		pt.fsMap.UpdatePage(0, id)
	}
	pt.pageIds = pt.pageIds[:0]

	_, err := pt.fsMap.FindPage(1, pt.ctx)
	assert.NotNil(err, pt.t)
}

func (pt *updatePageTests) toAbsSize(compressedSize uint8) uint16 {
	divider := dbbs.PageSize / 256
	return uint16(compressedSize) * divider
}

func (pt *updatePageTests) close() {
	err := os.Remove(pt.filename)
	if err != nil {
		pt.t.Error(err.Error())
	}
}
