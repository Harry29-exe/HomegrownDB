package validator

import (
	tk "HomegrownDB/backend/new/internal/parser/tokenizer/token"
)

// This struct implements all validator so when validator encounters
// error it can return this struct that has dummy methods and at the
// end will return encountered error
//
// #validator.Validator##validator.TokenValidator#
// #validator.TextTokenValidator##validator.TokenSkipper#
type afterErrorValidator struct {
	token tk.Token
	err   error
}

var _ Validator = afterErrorValidator{}

var _ TokenValidator = afterErrorValidator{}

var _ TextTokenValidator = afterErrorValidator{}

func (v afterErrorValidator) Next() TokenValidator {
	return v
}

func (v afterErrorValidator) Current() TokenValidator {
	return v
}

// Validator impl

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

func (v afterErrorValidator) SkipNextSB() error {
	return v.err
}

func (v afterErrorValidator) SkipNextSBAnd() Validator {
	return v
}

func (v afterErrorValidator) SkipCurrentSB() error {
	return v.err
}

func (v afterErrorValidator) SkipCurrentSBAnd() Validator {
	return v
}

func (v afterErrorValidator) SkipTokens() TokenSkipper {
	return v
}

func (v afterErrorValidator) SkipSpaceFromNext() error {
	return v.err
}

func (v afterErrorValidator) Check() (tk.Token, error) {
	return v.token, v.err
}

// TokenValidator impl

func (v afterErrorValidator) CheckAnd() Validator {
	return v
}

func (v afterErrorValidator) Has(code tk.Code) TokenValidator {
	return v
}

func (v afterErrorValidator) IsTextToken() TextTokenValidator {
	return v
}

func (v afterErrorValidator) StartsWithDigit() TextTokenValidator {
	return v
}

// TextTokenValidator impl

func (v afterErrorValidator) DontStartWithDigit() TextTokenValidator {
	return v
}

func (v afterErrorValidator) AsciiOnly() TextTokenValidator {
	return v
}

// TokenSkipper

func (v afterErrorValidator) SkipFromNext() error {
	return v.err
}

func (v afterErrorValidator) SkipFromCurrent() error {
	return v.err
}

func (v afterErrorValidator) Type(code tk.Code) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMin(code tk.Code, min int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMax(code tk.Code, max int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMinMax(code tk.Code, min, max int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeExactly(code tk.Code, occurrences int16) TokenSkipper {
	return v
}
