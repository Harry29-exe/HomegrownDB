package helpers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/sqlerr"
	tk "HomegrownDB/sql/querry/tokenizer"
)

func NextToken(source defs.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Next(),
		err:    nil,
	}
}

func CurrentToken(source defs.TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Current(),
		err:    nil,
	}
}

type tokenChecker struct {
	source defs.TokenSource
	token  tk.Token
	err    error
}

func (tc *tokenChecker) Check() (tk.Token, error) {
	if tc.err != nil {
		return nil, tc.err
	}
	return tc.token, nil
}

func (tc *tokenChecker) HasCode(code tk.TokenCode) *tokenChecker {
	if tc.err != nil {
		return tc
	}

	if tc.token.Code() != code {
		tc.err = sqlerr.NewSyntaxError(tk.ToString(code), tc.token.Value(), tc.source)
	}

	return tc
}

func (tc *tokenChecker) IsTextToken() *textTokenChecker {
	tc.HasCode(tk.Text)
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
