package validator

import (
	token2 "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
)

type Validator interface {
	Next() TokenValidator
	Current() TokenValidator

	NextIs(code token2.Code) error
	NextIsAnd(code token2.Code) Validator

	CurrentIs(code token2.Code) error
	CurrentIsAnd(code token2.Code) Validator

	NextSequence(codes ...token2.Code) error
	NextSequenceAnd(codes ...token2.Code) Validator

	CurrentSequence(codes ...token2.Code) error
	CurrentSequenceAnd(codes ...token2.Code) Validator

	SkipBreaks() *breaksSkipper
}

type TokenValidator interface {
	Check() (token2.Token, error)
	CheckAnd() Validator
	Has(code token2.Code) TokenValidator
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
