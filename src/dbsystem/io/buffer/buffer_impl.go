package buffer

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema"
	"HomegrownDB/dbsystem/schema/table"
	"sync"
)

type buffer struct {
	bufferMap     map[bstructs.PageTag]ArrayIndex
	bufferMapLock sync.RWMutex

	descriptorArray []pageDescriptor
	pageBufferArray []byte
}

func (b *buffer) RPage(tag bstructs.PageTag) (bstructs.RPage, error) {
	tableDef := schema.Tables.Table(tag.TableId)
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.newUsage()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.RLock()
	} else {
		b.bufferMapLock.RUnlock()
		index, err := b.fetchPageAndRLock(tag, tableDef)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bstructs.PageSize)
	return bstructs.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bstructs.PageSize)]), nil
}

func (b *buffer) WPage(tag bstructs.PageTag) (bstructs.WPage, error) {
	tableDef := schema.Tables.Table(tag.TableId)
	b.bufferMapLock.RLock()

	pageArrIndex, ok := b.bufferMap[tag]
	var descriptor *pageDescriptor
	if ok {
		descriptor = &b.descriptorArray[pageArrIndex]
		descriptor.newUsage()
		b.bufferMapLock.RUnlock()

		descriptor.contentLock.Lock()
	} else {
		b.bufferMapLock.RUnlock()
		index, err := b.fetchPageAndWLock(tag, tableDef)
		if err != nil {
			return nil, err
		}

		descriptor = &b.descriptorArray[index]
	}

	pageStart := uintptr(pageArrIndex) * uintptr(bstructs.PageSize)
	return bstructs.NewPage(tableDef, b.pageBufferArray[pageStart:pageStart+uintptr(bstructs.PageSize)]), nil
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
func (b *buffer) fetchPageAndRLock(tag bstructs.PageTag, table table.Definition) (ArrayIndex, error) {
	panic("Not implemented")
}

// fetchPage fetches page with given tag from drive and increases it usage count
// by 1, so it can not instantly become victim page, therefore function invoking this
// method should not increase it
func (b *buffer) fetchPageAndWLock(tag bstructs.PageTag, table table.Definition) (ArrayIndex, error) {
	panic("Not implemented")
}
