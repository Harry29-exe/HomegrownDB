package validator

import (
	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
)

func (v *validator) StartsWithDigit() TextTokenValidator {
	if v.currentText.StartsWithDigit {
		return v
	}
	return afterErrorValidator{
		token: v.currentText,
		err:   sqlerr.NewSyntaxTextError("Expected string starting with digit", v.source),
	}
}

func (v *validator) DontStartWithDigit() TextTokenValidator {
	if !v.currentText.StartsWithDigit {
		return v
	}
	return afterErrorValidator{
		token: v.currentText,
		err:   sqlerr.NewSyntaxTextError("String can not start with digit", v.source),
	}
}

func (v *validator) AsciiOnly() TextTokenValidator {
	if v.currentText.IsAscii {
		return v
	}
	return afterErrorValidator{
		token: v.currentText,
		err:   sqlerr.NewSyntaxTextError("expected string containing only ascii characters", v.source),
	}
}
