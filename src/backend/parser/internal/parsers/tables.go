package parsers

import (
	"HomegrownDB/backend/parser/internal/source"
	"HomegrownDB/backend/parser/internal/validator"
	"HomegrownDB/backend/tokenizer/token"
)

var Tables = tablesParser{}

type tablesParser struct {
	validator.Validator
}

// Parse table declarations which are usually found after FROM keyword
// todo add tests
func (t tablesParser) Parse(source source.TokenSource) (*TablesNode, error) {
	t.Init(source)
	source.Checkpoint()

	tables := TablesNode{Tables: make([]*TableNode, 0, 3)}
	for {
		table, err := Table.Parse(source)
		if err != nil {
			return nil, err
		}
		tables.Tables = append(tables.Tables, table)

		err = t.SkipBreaks().
			Type(token.SpaceBreak).
			TypeExactly(token.Comma, 1).
			SkipFromNext()
		if err != nil {
			source.Commit()
			return &tables, nil
		}
	}
}

type TablesNode struct {
	Tables []*TableNode
}
