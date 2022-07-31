package validator

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/tokenizer/token"
)

type Validator struct {
	source source.TokenSource
}

func (v *Validator) Init(source source.TokenSource) {
	v.source = source
}

func (v *Validator) Next() *tokenValidator {
	return Next(v.source)
}

func (v *Validator) NextIs(code token.Code) error {
	_, err := Next(v.source).Has(code).Check()
	return err
}

func (v *Validator) NextSequence(codes ...token.Code) error {
	return NextSequence(v.source, codes...)
}

func (v *Validator) Current() *tokenValidator {
	return Current(v.source)
}

func (v *Validator) CurrentIs(code token.Code) error {
	_, err := Current(v.source).Has(code).Check()
	return err
}

func (v *Validator) CurrentSequence(codes ...token.Code) error {
	return CurrentSequence(v.source, codes...)
}

func (v *Validator) SkipBreaks() *breaksSkipper {
	return SkipBreaks(v.source)
}
