package buffer

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm/fsmpage"
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

func (b *bufferProxy) TableRPage(tag PageTag, table table.Definition) (tpage.TableRPage, error) {
	rPage, err := b.buffer.RPage(tag)
	if err != nil {
		return nil, err
	}
	return tpage.NewPage(table, rPage.bytes), nil
}

func (b *bufferProxy) TableWPage(tag PageTag, table table.Definition) (tpage.TableWPage, error) {
	wPage, err := b.buffer.WPage(tag)
	if err != nil {
		return nil, err
	}
	return tpage.NewPage(table, wPage.bytes), nil
}

func (b *bufferProxy) RFsmPage(tag PageTag) (fsmpage.Page, error) {
	rPage, err := b.buffer.RPage(tag)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: rPage.bytes}, nil
}

func (b *bufferProxy) WFsmPage(tag PageTag) (fsmpage.Page, error) {
	wPage, err := b.buffer.WPage(tag)
	if err != nil {
		return fsmpage.Page{}, err
	}
	return fsmpage.Page{Bytes: wPage.bytes}, nil
}

func (b *bufferProxy) ReleaseWPage(tag PageTag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) ReleaseRPage(tag PageTag) {
	b.buffer.ReleaseRPage(tag)
}
