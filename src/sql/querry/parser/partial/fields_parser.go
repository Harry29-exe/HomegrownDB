package partial

import (
	"HomegrownDB/sql/querry/parser/common"
	"HomegrownDB/sql/querry/parser/parsetree"
	"HomegrownDB/sql/querry/tokenizer"
)

var FieldParser fieldParser = fieldParser{}

type fieldParser struct{}

func (p fieldParser) Parse(ctx common.TokenSource) (parsetree.Node, error) {
	parsingToken := ctx.GetNextToken()
	for {
		switch parsingToken.Code() {
		case tokenizer.Text:

		}
	}
}
