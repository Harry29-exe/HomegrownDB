package buffer

import (
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
)

func NewSharedBuffer(buffSize uint, store pageio.Store) SharedBuffer {
	return &bufferProxy{NewStdBuffer(buffSize, store)}
}

func AsSharedBuffer(buff StdBuffer) SharedBuffer {
	return &bufferProxy{buff}
}

type bufferProxy struct {
	buffer StdBuffer
}

var _ SharedBuffer = &bufferProxy{}

func (b *bufferProxy) RTablePage(table table.RDefinition, pageId page.Id) (tpage.RPage, error) {
	rPage, err := b.buffer.ReadRPage(table, pageId, RbmRead)
	if err != nil {
		return nil, err
	}

	return tpage.AsPage(rPage.Bytes, pageId, table), nil
}

func (b *bufferProxy) WTablePage(table table.RDefinition, pageId page.Id) (tpage.WPage, error) {
	wPage, err := b.buffer.ReadWPage(table, pageId, RbmReadOrCreate)
	if err != nil {
		return nil, err
	}

	if wPage.IsNew {
		return tpage.InitNewPage(table, pageId, wPage.Bytes), nil
	} else {
		return tpage.AsPage(wPage.Bytes, pageId, table), nil
	}
}

func (b *bufferProxy) RFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error) {
	rPage, err := b.buffer.ReadRPage(rel, pageId, RbmRead)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: rPage.Bytes}, nil
}

func (b *bufferProxy) WFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error) {
	wPage, err := b.buffer.ReadWPage(rel, pageId, RbmReadOrCreate)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: wPage.Bytes}, nil
}

func (b *bufferProxy) WPageRelease(tag pageio.PageTag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) RPageRelease(tag pageio.PageTag) {
	b.buffer.ReleaseRPage(tag)
}
