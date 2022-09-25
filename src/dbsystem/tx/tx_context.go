package tx

import (
	"HomegrownDB/dbsystem/schema/table"
)

type Ctx struct {
	rLockedTables map[table.Id]bool
	tableStore    table.RDefsStore

	Info         *InfoCtx
	CurrentQuery *QueryCtx
}

func NewContext(txId Id, tableStore table.RDefsStore) *Ctx {
	return &Ctx{
		rLockedTables: make(map[table.Id]bool, 20),
		tableStore:    tableStore,
		CurrentQuery:  NewQueryCtx(),
		Info:          NewInfoCtx(txId),
	}
}

func (c *Ctx) GetTable(name string) (table.Definition, error) {
	def, err := c.tableStore.GetTable(name)
	if err == nil {
		c.rLockedTables[def.TableId()] = true
	}

	return def, err
}

func (c *Ctx) Table(id table.Id) table.Definition {
	c.rLockedTables[id] = true
	return c.tableStore.Table(id)
}
