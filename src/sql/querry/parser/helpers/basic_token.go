package helpers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/sqlerr"
	token2 "HomegrownDB/sql/querry/tokenizer/token"
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
	token  token2.Token
	err    error
}

func (tc *tokenChecker) Check() (token2.Token, error) {
	if tc.err != nil {
		return nil, tc.err
	}
	return tc.token, nil
}

func (tc *tokenChecker) HasCode(code token2.Code) *tokenChecker {
	if tc.err != nil {
		return tc
	}

	if tc.token.Code() != code {
		tc.err = sqlerr.NewSyntaxError(token2.ToString(code), tc.token.Value(), tc.source)
	}

	return tc
}

func (tc *tokenChecker) IsTextToken() *textTokenChecker {
	tc.HasCode(token2.Text)
	if tc.err != nil {
		return nilTextTokenChecker(tc)
	}

	switch textToken := tc.token.(type) {
	case *token2.TextToken:
		return &textTokenChecker{
			tokenChecker: tc,
			textToken:    textToken,
		}
	default:
		tc.err = sqlerr.NewSyntaxError("Text", tc.token.Value(), tc.source)
		return nilTextTokenChecker(tc)
	}
}
