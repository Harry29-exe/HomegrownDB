package tx

import (
	"HomegrownDB/dbsystem/schema/table"
)

type Ctx interface {
	TxId() Id
	CommandExecuted() uint16
	IncrementCommandCounter()

	CurrentQuery() *QueryCtx
	table.RDefsStore
}

func NewContext(txId Id, tableStore table.RDefsStore) Ctx {
	return &context{
		txId:            txId,
		commandExecuted: 0,
		rLockedTables:   make(map[table.Id]bool, 20),
		tableStore:      tableStore,
	}
}

type context struct {
	txId            Id
	commandExecuted uint16
	rLockedTables   map[table.Id]bool
	tableStore      table.RDefsStore

	currentQuery *QueryCtx
}

func (c *context) TxId() Id {
	return c.txId
}

func (c *context) CommandExecuted() uint16 {
	return c.commandExecuted
}

func (c *context) IncrementCommandCounter() {
	c.commandExecuted++
}

func (c *context) CurrentQuery() *QueryCtx {
	return c.currentQuery
}

func (c *context) GetTable(name string) (table.Definition, error) {
	def, err := c.tableStore.GetTable(name)
	if err == nil {
		c.rLockedTables[def.TableId()] = true
	}

	return def, err
}

func (c *context) Table(id table.Id) table.Definition {
	c.rLockedTables[id] = true
	return c.tableStore.Table(id)
}
