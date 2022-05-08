package parsers

import (
	"HomegrownDB/backend/parser/parsers/helpers"
	"HomegrownDB/backend/parser/parsers/source"
	"HomegrownDB/backend/tokenizer/token"
)

var Table = tableParser{}

type tableParser struct {
	helpers.ParserHelper
}

func (t tableParser) Parse(source source.TokenSource) (*TableNode, error) {
	t.Init(source)
	source.Checkpoint()

	name, err := t.Current().
		IsTextToken().
		AsciiOnly().
		DontStartWithDigit().
		Check()
	if err != nil {
		source.Rollback()
		return nil, err
	}

	err = t.NextSequence(token.SpaceBreak, token.Identifier)
	if err != nil {
		source.Commit()
		return &TableNode{name.Value(), name.Value()}, nil
	}

	alias, err := t.Current().
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
