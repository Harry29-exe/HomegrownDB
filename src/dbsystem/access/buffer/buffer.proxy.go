package buffer

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
	"HomegrownDB/dbsystem/storage/tpage"
)

func NewSharedBuffer(buffSize uint, store *pageio.Store) SharedBuffer {
	return &bufferProxy{newSharedBuffer(buffSize, store)}
}

type bufferProxy struct {
	buffer sharedBuffer
}

var _ SharedBuffer = &bufferProxy{}

func (b *bufferProxy) RTablePage(pageId page.Id, table table.Definition) (tpage.TableRPage, error) {
	rPage, err := b.buffer.RPage(page.Tag{PageId: pageId, Relation: table.RelationId()})
	if err != nil {
		return nil, err
	}

	return tpage.AsPage(rPage.bytes, pageId, table), nil
}

func (b *bufferProxy) WTablePage(pageId page.Id, table table.Definition) (tpage.TableWPage, error) {
	wPage, err := b.buffer.WPage(page.Tag{PageId: pageId, Relation: table.RelationId()})
	if err != nil {
		return nil, err
	}

	if wPage.isNew {
		return tpage.InitNewPage(table, wPage.bytes), nil
	} else {
		return tpage.AsPage(wPage.bytes, pageId, table), nil
	}
}

func (b *bufferProxy) RFsmPage(tag page.Tag) (fsmpage.Page, error) {
	rPage, err := b.buffer.RPage(tag)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: rPage.bytes}, nil
}

func (b *bufferProxy) WFsmPage(tag page.Tag) (fsmpage.Page, error) {
	wPage, err := b.buffer.WPage(tag)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: wPage.bytes}, nil
}

func (b *bufferProxy) WPageRelease(tag page.Tag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) RPageRelease(tag page.Tag) {
	b.buffer.ReleaseRPage(tag)
}
