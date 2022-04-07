package helpers

import (
	"HomegrownDB/sql/querry/parser/def"
	"HomegrownDB/sql/querry/parser/sqlerr"
	"HomegrownDB/sql/querry/tokenizer/token"
	"math"
	"strings"
)

func SkipBreaks(source def.TokenSource) *breaksSkipper {
	return &breaksSkipper{
		breakTypes: map[token.Code]*breakType{},
		source:     source,
	}
}

func (h *ParserHelper) SkipBreaks() *breaksSkipper {
	return SkipBreaks(h.source)
}

type breaksSkipper struct {
	breakTypes map[token.Code]*breakType
	source     def.TokenSource
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

	var currentToken token.Token
	if fromCurrent {
		currentToken = b.source.Current()
	} else {
		currentToken = b.source.Next()
	}

	for currentToken != nil && token.IsBreak(currentToken.Code()) {
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
			err := sqlerr.NewSyntaxTextError("expected more of: "+token.ToString(breakType), b.source)
			b.source.Rollback()
			return err
		}
		if data.maxOccurrences < 0 {
			err := sqlerr.NewSyntaxTextError("expected less of: "+token.ToString(breakType), b.source)
			b.source.Rollback()
			return err
		}
	}

	b.source.Commit()
	return nil
}

func (b *breaksSkipper) Type(code token.Code) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: 0,
	}

	return b
}

func (b *breaksSkipper) TypeMin(code token.Code, min int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: math.MaxInt16,
		minOccurrences: min,
	}

	return b
}

func (b *breaksSkipper) TypeMax(code token.Code, max int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: max,
		minOccurrences: 0,
	}

	return b
}

func (b *breaksSkipper) TypeMinMax(code token.Code, min, max int16) *breaksSkipper {
	b.breakTypes[code] = &breakType{
		maxOccurrences: max,
		minOccurrences: min,
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
		builder.WriteString(token.ToString(code))
	}

	return builder.String()
}
