package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/parser/sqlerr"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Fields fieldsParser = fieldsParser{}

type fieldsParser struct{}

// Parse creates ptree.Fields ptree.Node from token sequence like following:
/* table_alias1.col1, table_alias1.col1, table_alias1.col1 */
func (p fieldsParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	source.Checkpoint()
	parsingToken := source.Current()
	fieldsToken := ptree.NewFieldsNode()
	//TODO change NextToken to CurrentToken

	for {
		if parsingToken.Code() != token.Text {
			return nil, sqlerr.NewSyntaxError(token.TextStr, parsingToken.Value(), source)
		}

		field, err := Field.Parse(source)
		if err != nil {
			return nil, err
		}

		err = fieldsToken.AddChild(field)
		if err != nil {
			return nil, err
		}

		parsingToken = source.Next()
		if parsingToken.Code() == token.SpaceBreak {
			_, err := helpers.NextToken(source).
				HasCode(token.Comma).
				Check()
			if err != nil {
				source.Rollback()
				return nil, err
			}
		}
	}
}
