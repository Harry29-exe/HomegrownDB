package sqlerr

import (
	"HomegrownDB/dbsystem/hglib"
)

func (s *syntaxError) Error() string {
	return "expected: \"" + s.expected + "\" instead got: \"" +
		s.actual + "\"\n" +
		s.currentQuery + " <- here "
}

type tokenizerError struct {
	msg string
}

func (t tokenizerError) Error() string {
	return t.msg
}

func (t tokenizerError) MsgCanBeReturnedToClient() bool {
	return true
}

func (t tokenizerError) Area() hglib.Area {
	return hglib.Tokenizer
}
