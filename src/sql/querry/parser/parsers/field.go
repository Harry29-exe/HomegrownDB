package parsers

import (
	"HomegrownDB/sql/querry/parser/def"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Field fieldParser = fieldParser{}

type fieldParser struct {
	helpers.ParserHelper
}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source def.TokenSource) (*FieldNode, error) {
	f.Init(source)
	source.Checkpoint()

	tableToken, err := f.Current().
		Has(token.Text).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	err = helpers.SkipBreaks(source).
		Type(token.SpaceBreak).
		TypeMinMax(token.Dot, 1, 1).
		SkipFromNext()
	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := helpers.Current(source).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	source.Commit()
	return &FieldNode{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: columnToken.Value(),
	}, nil
}

type FieldNode struct {
	TableAlias string
	FieldName  string
	FieldAlias string
}
