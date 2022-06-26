package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

type buffer struct {
	bufferMap     map[bstructs.PageTag]ArrayIndex
	bufferMapLock sync.RWMutex

	descriptorArray []pageDescriptor
	pageBufferArray []byte
}

func (b *buffer) RPage(tag bstructs.PageTag, table table.Definition) (bstructs.RPage, error) {
	b.bufferMapLock.Lock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.incrementRefCount()
		descriptor.contentLock.RLock()
		b.bufferMapLock.Unlock()
	} else {
		b.bufferMapLock.Unlock()
		index, err := b.fetchPageAndRLock(tag)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bstructs.PageSize)
	return bstructs.NewPage(table, b.pageBufferArray[pageStart:pageStart+uintptr(bstructs.PageSize)]), nil
}

func (b *buffer) WPage(tag bstructs.PageTag, table table.Definition) (bstructs.WPage, error) {
	//b.bufferMapLock.Lock()
	//
	//pageArrIndex, ok := b.bufferMap[]
	//var descriptor *pageDescriptor
	//if ok {
	//	descriptor = &b.descriptorArray[pageArrIndex]
	//	descriptor.
	//}
	//todo
	panic("not implemented")
}

func (b *buffer) ReleaseWPage(page bstructs.WPage) {
	panic("Not implemented")
}

func (b *buffer) ReleaseRPage(page bstructs.RPage) {
	panic("Not implemented")
}

// fetchPage fetches page with given tag from drive and increases it usage count
// by 1, so it can not instantly become victim page, therefore function invoking this
// method should not increase it
func (b *buffer) fetchPageAndRLock(tag bstructs.PageTag) (ArrayIndex, error) {
	panic("Not implemented")
}

// fetchPage fetches page with given tag from drive and increases it usage count
// by 1, so it can not instantly become victim page, therefore function invoking this
// method should not increase it
func (b *buffer) fetchPageAndWLock(tag bstructs.PageTag) (ArrayIndex, error) {
	panic("Not implemented")
}
