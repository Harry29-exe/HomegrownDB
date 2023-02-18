package sqlerr

import (
	"HomegrownDB/backend/internal/parser/tokenizer"
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/dbsystem/hglib"
	"strings"
)

func NewSyntaxError(expected string, actual string, source tokenizer.TokenSource) hglib.DBError {
	current := source.Current()
	if current.Code() == token2.Error {
		errorToken := current.(token2.ErrorToken)
		return tokenizerError{msg: errorToken.Error()}
	}
	return &syntaxError{
		expected:     expected,
		actual:       actual,
		currentQuery: recreateQuery(source),
	}
}

func NewTokenSyntaxError(expected, actual token2.Code, source tokenizer.TokenSource) hglib.DBError {
	current := source.Current()
	if current.Code() == token2.Error {
		errorToken := current.(token2.ErrorToken)
		return tokenizerError{msg: errorToken.Error()}
	}
	return &syntaxError{
		expected:     token2.ToString(expected),
		actual:       token2.ToString(actual),
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

func (s *syntaxError) Area() hglib.Area {
	return hglib.Parser
}

func NewSyntaxTextError(reason string, source tokenizer.TokenSource) hglib.DBError {
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

func (s syntaxTextError) Area() hglib.Area {
	return hglib.Parser
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
