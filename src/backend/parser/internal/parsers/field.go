package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Field = fieldParser{}

type fieldParser struct {
	validator.Validator
}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source source.TokenSource) (*FieldNode, error) {
	f.Init(source)
	source.Checkpoint()

	tableToken, err := f.Current().
		Has(token.Identifier).
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
