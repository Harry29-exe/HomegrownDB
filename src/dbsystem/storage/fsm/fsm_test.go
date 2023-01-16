package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/access/buffer"
	relation "HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"github.com/spf13/afero"
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
	usersTable := tt_user.Def(t)
	fsmRelation := relation.NewBaseRelation(0, relation.TypeFsm, "/", 0)
	fs := afero.NewMemMapFs()
	file, err := fs.Create("fsm_pageio")
	assert.IsNil(err, t)

	store := pageio.NewStore(hgtest.CreateAndInitTestFS(t))
	io, err := pageio.NewPageIO(file)
	assert.IsNil(err, t)
	store.Register(fsmRelation.RelationID(), io)
	buff := buffer.NewSharedBuffer(10_000, store)

	fsMap, err := fsm.CreateFreeSpaceMap(fsmRelation, usersTable.RelationID(), buff)
	assert.IsNil(err, t)

	return &fsmTestHelper{
		fsMap:   fsMap,
		t:       t,
		tx:      nil,
		pageIds: make([]page.Id, 0, 10),
	}
}

type fsmTestHelper struct {
	fsMap *fsm.FreeSpaceMap
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
	if err != nil {
		pt.t.Error(err.Error())
	}

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
