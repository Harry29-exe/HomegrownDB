package buffer

import (
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
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

func (b *bufferProxy) RTablePage(table table.RDefinition, pageId page.Id) (page.RPage, error) {
	rPage, err := b.buffer.ReadRPage(table.OID(), pageId, RbmRead)
	if err != nil {
		return nil, err
	}

	return page.AsTablePage(rPage.Bytes, pageId, table), nil
}

func (b *bufferProxy) WTablePage(table table.RDefinition, pageId page.Id) (page.WPage, error) {
	wPage, err := b.buffer.ReadWPage(table.OID(), pageId, RbmReadOrCreate)
	if err != nil {
		return nil, err
	}

	if wPage.IsNew {
		return page.InitNewTablePage(wPage.Bytes, table, wPage.Tag.PageId), nil
	} else {
		return page.AsTablePage(wPage.Bytes, wPage.Tag.PageId, table), nil
	}
}

func (b *bufferProxy) RFsmPage(ownerID dbobj.OID, pageId page.Id) (fsmpage.Page, error) {
	rPage, err := b.buffer.ReadRPage(ownerID, pageId, RbmRead)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: rPage.Bytes}, nil
}

func (b *bufferProxy) WFsmPage(ownerID dbobj.OID, pageId page.Id) (fsmpage.Page, error) {
	wPage, err := b.buffer.ReadWPage(ownerID, pageId, RbmReadOrCreate)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: wPage.Bytes}, nil
}

func (b *bufferProxy) WPageRelease(tag page.PageTag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) RPageRelease(tag page.PageTag) {
	b.buffer.ReleaseRPage(tag)
}
