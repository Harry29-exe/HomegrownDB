package validator

import (
	"HomegrownDB/backend/internal/parser/internal"
	token2 "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/sqlerr"
	"math"
	"strings"
)

func SkipTokens(source internal.TokenSource) *tokenSkipper {
	return &tokenSkipper{
		tokenTypes: map[token2.Code]*skippingTokenType{},
		source:     source,
	}
}

type tokenSkipper struct {
	tokenTypes map[token2.Code]*skippingTokenType
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

	var currentToken token2.Token
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

func (b *tokenSkipper) Type(code token2.Code) *tokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMin(code token2.Code, min int16) *tokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeMax(code token2.Code, max int16) *tokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: 0,
	}

	return b
}

func (b *tokenSkipper) TypeMinMax(code token2.Code, min, max int16) *tokenSkipper {
	b.tokenTypes[code] = &skippingTokenType{
		maxOccurrences: max,
		minOccurrences: min,
	}

	return b
}

func (b *tokenSkipper) TypeExactly(code token2.Code, occurrences int16) *tokenSkipper {
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
