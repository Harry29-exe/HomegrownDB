package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Field = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source source.TokenSource, validator validator.Validator) (*FieldNode, error) {
	source.Checkpoint()

	tableToken, err := validator.Current().
		Has(token.Identifier).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return nil, err
	}

	columnToken, err := validator.
		NextIsAnd(token.Dot).
		Next().IsTextToken().
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
