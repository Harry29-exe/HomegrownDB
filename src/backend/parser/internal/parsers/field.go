package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/parser/pnode"
	"HomegrownDB/backend/tokenizer/token"
)

var Field = fieldParser{}

type fieldParser struct{}

// Parse todo add support for field without table alias
func (f fieldParser) Parse(source source.TokenSource, validator validator.Validator) (pnode.FieldNode, error) {
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
