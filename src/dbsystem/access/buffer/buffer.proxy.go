package buffer

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/page"
)

type bufferProxy struct {
	buffer genericBuffer
}

var _ SharedBuffer = &bufferProxy{}

func (b *bufferProxy) TableRPage(tag PageTag, table table.Definition) (page.TableRPage, error) {
	rPage, err := b.buffer.RPage(tag)
	if err != nil {
		return nil, err
	}
	return page.NewPage(table, rPage), nil
}

func (b *bufferProxy) TableWPage(tag PageTag, table table.Definition) (page.TableWPage, error) {
	wPage, err := b.buffer.WPage(tag)
	if err != nil {
		return nil, err
	}
	return page.NewPage(table, wPage), nil
}

func (b *bufferProxy) RGenericPage(tag PageTag, relation relation.Relation) (page.GenericPage, error) {
	rPage, err := b.buffer.RPage(tag)
	if err != nil {
		return page.GenericPage{}, err
	}
	return page.NewGenericPage(rPage, relation), nil
}

func (b *bufferProxy) WGenericPage(tag PageTag, relation relation.Relation) (page.GenericPage, error) {
	wPage, err := b.buffer.WPage(tag)
	if err != nil {
		return page.GenericPage{}, err
	}
	return page.NewGenericPage(wPage, relation), nil
}

func (b *bufferProxy) ReleaseWPage(tag PageTag) {
	b.buffer.ReleaseWPage(tag)
}

func (b *bufferProxy) ReleaseRPage(tag PageTag) {
	b.buffer.ReleaseRPage(tag)
}
