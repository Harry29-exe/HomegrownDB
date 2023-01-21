package segparser

import (
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
)

var RangeVar = rangeVar{}

type rangeVar struct {
}

func (t rangeVar) Parse(source tkSource, validator tkValidator) (pnode.RangeVar, error) {
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

	err = validator.NextSequence(token.SpaceBreak, token.Identifier)
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
