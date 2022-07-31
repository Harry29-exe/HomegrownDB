package validator

import (
	"HomegrownDB/backend/parser/sqlerr"
	"HomegrownDB/backend/tokenizer/token"
)

func (v *validator) Check() (token.Token, error) {
	return v.current, nil
}

func (v *validator) CheckAnd() Validator {
	return v
}

func (v *validator) Has(code token.Code) TokenValidator {
	if v.current == nil {
		return afterErrorValidator{
			token: nil,
			err:   sqlerr.NewSyntaxError(token.ToString(code), "nil", v.source),
		}
	} else if v.current.Code() != code {
		return afterErrorValidator{
			token: v.current,
			err:   sqlerr.NewSyntaxError(token.ToString(code), v.current.Value(), v.source),
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
	case *token.TextToken:
		v.currentText = textToken
		return v
	default:
		return afterErrorValidator{
			token: v.current,
			err:   sqlerr.NewSyntaxError("Identifier", v.current.Value(), v.source),
		}
	}
}
