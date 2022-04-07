package helpers

import (
	"HomegrownDB/sql/querry/parser/def"
	"HomegrownDB/sql/querry/parser/sqlerr"
	tk "HomegrownDB/sql/querry/tokenizer/token"
)

func Next(source def.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Next(),
		err:    nil,
	}
}

func (h *ParserHelper) Next() *tokenChecker {
	return Next(h.source)
}

func Current(source def.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Current(),
		err:    nil,
	}
}

func (h *ParserHelper) Current() *tokenChecker {
	return Current(h.source)
}

func NextIs(source def.TokenSource, code tk.Code) error {
	_, err := Next(source).Has(code).Check()
	return err
}

func (h *ParserHelper) NextIs(code tk.Code) error {
	return NextIs(h.source, code)
}

func CurrentIs(source def.TokenSource, code tk.Code) error {
	_, err := Current(source).Has(code).Check()
	return err
}

func (h *ParserHelper) CurrentIs(code tk.Code) error {
	return CurrentIs(h.source, code)
}

type tokenChecker struct {
	source def.TokenSource
	token  tk.Token
	err    error
}

func (tc *tokenChecker) Check() (tk.Token, error) {
	if tc.err != nil {
		return nil, tc.err
	}
	return tc.token, nil
}

func (tc *tokenChecker) Has(code tk.Code) *tokenChecker {
	switch {
	case tc.err != nil:
		break

	case tc.token == nil:
		tc.err = sqlerr.NewSyntaxError(tk.ToString(code), "nil", tc.source)

	case tc.token.Code() != code:
		tc.err = sqlerr.NewSyntaxError(tk.ToString(code), tc.token.Value(), tc.source)
	}

	return tc
}

func (tc *tokenChecker) IsTextToken() *textTokenChecker {
	tc.Has(tk.Text)
	if tc.err != nil {
		return nilTextTokenChecker(tc)
	}

	switch textToken := tc.token.(type) {
	case *tk.TextToken:
		return &textTokenChecker{
			tokenChecker: tc,
			textToken:    textToken,
		}
	default:
		tc.err = sqlerr.NewSyntaxError("Text", tc.token.Value(), tc.source)
		return nilTextTokenChecker(tc)
	}
}
