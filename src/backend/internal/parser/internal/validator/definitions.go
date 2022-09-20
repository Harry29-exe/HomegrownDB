package validator

import (
	token "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

type Validator interface {
	Next() TokenValidator
	Current() TokenValidator

	NextIs(code token.Code) error
	NextIsAnd(code token.Code) Validator

	CurrentIs(code token.Code) error
	CurrentIsAnd(code token.Code) Validator

	NextSequence(codes ...token.Code) error
	NextSequenceAnd(codes ...token.Code) Validator

	CurrentSequence(codes ...token.Code) error
	CurrentSequenceAnd(codes ...token.Code) Validator

	SkipTokens() TokenSkipper
	// SkipSpaceFromNext is the same as
	// Validator.SkipTokens().TypeMax(token.Space, 1).SkipFromNext)
	SkipSpaceFromNext() error
}

type TokenValidator interface {
	Check() (token.Token, error)
	CheckAnd() Validator
	Has(code token.Code) TokenValidator
	IsTextToken() TextTokenValidator
}

type TextTokenValidator interface {
	TokenValidator
	StartsWithDigit() TextTokenValidator
	DontStartWithDigit() TextTokenValidator
	AsciiOnly() TextTokenValidator
}

//todo delete breakSkiper and implement it on top of validator
type BreaksValidator interface {
}
