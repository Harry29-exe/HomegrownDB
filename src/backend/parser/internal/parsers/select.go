package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Select = selectParser{}

type selectParser struct {
	validator.Validator
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
	err = s.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Identifier)
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
