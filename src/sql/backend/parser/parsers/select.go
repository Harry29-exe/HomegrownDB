package parsers

import (
	"HomegrownDB/sql/backend/parser/parsers/helpers"
	"HomegrownDB/sql/backend/parser/parsers/source"
	"HomegrownDB/sql/backend/tokenizer/token"
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

	// From
	err = s.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Text)
	if err != nil {
		return nil, err
	}

	// Tables
	tables, err := Tables.Parse(source)
	if err != nil {
		return nil, err
	}
	selectNode.Tables = tables

	source.Commit()
	return &selectNode, nil
}

type SelectNode struct {
	Fields *FieldsNode
	Tables *TablesNode
}
