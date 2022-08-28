package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Field = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source internal.TokenSource, validator validator.Validator) (pnode.FieldNode, error) {
	source.Checkpoint()

	tableToken, err := validator.Current().
		Has(token.Identifier).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return pnode.FieldNode{}, err
	}

	columnToken, err := validator.
		NextIsAnd(token.Dot).
		Next().IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return pnode.FieldNode{}, err
	}

	source.Commit()
	return pnode.FieldNode{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: columnToken.Value(),
	}, nil
}
