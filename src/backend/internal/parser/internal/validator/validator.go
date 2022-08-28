package validator

import (
	"HomegrownDB/backend/internal/parser/internal"
	token2 "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/sqlerr"
)

func NewValidator(source internal.TokenSource) Validator {
	v := new(validator)
	v.source = source
	return v
}

type validator struct {
	source      internal.TokenSource
	current     token2.Token
	currentText *token2.TextToken
}

func (v *validator) Init(source internal.TokenSource) {
	v.source = source
}

func (v *validator) Next() TokenValidator {
	v.current = v.source.Next()
	return v
}

func (v *validator) Current() TokenValidator {
	v.current = v.source.Current()
	return v
}

func (v *validator) NextIs(code token2.Code) error {
	if tk := v.source.Next(); tk == nil {
		return sqlerr.NewSyntaxError(token2.ToString(code), "nil", v.source)
	} else if tk.Code() != code {
		return sqlerr.NewSyntaxError(token2.ToString(code), tk.Value(), v.source)
	}
	return nil
}

func (v *validator) NextIsAnd(code token2.Code) Validator {
	if tk := v.source.Next(); tk == nil {
		return afterErrorValidator{
			token: tk,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), "nil", v.source),
		}
	} else if tk.Code() != code {
		return afterErrorValidator{
			token: tk,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), tk.Value(), v.source),
		}
	}
	return v
}

func (v *validator) CurrentIs(code token2.Code) error {
	if tk := v.source.Current(); tk == nil {
		return sqlerr.NewSyntaxError(token2.ToString(code), "nil", v.source)
	} else if tk.Code() != code {
		return sqlerr.NewSyntaxError(token2.ToString(code), tk.Value(), v.source)
	}
	return nil
}

func (v *validator) CurrentIsAnd(code token2.Code) Validator {
	if tk := v.source.Current(); tk == nil {
		return afterErrorValidator{
			token: tk,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), "nil", v.source),
		}
	} else if tk.Code() != code {
		return afterErrorValidator{
			token: tk,
			err:   sqlerr.NewSyntaxError(token2.ToString(code), tk.Value(), v.source),
		}
	}
	return v
}

func (v *validator) NextSequence(codes ...token2.Code) error {
	v.source.Checkpoint()
	for _, code := range codes {
		next := v.source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), v.source)
			v.source.Rollback()
			return err
		}
	}

	v.source.Commit()
	return nil
}

func (v *validator) NextSequenceAnd(codes ...token2.Code) Validator {
	v.source.Checkpoint()
	for _, code := range codes {
		next := v.source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), v.source)
			v.source.Rollback()
			return afterErrorValidator{
				token: next,
				err:   err,
			}
		}
	}

	v.source.Commit()
	return nil
}

func (v *validator) CurrentSequence(codes ...token2.Code) error {
	currentToken := v.source.Current()
	if currentToken.Code() != codes[0] {
		return sqlerr.NewTokenSyntaxError(codes[0], currentToken.Code(), v.source)
	}

	v.source.Checkpoint()
	for _, code := range codes[1:] {
		next := v.source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), v.source)
			v.source.Rollback()
			return err
		}
	}

	v.source.Commit()
	return nil
}

func (v *validator) CurrentSequenceAnd(codes ...token2.Code) Validator {
	currentToken := v.source.Current()
	if currentToken.Code() != codes[0] {
		return afterErrorValidator{
			token: currentToken,
			err:   sqlerr.NewTokenSyntaxError(codes[0], currentToken.Code(), v.source),
		}
	}

	v.source.Checkpoint()
	for _, code := range codes[1:] {
		next := v.source.Next()
		if next.Code() != code {
			err := sqlerr.NewTokenSyntaxError(code, next.Code(), v.source)
			v.source.Rollback()
			return afterErrorValidator{
				token: next,
				err:   err,
			}
		}
	}

	v.source.Commit()
	return nil
}

func (v *validator) SkipBreaks() *breaksSkipper {
	return SkipBreaks(v.source)
}
