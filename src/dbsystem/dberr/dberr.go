package dberr

type DBError interface {
	error
	Area() Area
	MsgCanBeReturnedToClient() bool
}

type Area = uint8

const (
	Tokenizer Area = iota
	Parser
	Analyser
	Planer
	Executor

	DBSystem
)
