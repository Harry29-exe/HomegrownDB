package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Select = selectParser{}

type selectParser struct{}

func (s selectParser) Parse(source source.TokenSource) (*SelectNode, error) {
	source.Checkpoint()
	v := validator.NewValidator(source)

	// Select
	err := v.CurrentIsAnd(token.Select).
		NextIs(token.SpaceBreak)
	if err != nil {
		return nil, err
	}
	selectNode := SelectNode{}

	// Fields
	source.Next()
	selectNode.Fields, err = Fields.Parse(source, v)
	if err != nil {
		return nil, err
	}

	// From
	err = v.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Identifier)
	if err != nil {
		return nil, err
	}

	// Tables
	tables, err := Tables.Parse(source, v)
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
