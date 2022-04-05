package common

import tk "HomegrownDB/sql/querry/tokenizer"

func CheckNextToken(source TokenSource) *tokenChecker {
	return &tokenChecker{
		source: source,
		token:  source.Next(),
		err:    nil,
	}
}

type tokenChecker struct {
	source TokenSource
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
		tc.err = NewSyntaxError(tk.ToString(code), tc.token.Value(), tc.source)
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
		tc.err = NewSyntaxError("Text", tc.token.Value(), tc.source)
		return nilTextTokenChecker(tc)
	}
}

/* ====textTokenChecker==== */

// textTokenChecker is specialization of tokenChecker
// created just to offer checks for tokenizer.TextToken
type textTokenChecker struct {
	*tokenChecker
	textToken *tk.TextToken
}

func nilTextTokenChecker(checker *tokenChecker) *textTokenChecker {
	return &textTokenChecker{
		tokenChecker: checker,
		textToken:    nil,
	}
}

func (ttc *textTokenChecker) StartsWithDigit() *textTokenChecker {
	if ttc.err != nil {
		return ttc
	}
	if !ttc.textToken.StartsWithDigit {
		ttc.err = NewSyntaxTextError("Expected string starting with digit", ttc.source)
	}

	return ttc
}

func (ttc *textTokenChecker) DontStartWithDigit() *textTokenChecker {
	if ttc.err != nil {
		return ttc
	}
	if ttc.textToken.StartsWithDigit {
		ttc.err = NewSyntaxTextError("String can not start with digit", ttc.source)
	}

	return ttc
}

func (ttc *textTokenChecker) AsciiOnly() *textTokenChecker {
	if ttc.err != nil {
		return ttc
	}
	if !ttc.textToken.IsAscii {
		ttc.err = NewSyntaxTextError("expected string containing only ascii characters", ttc.source)
	}

	return ttc
}
