package parser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Tables = tablesParser{}

type tablesParser struct {
}

// Parse table declarations which are usually found after FROM keyword
// todo add tests
func (t tablesParser) Parse(source internal.TokenSource, validator validator.Validator) (pnode.TablesNode, error) {
	source.Checkpoint()

	tables := pnode.TablesNode{Tables: make([]pnode.TableNode, 0, 3)}
	for {
		table, err := Table.Parse(source, validator)
		if err != nil {
			return tables, err
		}
		tables.Tables = append(tables.Tables, table)

		err = validator.SkipTokens().
			Type(token.SpaceBreak).
			TypeExactly(token.Comma, 1).
			SkipFromNext()
		if err != nil {
			source.Commit()
			return tables, nil
		}
	}
}
