package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/access/fsm"
	"HomegrownDB/dbsystem/dbbs"
	"github.com/google/uuid"
	"os"
	"testing"
)

func TestFreeSpaceMap_UpdatePage(t *testing.T) {
	// given
	newUUID, _ := uuid.NewUUID()
	filename := os.TempDir() + "/" + newUUID.String()

	defer func() {
		_ = os.Remove(filename)
	}()

	freeSpaceMap := fsm.CreateFreeSpaceMap(filename)

	givenPageId, space := uint32(0), uint16(512)
	freeSpaceMap.UpdatePage(space, givenPageId)
	receivedPageId, err := freeSpaceMap.FindPage(space, nil)
	assert.IsNil(err, t)
	assert.Eq(receivedPageId, givenPageId, t)

	givenPageId, space = uint32(3), uint16(768)
	freeSpaceMap.UpdatePage(space, givenPageId)
	receivedPageId, err = freeSpaceMap.FindPage(space, nil)
	assert.IsNil(err, t)
	assert.Eq(receivedPageId, givenPageId, t)

	givenPageId, space = uint32(dbbs.PageSize+5), uint16(1022)
	freeSpaceMap.UpdatePage(space, givenPageId)
	receivedPageId, err = freeSpaceMap.FindPage(space, nil)
	assert.IsNil(err, t)
	assert.Eq(receivedPageId, givenPageId, t)

	biggestPageId, biggestSpace := givenPageId, space
	givenPageId, space = uint32(dbbs.PageSize*2+5), uint16(1021)
	freeSpaceMap.UpdatePage(space, givenPageId)
	receivedPageId, err = freeSpaceMap.FindPage(biggestSpace, nil)
	assert.IsNil(err, t)
	assert.Eq(receivedPageId, biggestPageId, t)
}
