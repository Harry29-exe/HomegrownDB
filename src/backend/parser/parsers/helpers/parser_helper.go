package helpers

import (
	"HomegrownDB/backend/parser/parsers/source"
	"HomegrownDB/backend/tokenizer/token"
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

func (h *ParserHelper) NextIs(code token.Code) error {
	_, err := Next(h.source).Has(code).Check()
	return err
}

func (h *ParserHelper) NextSequence(codes ...token.Code) error {
	return NextSequence(h.source, codes...)
}

func (h *ParserHelper) Current() *tokenChecker {
	return Current(h.source)
}

func (h *ParserHelper) CurrentIs(code token.Code) error {
	_, err := Current(h.source).Has(code).Check()
	return err
}

func (h *ParserHelper) CurrentSequence(codes ...token.Code) error {
	return CurrentSequence(h.source, codes...)
}
