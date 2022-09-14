package dberr

type DBError interface {
	error
	Area() Area
}

type Area = uint8

const (
	Tokenizer Area = iota
	Parser
	Analyser
	Planer
	Executor
)
