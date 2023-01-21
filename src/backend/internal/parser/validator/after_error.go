package validator

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
)

// This struct implements all validator so when validator encounters
// error it can return this struct that has dummy methods and at the
// end will return encountered error
//
// #validator.Validator##validator.TokenValidator#
// #validator.TextTokenValidator##validator.TokenSkipper#
type afterErrorValidator struct {
	token token.Token
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

func (v afterErrorValidator) NextIs(code token.Code) error {
	return v.err
}

func (v afterErrorValidator) NextIsAnd(code token.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentIs(code token.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentIsAnd(code token.Code) Validator {
	return v
}

func (v afterErrorValidator) NextSequence(codes ...token.Code) error {
	return v.err
}

func (v afterErrorValidator) NextSequenceAnd(codes ...token.Code) Validator {
	return v
}

func (v afterErrorValidator) CurrentSequence(codes ...token.Code) error {
	return v.err
}

func (v afterErrorValidator) CurrentSequenceAnd(codes ...token.Code) Validator {
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

func (v afterErrorValidator) Check() (token.Token, error) {
	return v.token, v.err
}

// TokenValidator impl

func (v afterErrorValidator) CheckAnd() Validator {
	return v
}

func (v afterErrorValidator) Has(code token.Code) TokenValidator {
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

func (v afterErrorValidator) Type(code token.Code) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMin(code token.Code, min int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMax(code token.Code, max int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeMinMax(code token.Code, min, max int16) TokenSkipper {
	return v
}

func (v afterErrorValidator) TypeExactly(code token.Code, occurrences int16) TokenSkipper {
	return v
}
