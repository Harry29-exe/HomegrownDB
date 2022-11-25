package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
)

var Table = tableParser{}

type tableParser struct {
}

func (t tableParser) Parse(source internal.TokenSource, validator validator.Validator) (pnode.RangeVar, error) {
	source.Checkpoint()

	name, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return pnode.RangeVar{}, err
	}

	err = validator.NextSequence(token.SpaceBreak, token.Identifier)
	if err != nil {
		source.Commit()
		return pnode.RangeVar{RelName: name.Value(), Alias: name.Value()}, nil
	}

	alias, err := validator.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return pnode.RangeVar{}, err
	}

	source.Commit()
	return pnode.RangeVar{
		RelName: name.Value(),
		Alias:   alias.Value(),
	}, nil
}
