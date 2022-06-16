package tx

type Context interface {
	TxId() Id
	CommandExecuted() uint16
	IncrementCommandCounter()
}

func NewContext(txId Id) Context {
	return &context{
		txId:            txId,
		commandExecuted: 0,
	}
}

type context struct {
	txId            Id
	commandExecuted uint16
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
