package tx

type Tx struct {
	Id     Id
	Level  IsolationLevel
	Status Status
}

type Id = uint32

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
