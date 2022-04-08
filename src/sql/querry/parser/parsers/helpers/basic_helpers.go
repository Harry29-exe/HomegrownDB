package helpers

import (
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/parser/sqlerr"
	"HomegrownDB/sql/querry/tokenizer/token"
)

func Next(source source.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Next(),
		err:    nil,
	}
}

func NextSequence(source source.TokenSource, codes ...token.Code) error {
	for _, code := range codes {
		next := source.Next()
		if next.Code() != code {
			return sqlerr.NewTokenSyntaxError(code, next.Code(), source)
		}
	}

	return nil
}

func Current(source source.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Current(),
		err:    nil,
	}
}

func CurrentSequence(source source.TokenSource, codes ...token.Code) error {
	currentToken := source.Current()
	if currentToken.Code() != codes[0] {
		return sqlerr.NewTokenSyntaxError(codes[0], currentToken.Code(), source)
	}

	for _, code := range codes[1:] {
		next := source.Next()
		if next.Code() != code {
			return sqlerr.NewTokenSyntaxError(code, next.Code(), source)
		}
	}

	return nil
}

func NextIs(source source.TokenSource, code token.Code) error {
	_, err := Next(source).Has(code).Check()
	return err
}

func (h *ParserHelper) NextIs(code token.Code) error {
	return NextIs(h.source, code)
}

func CurrentIs(source source.TokenSource, code token.Code) error {
	_, err := Current(source).Has(code).Check()
	return err
}

func (h *ParserHelper) CurrentIs(code token.Code) error {
	return CurrentIs(h.source, code)
}

type tokenChecker struct {
	source source.TokenSource
	token  token.Token
	err    error
}

func (tc *tokenChecker) Check() (token.Token, error) {
	if tc.err != nil {
		return nil, tc.err
	}
	return tc.token, nil
}

func (tc *tokenChecker) Has(code token.Code) *tokenChecker {
	switch {
	case tc.err != nil:
		break

	case tc.token == nil:
		tc.err = sqlerr.NewSyntaxError(token.ToString(code), "nil", tc.source)

	case tc.token.Code() != code:
		tc.err = sqlerr.NewSyntaxError(token.ToString(code), tc.token.Value(), tc.source)
	}

	return tc
}

func (tc *tokenChecker) IsTextToken() *textTokenChecker {
	tc.Has(token.Text)
	if tc.err != nil {
		return nilTextTokenChecker(tc)
	}

	switch textToken := tc.token.(type) {
	case *token.TextToken:
		return &textTokenChecker{
			tokenChecker: tc,
			textToken:    textToken,
		}
	default:
		tc.err = sqlerr.NewSyntaxError("Text", tc.token.Value(), tc.source)
		return nilTextTokenChecker(tc)
	}
}
