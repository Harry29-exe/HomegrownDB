package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
)

var Buffer DBBuffer = &buffer{
	bufferMap:       map[bstructs.PageTag]ArrayIndex{},
	descriptorArray: make([]pageDescriptor, 0, bufferSize),
	pageBufferArray: make([]byte, 0, int64(bstructs.PageSize)*bufferSize),
}

type DBBuffer interface {
	RPage(tag bstructs.PageTag) (bstructs.RPage, error)
	WPage(id bstructs.PageId) (bstructs.WPage, error)

	ReleaseWPage(page bstructs.WPage)
	ReleaseRPage(page bstructs.RPage)
}

const bufferSize = 10_000

type ArrayIndex = uint
