package sqlerr

import (
	"HomegrownDB/backend/new/internal/parser/tokenizer"
	"HomegrownDB/backend/new/internal/parser/tokenizer/token"
	"HomegrownDB/dbsystem/dberr"
	"strings"
)

func NewSyntaxError(expected string, actual string, source tokenizer.TokenSource) dberr.DBError {
	current := source.Current()
	if current.Code() == token.Error {
		errorToken := current.(token.ErrorToken)
		return tokenizerError{msg: errorToken.Error()}
	}
	return &syntaxError{
		expected:     expected,
		actual:       actual,
		currentQuery: recreateQuery(source),
	}
}

func NewTokenSyntaxError(expected, actual token.Code, source tokenizer.TokenSource) dberr.DBError {
	current := source.Current()
	if current.Code() == token.Error {
		errorToken := current.(token.ErrorToken)
		return tokenizerError{msg: errorToken.Error()}
	}
	return &syntaxError{
		expected:     token.ToString(expected),
		actual:       token.ToString(actual),
		currentQuery: recreateQuery(source),
	}
}

type syntaxError struct {
	expected     string
	actual       string
	currentQuery string
}

func (s *syntaxError) MsgCanBeReturnedToClient() bool {
	return true
}

func (s *syntaxError) Area() dberr.Area {
	return dberr.Parser
}

func NewSyntaxTextError(reason string, source tokenizer.TokenSource) dberr.DBError {
	return syntaxTextError{
		reason:       reason,
		currentQuery: recreateQuery(source),
	}
}

type syntaxTextError struct {
	reason       string
	currentQuery string
}

func (s syntaxTextError) Error() string {
	return s.currentQuery + " <- " + s.reason
}

func (s syntaxTextError) Area() dberr.Area {
	return dberr.Parser
}

func (s syntaxTextError) MsgCanBeReturnedToClient() bool {
	return true
}

func recreateQuery(source tokenizer.TokenSource) string {
	tokens := source.History()
	strBuilder := strings.Builder{}
	for _, tk := range tokens {
		strBuilder.WriteString(tk.Value())
	}

	return strBuilder.String()
}
