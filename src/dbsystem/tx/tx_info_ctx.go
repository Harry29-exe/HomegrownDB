package tx

func NewInfoCtx(txId Id) *InfoCtx {
	return &InfoCtx{
		txId:            txId,
		commandExecuted: 0,
	}
}

type InfoCtx struct {
	txId            Id
	commandExecuted uint16
}

func (i *InfoCtx) TxId() Id {
	return i.txId
}

func (i *InfoCtx) CommandExecuted() uint16 {
	return i.commandExecuted
}

func (i *InfoCtx) IncrementCommandCounter() {
	i.commandExecuted++
}
