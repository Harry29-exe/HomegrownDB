package fsm_test

import (
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/access/fsm"
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
	freeSpaceMap.UpdatePage(512, 0)

	pageId, err := freeSpaceMap.FindPage(512, nil)
	assert.IsNil(err, t)
	assert.Eq(pageId, 0, t)
}
