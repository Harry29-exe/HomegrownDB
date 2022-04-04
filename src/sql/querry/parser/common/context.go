package common

import (
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
	"strings"
)

// TokenSource array like structure where Next moves pointer one token forward
// and return it while Prev moves pointer one token back and return it.
// If either method reach array end it returns nil without changing pointer
type TokenSource interface {
	Next() tokenizer.Token      // Next move pointer forwards and returns
	Prev() tokenizer.Token      // Prev move pointer backwards and returns
	Current() tokenizer.Token   // Current returns token which has been returned with last method
	History() []tokenizer.Token // History returns all token from beginning to the one that Next would return

	Checkpoint() // Checkpoint creates new checkpoint for parser to rollback
	Commit()     // Commit deletes last checkpoint
	Rollback()   // Rollback to last checkpoint and removes this checkpoint
}

type LesserParser interface {
	Parse(source TokenSource) (parsetree.Node, error)
}

func NewSyntaxError(expected string, actual string, source TokenSource) *syntaxError {
	return &syntaxError{
		expected: expected,
		actual:   actual,
		source:   source,
	}
}

type syntaxError struct {
	expected string
	actual   string
	source   TokenSource
}

func (s *syntaxError) Error() string {
	return "expected: \"" + s.expected + "\" instead got: \"" +
		s.actual + "\"\n" +
		s.recreateQuery() + " <- here "
}

func (s *syntaxError) recreateQuery() string {
	tokens := s.source.History()
	strBuilder := strings.Builder{}
	for _, token := range tokens {
		strBuilder.WriteString(token.Value())
	}

	return strBuilder.String()
}

func NewSyntaxTextError(reason string, source TokenSource) *syntaxTextError {
	return &syntaxTextError{
		reason: reason,
		source: source,
	}
}

type syntaxTextError struct {
	reason string
	source TokenSource
}

func (s *syntaxTextError) Error() string {
	return s.recreateQuery() + " <- " + s.reason
}

//todo move this function to interface or something (it's copied from syntaxError)
func (s *syntaxTextError) recreateQuery() string {
	tokens := s.source.History()
	strBuilder := strings.Builder{}
	for _, token := range tokens {
		strBuilder.WriteString(token.Value())
	}

	return strBuilder.String()
}
