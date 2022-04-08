package helpers

import (
	"HomegrownDB/sql/querry/parser/parsers/source"
)

type ParserHelper struct {
	source source.TokenSource
}

func (h *ParserHelper) Init(source source.TokenSource) {
	h.source = source
}

func (h *ParserHelper) Next() *tokenChecker {
	return Next(h.source)
}

func (h *ParserHelper) Current() *tokenChecker {
	return Current(h.source)
}
