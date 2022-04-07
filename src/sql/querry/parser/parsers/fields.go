package parsers

import (
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/sqlerr"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Fields fieldsParser = fieldsParser{}

type fieldsParser struct{}

// Parse creates ptree.Fields ptree.Node from token sequence like following:
/* table_alias1.col1, table_alias1.col1, table_alias1.col1 */
func (p fieldsParser) Parse(source TokenSource) (*FieldsNode, error) {
	source.Checkpoint()
	parsingToken := source.Current()
	fields := FieldsNode{Fields: make([]*FieldNode, 0, 5)}
	//TODO change NextToken to CurrentToken

	for {
		if parsingToken.Code() != token.Text {
			return nil, sqlerr.NewSyntaxError(token.TextStr, parsingToken.Value(), source)
		}

		field, err := Field.Parse(source)
		if err != nil {
			return nil, err
		}
		fields.AddField(field)

		err = helpers.SkipBreaks(source).
			Type(token.SpaceBreak).
			TypeMinMax(token.Comma, 1, 1).
			SkipFromCurrent()
		if err != nil {
			return nil, err
		}
	}
}

type FieldsNode struct {
	Fields []*FieldNode
}

func (f *FieldsNode) AddField(field *FieldNode) {
	f.Fields = append(f.Fields, field)
}
