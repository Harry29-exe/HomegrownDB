package sqlerr

type SqlError interface {
	error
	Level() Level
}

type Level = uint8

const (
	Tokenizer Level = iota
	Parser
	Analyser
	Planer
	Executor
)
