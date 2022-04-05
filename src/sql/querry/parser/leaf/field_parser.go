package leaf

import (
	"HomegrownDB/sql/querry/parser/common"
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
)

var FieldParser fieldParser = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source common.TokenSource) (parsetree.Node, error) {
	source.Checkpoint()

	tableToken, err := common.CheckNextToken(source).
		HasCode(tokenizer.Text).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	_, err = common.CheckNextToken(source).
		HasCode(tokenizer.Dot).
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := common.CheckNextToken(source).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	source.Commit()
	return parsetree.NewFieldNode(parsetree.FieldNodeValue{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: "",
	}), nil
}
