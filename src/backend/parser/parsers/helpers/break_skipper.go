package helpers

import (
	"HomegrownDB/backend/parser/parsers/source"
	"HomegrownDB/backend/parser/sqlerr"
	token2 "HomegrownDB/backend/tokenizer/token"
	"math"
	"strings"
)

func SkipBreaks(source source.TokenSource) *breaksSkipper {
	return &breaksSkipper{
		breakTypes: map[token2.Code]*breakType{},
		source:     source,
	}
}

func (h *ParserHelper) SkipBreaks() *breaksSkipper {
	return SkipBreaks(h.source)
}

type breaksSkipper struct {
	breakTypes map[token2.Code]*breakType
	source     source.TokenSource
}

type breakType struct {
	maxOccurrences int16
	minOccurrences int16
}

func (b *breaksSkipper) SkipFromNext() error {
	return b.skip(false)
}

func (b *breaksSkipper) SkipFromCurrent() error {
	return b.skip(true)
}

func (b *breaksSkipper) skip(fromCurrent bool) error {
	b.source.Checkpoint()

	var currentToken token2.Token
	if fromCurrent {
		currentToken = b.source.Current()
	} else {
		currentToken = b.source.Next()
	}

	for token2.IsBreak(currentToken.Code()) {
		breakType, ok := b.breakTypes[currentToken.Code()]
		if !ok {
			err := sqlerr.NewSyntaxError(b.breakTypesToStr(), currentToken.Value(), b.source)
			b.source.Rollback()
			return err
		}
		breakType.minOccurrences--
		breakType.maxOccurrences--

		currentToken = b.source.Next()
	}

	for breakType, data := range b.breakTypes {
		if data.minOccurrences > 0 {
			err := sqlerr.NewSyntaxTextError("expected more of: "+token2.ToString(breakType), b.source)
			b.source.Rollback()
			return err
		}
		if data.maxOccurrences < 0 {
			err := sqlerr.NewSyntaxTextError("expected less of: "+token2.ToString(breakType), b.source)
			b.source.Rollback()
			return err
		}
	}

	b.source.Commit()
	return nil
}

func (b *breaksSkipper) Type(code token2.Code) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: 0,
	}

	return b
}

func (b *breaksSkipper) TypeMin(code token2.Code, min int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: min,
	}

	return b
}

func (b *breaksSkipper) TypeMax(code token2.Code, max int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: max,
		minOccurrences: 0,
	}

	return b
}

func (b *breaksSkipper) TypeMinMax(code token2.Code, min, max int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: max,
		minOccurrences: min,
	}

	return b
}

func (b *breaksSkipper) TypeExactly(code token2.Code, occurrences int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: occurrences,
		minOccurrences: occurrences,
	}

	return b
}

func (b *breaksSkipper) breakTypesToStr() string {
	builder := strings.Builder{}

	notFirst := false
	for code := range b.breakTypes {
		if notFirst {
			builder.WriteString("||")
		} else {
			notFirst = true
		}
		builder.WriteString(token2.ToString(code))
	}

	return builder.String()
}
