package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode
)

var Select = _select{}

type _select struct{}

func (s _select) Parse(source internal.TokenSource, v validator.Validator) (pnode.SelectStmt, error) {
	source.Checkpoint()

	// Select
	selectNode := pnode.NewSelectStmt()
	err := v.CurrentIsAnd(token.Select).
		NextIs(token.SpaceBreak)
	if err != nil {
		source.Rollback()
		return selectNode, err
	}

	// Fields
	source.Next()
	err = s.parseFields(&selectNode, source, v)
	if err != nil {
		source.Rollback()
		return selectNode, err
	}

	// From
	err = v.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Identifier)
	if err != nil {
		source.Rollback()
		return selectNode, err
	}

	// tables
	err = s.parseTables(&selectNode, source, v)
	if err != nil {
		source.Rollback()
		return selectNode, err
	}

	source.CommitAndInitNode(&selectNode.node)
	return selectNode, nil
}

func (s _select) parseFields(
	selectNode *pnode.SelectStmt,
	source internal.TokenSource,
	v validator.Validator,
) error {
	source.Checkpoint()

	parsingToken := source.Current()
	for {
		if parsingToken.Code() != token.Identifier {
			source.Rollback()
			return sqlerr.NewSyntaxError(token.ToString(token.Identifier), parsingToken.Value(), source)
		}

		field, err := TargetEntry.Parse(source, v)
		if err != nil {
			source.Rollback()
			return err
		}
		selectNode.AddField(field)

		err = v.SkipTokens().
			Type(token.SpaceBreak).
			TypeMinMax(token.Comma, 1, 1).
			SkipFromNext()

		if err != nil {
			source.Commit()
			return nil
		}
		source.Next()
	}
}

func (s _select) parseTables(
	selectNode *pnode.Select,
	source internal.TokenSource,
	v validator.Validator,
) error {
	source.Checkpoint()

	for {
		table, err := Table.Parse(source, v)
		if err != nil {
			return err
		}
		selectNode.Tables = append(selectNode.Tables, table)

		err = v.SkipTokens().
			Type(token.SpaceBreak).
			TypeExactly(token.Comma, 1).
			SkipFromNext()
		if err != nil {
			source.Commit()
			return nil
		}
	}
}
