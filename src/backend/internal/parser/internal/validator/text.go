package validator

import (
	tk "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/sqlerr"
)

// textTokenChecker is specialization of tokenValidator
// created just to offer checks for tokenizer.TextToken
type textTokenChecker struct {
	*tokenValidator
	textToken *tk.TextToken
}

func nilTextTokenChecker(checker *tokenValidator) *textTokenChecker {
	return &textTokenChecker{
		tokenValidator: checker,
		textToken:      nil,
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
