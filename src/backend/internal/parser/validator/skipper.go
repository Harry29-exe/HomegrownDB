package validator

import (
	"HomegrownDB/backend/internal/parser/tokenizer"
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/sqlerr"
	"math"
	"strings"
)

func SkipTokens(source tokenizer.TokenSource) TokenSkipper {
	return &tokenSkipper{
		tokenTypes: map[token2.Code]*skippingTokenType{},
		source:     source,
	}
}

type TokenSkipper interface {
	SkipFromNext() error
	SkipFromCurrent() error
	Type(code token2.Code) TokenSkipper
	TypeMin(code token2.Code, min int16) TokenSkipper
	TypeMax(code token2.Code, max int16) TokenSkipper
	TypeMinMax(code token2.Code, min, max int16) TokenSkipper
	TypeExactly(code token2.Code, occurrences int16) TokenSkipper
}

type tokenSkipper struct {
	tokenTypes map[token2.Code]*skippingTokenType
	source     tokenizer.TokenSource
}

type skippingTokenType struct {
	maxOccurrences int16
	minOccurrences int16
}

func (b *tokenSkipper) SkipFromNext() error {
	return b.skip(false)
}

func (b *tokenSkipper) SkipFromCurrent() error {
	return b.skip(true)
}

func (b *tokenSkipper) skip(fromCurrent bool) error {
	b.source.Checkpoint()

	var currentToken token2.Token
	if fromCurrent {
		currentToken = b.source.Current()
	} else {
		currentToken = b.source.Next()
	}

	for {
		skippingToken, ok := b.tokenTypes[currentToken.Code()]
		if !ok {
			break
		}
		skippingToken.maxOccurrences--
		skippingToken.minOccurrences--
		currentToken = b.source.Next()
	}

	for tokenType, data := range b.tokenTypes {
		if data.minOccurrences > 0 {
			err := sqlerr.NewSyntaxTextError("expected more of: "+token2.ToString(tokenType), b.source)
			b.source.Rollback()
			return err
		}
		if data.maxOccurrences < 0 {
			err := sqlerr.NewSyntaxTextError("expected less of: "+token2.ToString(tokenType), b.source)
			b.source.Rollback()
			return err
		}
	}

	b.source.Commit()
	return nil
}

func (b *tokenSkipper) Type(code token2.Code) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMin(code token2.Code, min int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeMax(code token2.Code, max int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMinMax(code token2.Code, min, max int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeExactly(code token2.Code, occurrences int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: occurrences,
		minOccurrences: occurrences,
	}

	return b
}

func (b *tokenSkipper) breakTypesToStr() string {
	builder := strings.Builder{}

	notFirst := false
	for code := range b.tokenTypes {
		if notFirst {
			builder.WriteString("||")
		} else {
			notFirst = true
		}
		builder.WriteString(token2.ToString(code))
	}

	return builder.String()
}
