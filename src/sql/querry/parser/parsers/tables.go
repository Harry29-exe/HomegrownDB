package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
)

var Tables tablesParser = tablesParser{}

type tablesParser struct {
	helpers.ParserHelper
}

// Parse table declarations which are usually found after FROM keyword
//
func (t tablesParser) Parse(source source.TokenSource) (*TablesNode, error) {
	t.Init(source)
	source.Checkpoint()

	//if {
	//
	//}
	panic("not implemented")
}

type TablesNode struct {
	Tables []TableNode
}
