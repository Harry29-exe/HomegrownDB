package tx

import "HomegrownDB/dbsystem/schema/table"

type Ctx interface {
	TxId() Id
	CommandExecuted() uint16
	IncrementCommandCounter()
}

func NewContext(txId Id) Ctx {
	return &context{
		txId:            txId,
		commandExecuted: 0,
	}
}

type context struct {
	txId            Id
	commandExecuted uint16
	rLockedTables   []table.Id
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

func (c *context) QueryWillAccessTable() {

}
