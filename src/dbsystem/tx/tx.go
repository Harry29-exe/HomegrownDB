package tx

type Tx interface {
	TxID() Id
	Level() IsolationLevel
	Status() Status
	CommandExecuted() uint16
}

type StdTx struct {
	Id                Id
	TxLevel           IsolationLevel
	TxStatus          Status
	TxCommandExecuted uint16
}

func (s StdTx) TxID() Id {
	return s.Id
}

func (s StdTx) Level() IsolationLevel {
	return s.TxLevel
}

func (s StdTx) Status() Status {
	return s.TxLevel
}

func (s StdTx) CommandExecuted() uint16 {
	return s.TxCommandExecuted
}

type Id int32
type Counter = uint16

const (
	IdSize             = 4
	CommandCounterSize = 2
)

type IsolationLevel = uint8

const (
	UncommittedRead = iota
	CommittedRead
	RepeatableRead
	SI
	SSI
)

type Status = uint8

const (
	Finished = iota
	InProgress
	Rollback
	NotStarted
)
