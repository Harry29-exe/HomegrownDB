package parsers

import (
	"HomegrownDB/sql/querry/parser/parsers/helpers"
	"HomegrownDB/sql/querry/parser/parsers/source"
	"HomegrownDB/sql/querry/tokenizer/token"
)

var Tables tablesParser = tablesParser{}

type tablesParser struct {
	helpers.ParserHelper
}

// Parse table declarations which are usually found after FROM keyword
//
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

		source.CommitAndCheckpoint()
		err = t.SkipBreaks().
			Type(token.SpaceBreak).
			TypeExactly(token.Comma, 1).
			SkipFromNext()
		if err != nil {
			source.Rollback()
			return &tables, nil
		}
	}
}

type TablesNode struct {
	Tables []*TableNode
}
