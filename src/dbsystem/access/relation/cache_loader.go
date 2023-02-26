package relation

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"errors"
)

func newLoaderCache(buffer buffer.SharedBuffer) *loaderCache {
	return &loaderCache{
		buffer:    buffer,
		relations: []reldef.BaseRelation{},
		columns:   map[reldef.OID][]reldef.ColumnDefinition{},
		sequences: map[reldef.OID]systable.FutureSequence{},
	}
}

type loaderCache struct {
	buffer    buffer.SharedBuffer
	relations []reldef.BaseRelation
	columns   map[reldef.OID][]reldef.ColumnDefinition
	sequences map[reldef.OID]systable.FutureSequence
}

func (c *loaderCache) load() error {
	err := c.loadPages(c.loadRelPage)
	if err != nil {
		return err
	}
	err = c.loadPages(c.loadColPage)
	if err != nil {
		return err
	}
	err = c.loadPages(c.loadSequencePage)
	if err != nil {
		return err
	}

	return nil
}

func (c *loaderCache) loadRelPage(
	pageId page.Id,
) error {
	rPage, err := c.buffer.RTablePage(systable.RelationsTableDef(), pageId)
	if err != nil {
		return canNotReadPage{err}
	}
	defer c.buffer.RPageRelease(rPage.PageTag())

	for tupleIndex := page.TupleIndex(0); tupleIndex < rPage.TupleCount(); tupleIndex++ {
		tuple := rPage.Tuple(tupleIndex)
		relation := systable.RelationsOps.RowAsData(tuple)
		c.relations = append(c.relations, relation)
	}

	return nil
}

func (c *loaderCache) loadColPage(pageID page.Id) error {
	rPage, err := c.buffer.RTablePage(systable.ColumnsTableDef(), pageID)
	if err != nil {
		return canNotReadPage{err}
	}

	defer c.buffer.RPageRelease(rPage.PageTag())

	for tupleIndex := page.TupleIndex(0); tupleIndex < rPage.TupleCount(); tupleIndex++ {
		tuple := rPage.Tuple(tupleIndex)
		columnDef := systable.ColumnsOps.RowToData(tuple)
		relOID := systable.ColumnsOps.TableOID(tuple)

		columns, ok := c.columns[relOID]
		if !ok {
			c.columns[relOID] = []reldef.ColumnDefinition{columnDef}
		} else {
			c.columns[relOID] = append(columns, columnDef)
		}
	}

	return nil
}

func (c *loaderCache) loadSequencePage(pageID page.Id) error {
	rPage, err := c.buffer.RTablePage(systable.SequencesTableDef(), pageID)
	if err != nil {
		return canNotReadPage{err}
	}

	defer c.buffer.RPageRelease(rPage.PageTag())

	for tupleIndex := page.TupleIndex(0); tupleIndex < rPage.TupleCount(); tupleIndex++ {
		tuple := rPage.Tuple(tupleIndex)
		futureSeq := systable.SequencesOps.RowToData(tuple)

		c.sequences[futureSeq.OID()] = futureSeq
	}

	return nil
}

func (c *loaderCache) loadPages(handler func(pageID page.Id) error) error {
	nextPageId := page.Id(0)

	for {
		err := handler(nextPageId)
		if err != nil {
			if errors.Is(err, canNotReadPageErr) {
				break
			} else {
				return err
			}
		}
		nextPageId++
	}
	return nil
}
