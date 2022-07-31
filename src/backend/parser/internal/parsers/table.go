package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Table = tableParser{}

type tableParser struct {
}

func (t tableParser) Parse(source source.TokenSource, validator validator.Validator) (*TableNode, error) {
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
		return &TableNode{name.Value(), name.Value()}, nil
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
	return &TableNode{
		TableName:  name.Value(),
		TableAlias: alias.Value(),
	}, nil
}

type TableNode struct {
	TableName  string
	TableAlias string
}