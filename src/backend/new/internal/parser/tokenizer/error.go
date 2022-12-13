package tokenizer

import "fmt"

func eofError(whileReadingTokenType, query string) Error {
	return Error{
		msg: fmt.Sprintf(
			"Enexpected end of query while reading %s.\n%s <-- here",
			whileReadingTokenType,
			query),
	}
}

type Error struct {
	msg string
}

func (e Error) Error() string {
	return e.msg
}
