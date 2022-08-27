package validator

import (
	token2 "HomegrownDB/backend/parser/tokenizer/token"
)

type afterErrorValidator struct {
	token token2.Token
	err   error
}

// Validator impl

func (v afterErrorValidator) Next() TokenValidator {
	return v
}

func (v afterErrorValidator) Current() TokenValidator {
	return v
}

func (v afterErrorValidator) NextIs(code token2.Code) error {
	return v.err
}

func (v afterErrorValidator) NextIsAnd(code token2.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentIs(code token2.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentIsAnd(code token2.Code) Validator {
	return v
}

func (v afterErrorValidator) NextSequence(codes ...token2.Code) error {
	return v.err
}

func (v afterErrorValidator) NextSequenceAnd(codes ...token2.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentSequence(codes ...token2.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentSequenceAnd(codes ...token2.Code) Validator {
	return v
}

func (v afterErrorValidator) SkipBreaks() *breaksSkipper {
	return &breaksSkipper{} //todo
}

// TokenValidator impl

func (v afterErrorValidator) Check() (token2.Token, error) {
	return v.token, v.err
}

func (v afterErrorValidator) CheckAnd() Validator {
	return v
}

func (v afterErrorValidator) Has(code token2.Code) TokenValidator {
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
