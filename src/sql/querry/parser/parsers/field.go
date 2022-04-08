package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Field fieldParser = fieldParser{}

type fieldParser struct {
	helpers.ParserHelper
}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source source.TokenSource) (*FieldNode, error) {
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

	err = f.NextIs(token.Dot)
	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := f.Next().
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
