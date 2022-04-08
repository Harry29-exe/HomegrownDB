package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Select selectParser = selectParser{}

type selectParser struct {
	helpers.ParserHelper
}

func (s selectParser) Parse(source source.TokenSource) (*SelectNode, error) {
	source.Checkpoint()
	s.Init(source)

	// Select
	err := s.CurrentIs(token.Select)
	if err != nil {
		return nil, err
	}
	selectNode := SelectNode{}

	err = s.NextIs(token.SpaceBreak)
	if err != nil {
		return nil, err
	}

	// Fields
	source.Next()
	selectNode.Fields, err = Fields.Parse(source)
	if err != nil {
		return nil, err
	}

	err = s.NextIs(token.SpaceBreak)
	if err != nil {
		return nil, err
	}
	//todo create NextSequence
	err = s.NextIs(token.From)

	//// Table
	//err = s.Attach()
	return &selectNode, nil
}

type SelectNode struct {
	Fields *FieldsNode
	Tables *TablesNode
}
