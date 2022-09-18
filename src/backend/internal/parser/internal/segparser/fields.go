package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/sqlerr"
	token "HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Fields = fieldsParser{}

type fieldsParser struct {
	validator.Validator
}

// Parse starts at Current token and ends at last column name and creates Fields
// from token sequence like following:
//
// "table_alias1.col1, table_alias1.col1, table_alias1.col1"
//
// It does not support whitespace at the start of TokenSource,
// nor it does not touch/parse any chars after
//
// Because of it there are potentials gotchas as following sentences won't
// return error but won't be fully parsed either:
//
// "table1.col1.col2, table2.col1" - will be parsed to second dot and returned
//
// "table1.col1,, table2.col2" - will be parsed to first comma and returned
func (p fieldsParser) Parse(source internal.TokenSource, validator validator.Validator) (pnode.FieldsNode, error) {
	source.Checkpoint()

	parsingToken := source.Current()
	fields := pnode.FieldsNode{Fields: make([]pnode.FieldNode, 0, 5)}

	for {
		if parsingToken.Code() != token.Identifier {
			return fields, sqlerr.NewSyntaxError(token.ToString(token.Identifier), parsingToken.Value(), source)
		}

		field, err := Field.Parse(source, validator)
		if err != nil {
			return fields, err
		}
		fields.AddField(field)

		err = validator.SkipTokens().
			Type(token.SpaceBreak).
			TypeMinMax(token.Comma, 1, 1).
			SkipFromNext()

		if err != nil {
			source.Commit()
			return fields, nil
		}
		source.Next()
	}
}
