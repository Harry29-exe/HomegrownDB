package parsers

import (
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Field fieldParser = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source TokenSource) (*FieldNode, error) {
	source.Checkpoint()

	tableToken, err := helpers.CurrentToken(source).
		HasCode(token.Text).
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
		ShipFromNext()
	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := helpers.CurrentToken(source).
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
		FieldAlias: "",
	}, nil
}

type FieldNode struct {
	TableAlias string
	FieldName  string
	FieldAlias string
}
