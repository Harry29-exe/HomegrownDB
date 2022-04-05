package common

import "strings"

func NewSyntaxError(expected string, actual string, source TokenSource) *syntaxError {
	return &syntaxError{
		expected:     expected,
		actual:       actual,
		currentQuery: recreateQuery(source),
	}
}

type syntaxError struct {
	expected     string
	actual       string
	currentQuery string
}

func (s *syntaxError) Error() string {
	return "expected: \"" + s.expected + "\" instead got: \"" +
		s.actual + "\"\n" +
		s.currentQuery + " <- here "
}

func NewSyntaxTextError(reason string, source TokenSource) *syntaxTextError {
	return &syntaxTextError{
		reason:       reason,
		currentQuery: recreateQuery(source),
	}
}

type syntaxTextError struct {
	reason       string
	currentQuery string
}

func (s *syntaxTextError) Error() string {
	return s.currentQuery + " <- " + s.reason
}

func recreateQuery(source TokenSource) string {
	tokens := source.History()
	strBuilder := strings.Builder{}
	for _, token := range tokens {
		strBuilder.WriteString(token.Value())
	}

	return strBuilder.String()
}
