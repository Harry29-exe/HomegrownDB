package hgtest

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/storage/dbfs"
)

type TestCtx struct {
	FS        dbfs.FS
	Relations []relation.Relation
}

func CreateNewDBSystem() {

}
