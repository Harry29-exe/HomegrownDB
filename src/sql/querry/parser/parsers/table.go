package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/ptree"
)

var Table tableParser = tableParser{}

type tableParser struct {
	defs.ParserPrototype
}

func (t tableParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	//TODO implement me
	panic("implement me")
}
