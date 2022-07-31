package parser

import (
	"HomegrownDB/backend/sqlparser/internal"
	"HomegrownDB/backend/tokenizer/token"
)

type Parser struct {
	source internal.TokenSource
}

func NewParser(tokenSource internal.TokenSource) *Parser {
	return &Parser{source: tokenSource}
}

func (p *Parser) Next() *tokenChecker {
	return Next(p.source)
}

func (p *Parser) NextIs(code token.Code) error {
	_, err := Next(p.source).Has(code).Check()
	return err
}

func (p *Parser) NextSequence(codes ...token.Code) error {
	return NextSequence(p.source, codes...)
}

func (p *Parser) Current() *tokenChecker {
	return Current(p.source)
}

func (p *Parser) CurrentIs(code token.Code) error {
	_, err := Current(p.source).Has(code).Check()
	return err
}

func (p *Parser) CurrentSequence(codes ...token.Code) error {
	return CurrentSequence(p.source, codes...)
}
