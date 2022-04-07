package parsers

import (
	"HomegrownDB/sql/querry/parser/def"
	"HomegrownDB/sql/querry/parser/helpers"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Select selectParser = selectParser{}

type selectParser struct{}

func (s selectParser) Parse(source def.TokenSource) (*SelectNode, error) {
	source.Checkpoint()

	// Select
	_, err := helpers.Current(source).
		Has(token.Select).
		Check()
	if err != nil {
		return nil, err
	}
	selectNode := SelectNode{}

	err = helpers.SkipBreaks(source).
		Type(token.SpaceBreak).
		SkipFromNext()
	if err != nil {
		return nil, err
	}

	// Fields
	selectNode.Fields, err = Fields.Parse(source)
	if err != nil {
		return nil, err
	}

	//// Table
	//err = s.Attach()
	return &selectNode, nil
}

type SelectNode struct {
	Fields *FieldsNode
	Tables *TablesNode
}
