package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/parser/sqlerr"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Fields fieldsParser = fieldsParser{}

type fieldsParser struct {
	helpers.ParserHelper
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
func (p fieldsParser) Parse(source source.TokenSource) (*FieldsNode, error) {
	p.Init(source)
	source.Checkpoint()

	parsingToken := source.Current()
	fields := FieldsNode{Fields: make([]*FieldNode, 0, 5)}

	for {
		if parsingToken.Code() != token.Text {
			return nil, sqlerr.NewSyntaxError(token.TextStr, parsingToken.Value(), source)
		}

		field, err := Field.Parse(source)
		if err != nil {
			return nil, err
		}
		fields.AddField(field)

		err = p.SkipBreaks().
			Type(token.SpaceBreak).
			TypeMinMax(token.Comma, 1, 1).
			SkipFromNext()

		if err != nil {
			source.Commit()
			return &fields, nil
		}
	}
}

type FieldsNode struct {
	Fields []*FieldNode
}

func (f *FieldsNode) AddField(field *FieldNode) {
	f.Fields = append(f.Fields, field)
}
