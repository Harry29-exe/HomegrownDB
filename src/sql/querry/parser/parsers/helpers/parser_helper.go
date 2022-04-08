package helpers

import (
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/tokenizer/token"
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

func (h *ParserHelper) NextSequence(codes ...token.Code) error {
	return NextSequence(h.source, codes...)
}

func (h *ParserHelper) Current() *tokenChecker {
	return Current(h.source)
}

func (h *ParserHelper) CurrentSequence(codes ...token.Code) error {
	return CurrentSequence(h.source, codes...)
}
