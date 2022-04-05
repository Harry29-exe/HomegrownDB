package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer"
)

var Field fieldParser = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	source.Checkpoint()

	tableToken, err := helpers.CurrentToken(source).
		HasCode(tokenizer.Text).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	_, err = helpers.NextToken(source).
		HasCode(tokenizer.Dot).
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := helpers.NextToken(source).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	source.Commit()
	return ptree.NewFieldNode(ptree.FieldNodeValue{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: "",
	}), nil
}
