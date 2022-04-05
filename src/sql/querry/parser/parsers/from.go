package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer"
)

var From fromParser = fromParser{}

type fromParser struct {
	defs.ParserPrototype
}

func (p fromParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	source.Checkpoint()

	_, err := helpers.CurrentToken(source).
		HasCode(tokenizer.From).
		Check()
	if err != nil {
		return nil, err
	}
	p.Root = ptree.NewFromNode()

}
