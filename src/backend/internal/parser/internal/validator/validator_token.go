package validator

import (
	"HomegrownDB/backend/internal/parser/internal/sqlerr"
	token2 "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

func (v *validator) Check() (token2.Token, error) {
	return v.current, nil
}

func (v *validator) CheckAnd() Validator {
	return v
}

func (v *validator) Has(code token2.Code) TokenValidator {
	if v.current == nil {
		return afterErrorValidator{
			token: nil,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), "nil", v.source),
		}
	} else if v.current.Code() != code {
		return afterErrorValidator{
			token: v.current,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), v.current.Value(), v.source),
		}
	}
	return v
}

func (v *validator) IsTextToken() TextTokenValidator {
	if v.current == nil {
		return afterErrorValidator{
			token: nil,
			err:   sqlerr.NewSyntaxError("Identifier", "nil", v.source),
		}
	}

	switch textToken := v.current.(type) {
	case *token2.TextToken:
		v.currentText = textToken
		return v
	default:
		return afterErrorValidator{
			token: v.current,
			err:   sqlerr.NewSyntaxError("Identifier", v.current.Value(), v.source),
		}
	}
}
