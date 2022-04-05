package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer"
)

var Select selectParser = selectParser{}

type selectParser struct {
	defs.ParserPrototype
}

func (s selectParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	source.Checkpoint()

	// Select
	_, err := helpers.CurrentToken(source).
		HasCode(tokenizer.Select).
		Check()
	if err != nil {
		return nil, err
	}
	s.Root = ptree.NewSelectNode()

	// Fields
	err = s.Attach(Fields.Parse(source))
	if err != nil {
		return nil, err
	}

	// From
	err = s.Attach()
}
