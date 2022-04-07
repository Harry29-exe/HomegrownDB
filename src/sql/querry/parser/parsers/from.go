package parsers

import (
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Tables tablesParser = tablesParser{}

type tablesParser struct {
}

func (p tablesParser) Parse(source TokenSource) (*TablesNode, error) {
	source.Checkpoint()

	_, err := helpers.CurrentToken(source).
		HasCode(token.From).
		Check()
	if err != nil {
		return nil, err
	}
	p.Root = ptree.NewFromNode()

}

type TablesNode struct {
	Tables []TableNode
}
