package creator

import (
	"HomegrownDB/dbsystem/access/relation/dbobj"
	"HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
)

func newPageCache(fs dbfs.FS, tables ...table.RDefinition) *pageCache {
	cache := pageCache{
		fs:     fs,
		tables: map[dbobj.OID]table.RDefinition{},
		cache:  map[dbobj.OID][]page.WPage{},
	}
	for _, table := range tables {
		cache.tables[table.OID()] = table
		cache.cache[table.OID()] = []page.WPage{}
	}
	return &cache
}

type pageCache struct {
	fs     dbfs.FS
	tables map[dbobj.OID]table.RDefinition
	cache  map[dbobj.OID][]page.WPage
}

func (c *pageCache) insert(oid dbobj.OID, tuple page.WTuple) error {
	tupleBytes := tuple.Bytes()

	pages := c.cache[oid]
	for _, wPage := range pages {
		if int(wPage.FreeSpace()) > len(tuple.Bytes()) {
			return wPage.InsertTuple(tupleBytes)
		}
	}
	newPage := page.AsTablePage(
		make([]byte, page.Size),
		page.Id(len(pages)),
		c.tables[oid])
	if err := newPage.InsertTuple(tupleBytes); err != nil {
		return err
	}

	c.cache[oid] = append(pages, newPage)
	return nil
}

func (c *pageCache) flush() error {
	for oid, pages := range c.cache {
		dataFile, err := c.fs.OpenPageObjectFile(oid)
		if err != nil {
			return err
		}
		for _, wPage := range pages {
			_, err = dataFile.Write(wPage.Bytes())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
