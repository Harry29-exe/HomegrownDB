package parsers

import (
	"HomegrownDB/sql/querry/parser/defs"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/parser/ptree"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Select selectParser = selectParser{}

type selectParser struct {
	defs.ParserPrototype
}

func (s selectParser) Parse(source defs.TokenSource) (ptree.Node, error) {
	source.Checkpoint()

	// Select
	_, err := helpers.CurrentToken(source).
		HasCode(token.Select).
		Check()
	if err != nil {
		return nil, err
	}
	s.Root = ptree.NewSelectNode()

	err = helpers.SkipBreaks(source).
		TypeMax(token.SpaceBreak, 2).
		TypeMinMax(token.Comma, 1, 1).
		ShipFromNext()
	if err != nil {
		return nil, err
	}

	// Fields
	err = s.Attach(Fields.Parse(source))
	if err != nil {
		return nil, err
	}

	// From
	err = s.Attach()
}
