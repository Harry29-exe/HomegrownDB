package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
)

var Table = tableParser{}

type tableParser struct {
}

func (t tableParser) Parse(source internal.TokenSource, validator tkValidator) (pnode.RangeVar, error) {
	source.Checkpoint()

	name, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return nil, err
	}

	err = validator.NextSequence(token.SpaceBreak, token.As)
	if err != nil {
		source.Commit()
		return pnode.NewRangeVar(name.Value(), ""), nil
	}

	alias, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return nil, err
	}

	source.Commit()
	return pnode.NewRangeVar(name.Value(), alias.Value()), nil
}
