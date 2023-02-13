package hgtest

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

var PageIOUtils = pageIOUtils{}

type pageIOUtils struct{}

func (u pageIOUtils) With(t *testing.T, fs dbfs.FS, relations ...reldef.Relation) pageio.Store {
	store := pageio.NewStore(fs)
	for _, rel := range relations {
		err := store.Load(rel.OID())
		assert.ErrIsNil(err, t)
	}
	return store
}
