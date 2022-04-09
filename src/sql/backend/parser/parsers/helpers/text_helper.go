package helpers

import (
	"HomegrownDB/sql/backend/parser/sqlerr"
	tk "HomegrownDB/sql/backend/tokenizer/token"
)

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
		ttc.err = sqlerr.NewSyntaxTextError("Expected string starting with digit", ttc.source)
	}

	return ttc
}

func (ttc *textTokenChecker) DontStartWithDigit() *textTokenChecker {
	if ttc.err != nil {
		return ttc
	}
	if ttc.textToken.StartsWithDigit {
		ttc.err = sqlerr.NewSyntaxTextError("String can not start with digit", ttc.source)
	}

	return ttc
}

func (ttc *textTokenChecker) AsciiOnly() *textTokenChecker {
	if ttc.err != nil {
		return ttc
	}
	if !ttc.textToken.IsAscii {
		ttc.err = sqlerr.NewSyntaxTextError("expected string containing only ascii characters", ttc.source)
	}

	return ttc
}
