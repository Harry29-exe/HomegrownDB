package parsers

import (
	"HomegrownDB/sql/querry/parser/def"
	"HomegrownDB/sql/querry/parser/ptree"
)

var Table tableParser = tableParser{}

type tableParser struct {
}

func (t tableParser) Parse(source def.TokenSource) (ptree.Node, error) {
	//TODO implement me
	panic("implement me")
}

type TableNode struct {
	TableName  string
	TableAlias string
}
