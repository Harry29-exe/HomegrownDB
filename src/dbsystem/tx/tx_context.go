package tx

type Context interface {
	TxId() Id
	CommandExecuted() uint16
}
