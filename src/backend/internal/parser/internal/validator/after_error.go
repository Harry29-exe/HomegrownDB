package validator

import (
	tk "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

type afterErrorValidator struct {
	token tk.Token
	err   error
}

// Validator impl

func (v afterErrorValidator) Next() TokenValidator {
	return v
}

func (v afterErrorValidator) Current() TokenValidator {
	return v
}

func (v afterErrorValidator) NextIs(code tk.Code) error {
	return v.err
}

func (v afterErrorValidator) NextIsAnd(code tk.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentIs(code tk.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentIsAnd(code tk.Code) Validator {
	return v
}

func (v afterErrorValidator) NextSequence(codes ...tk.Code) error {
	return v.err
}

func (v afterErrorValidator) NextSequenceAnd(codes ...tk.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentSequence(codes ...tk.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentSequenceAnd(codes ...tk.Code) Validator {
	return v
}

func (v afterErrorValidator) SkipTokens() *tokenSkipper {
	return &tokenSkipper{} //todo
}

// TokenValidator impl

func (v afterErrorValidator) Check() (tk.Token, error) {
	return v.token, v.err
}

func (v afterErrorValidator) CheckAnd() Validator {
	return v
}

func (v afterErrorValidator) Has(code tk.Code) TokenValidator {
	return v
}

func (v afterErrorValidator) IsTextToken() TextTokenValidator {
	return v
}

// TextTokenValidator impl

func (v afterErrorValidator) StartsWithDigit() TextTokenValidator {
	return v
}

func (v afterErrorValidator) DontStartWithDigit() TextTokenValidator {
	return v
}

func (v afterErrorValidator) AsciiOnly() TextTokenValidator {
	return v
}
