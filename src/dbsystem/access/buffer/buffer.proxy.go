package buffer

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
)

func NewSharedBuffer(buffSize uint, store *pageio.Store) SharedBuffer {
	return &bufferProxy{newBuffer(buffSize, store)}
}

type bufferProxy struct {
	buffer internalBuffer
}

var _ SharedBuffer = &bufferProxy{}

func (b *bufferProxy) RTablePage(table table.RDefinition, pageId page.Id) (tpage.RPage, error) {
	rPage, err := b.buffer.ReadRPage(table, pageId, rbmRead)
	if err != nil {
		return nil, err
	}

	return tpage.AsPage(rPage.bytes, pageId, table), nil
}

func (b *bufferProxy) WTablePage(table table.RDefinition, pageId page.Id) (tpage.WPage, error) {
	wPage, err := b.buffer.ReadWPage(table, pageId, rbmReadOrCreate)
	if err != nil {
		return nil, err
	}

	if wPage.isNew {
		return tpage.InitNewPage(table, pageId, wPage.bytes), nil
	} else {
		return tpage.AsPage(wPage.bytes, pageId, table), nil
	}
}

func (b *bufferProxy) RFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error) {
	rPage, err := b.buffer.ReadRPage(rel, pageId, rbmRead)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: rPage.bytes}, nil
}

func (b *bufferProxy) WFsmPage(rel relation.Relation, pageId page.Id) (fsmpage.Page, error) {
	wPage, err := b.buffer.ReadWPage(rel, pageId, rbmReadOrCreate)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: wPage.bytes}, nil
}

func (b *bufferProxy) WPageRelease(tag pageio.PageTag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) RPageRelease(tag pageio.PageTag) {
	b.buffer.ReleaseRPage(tag)
}
