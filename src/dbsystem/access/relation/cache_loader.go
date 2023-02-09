package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/dbsystem/storage/page"
	"errors"
)

func newLoaderCache(buffer buffer.SharedBuffer) *loaderCache {
	return &loaderCache{
		buffer:    buffer,
		relations: []reldef.Relation{},
		columns:   map[reldef.OID][]column.WDef{},
	}
}

type loaderCache struct {
	buffer    buffer.SharedBuffer
	relations []reldef.Relation
	columns   map[reldef.OID][]column.WDef
}

func (c *loaderCache) loadRelations() error {
	nextPageId := page.Id(0)

	for {
		err := c.loadRelPage(nextPageId)
		if err != nil {
			if errors.Is(err, canReadPageErr) {
				break
			} else {
				return err
			}
		}
		nextPageId++
	}
	return nil
}

func (c *loaderCache) loadRelPage(
	pageId page.Id,
) error {
	rPage, err := c.buffer.RTablePage(systable.RelationsTableDef(), pageId)
	if err != nil {
		return canReadPage{canReadPageErr}
	}
	defer c.buffer.RPageRelease(rPage.PageTag())

	for tupleIndex := page.TupleIndex(0); tupleIndex < rPage.TupleCount(); tupleIndex++ {
		tuple := rPage.Tuple(tupleIndex)
		relation := systable.RelationsOps.RowAsData(tuple)
		c.relations = append(c.relations, relation)
	}

	return nil
}

func (c *loaderCache) loadColumns() error {
	nextPageId := page.Id(0)

	for {
		err := c.loadColPage(nextPageId)
		if err != nil {
			if errors.Is(err, canReadPageErr) {
				break
			} else {
				return err
			}
		}
		nextPageId++
	}
	return nil
}

func (c *loaderCache) loadColPage(pageID page.Id) error {
	rPage, err := c.buffer.RTablePage(systable.ColumnsTableDef(), pageID)
	if err != nil {
		return canReadPage{canReadPageErr}
	}

	defer c.buffer.RPageRelease(rPage.PageTag())

	for tupleIndex := page.TupleIndex(0); tupleIndex < rPage.TupleCount(); tupleIndex++ {
		tuple := rPage.Tuple(tupleIndex)
		columnDef := systable.ColumnsOps.RowToData(tuple)
		relOID := systable.ColumnsOps.TableOID(tuple)

		columns, ok := c.columns[relOID]
		if !ok {
			c.columns[relOID] = []column.WDef{columnDef}
		} else {
			c.columns[relOID] = append(columns, columnDef)
		}
	}

	return nil
}
