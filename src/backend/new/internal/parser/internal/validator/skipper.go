package validator

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"math"
	"strings"
)

func SkipTokens(source internal.TokenSource) TokenSkipper {
	return &tokenSkipper{
		tokenTypes: map[token.Code]*skippingTokenType{},
		source:     source,
	}
}

type TokenSkipper interface {
	SkipFromNext() error
	SkipFromCurrent() error
	Type(code token.Code) TokenSkipper
	TypeMin(code token.Code, min int16) TokenSkipper
	TypeMax(code token.Code, max int16) TokenSkipper
	TypeMinMax(code token.Code, min, max int16) TokenSkipper
	TypeExactly(code token.Code, occurrences int16) TokenSkipper
}

type tokenSkipper struct {
	tokenTypes map[token.Code]*skippingTokenType
	source     internal.TokenSource
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

	var currentToken token.Token
	if fromCurrent {
		currentToken = b.source.Current()
	} else {
		currentToken = b.source.Next()
	}

	for {
		skippingToken, ok := b.tokenTypes[currentToken.Code()]
		if !ok {
			b.source.Prev()
			break
		}
		skippingToken.maxOccurrences--
		skippingToken.minOccurrences--
		currentToken = b.source.Next()
	}

	for tokenType, data := range b.tokenTypes {
		if data.minOccurrences > 0 {
			err := sqlerr.NewSyntaxTextError("expected more of: "+token.ToString(tokenType), b.source)
			b.source.Rollback()
			return err
		}
		if data.maxOccurrences < 0 {
			err := sqlerr.NewSyntaxTextError("expected less of: "+token.ToString(tokenType), b.source)
			b.source.Rollback()
			return err
		}
	}

	b.source.Commit()
	return nil
}

func (b *tokenSkipper) Type(code token.Code) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMin(code token.Code, min int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeMax(code token.Code, max int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMinMax(code token.Code, min, max int16) TokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeExactly(code token.Code, occurrences int16) TokenSkipper {
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
		builder.WriteString(token.ToString(code))
	}

	return builder.String()
}
