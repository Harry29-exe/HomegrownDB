package parsers

import "HomegrownDB/sql/querry/parser/def"

var Tables tablesParser = tablesParser{}

type tablesParser struct {
}

func (p tablesParser) Parse(source def.TokenSource) (*TablesNode, error) {
	source.Checkpoint()

	//todo implement me
	panic("not implemented")
}

type TablesNode struct {
	Tables []TableNode
}
