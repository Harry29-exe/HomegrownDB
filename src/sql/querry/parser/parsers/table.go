package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Table tableParser = tableParser{}

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

	err = t.NextSequence(token.SpaceBreak, token.Text)
	if err != nil {
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
