package segparser

import (
	"HomegrownDB/backend/internal/parser/internal"
	"HomegrownDB/backend/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/internal/parser/internal/validator"
	"HomegrownDB/backend/internal/parser/pnode"
)

var Select = _select{}

type _select struct{}

func (s _select) Parse(source internal.TokenSource) (*pnode.Select, error) {
	source.Checkpoint()
	v := validator.NewValidator(source)

	// Select
	err := v.CurrentIsAnd(token.Select).
		NextIs(token.SpaceBreak)
	if err != nil {
		return nil, err
	}
	selectNode := pnode.NewSelect()

	// Fields
	source.Next()
	err = s.parseFields(&selectNode, source, v)
	if err != nil {
		return nil, err
	}

	// From
	err = v.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Identifier)
	if err != nil {
		return nil, err
	}

	// Tables
	err = s.parseTables(&selectNode, source, v)
	if err != nil {
		return nil, err
	}

	source.Commit()
	return &selectNode, nil
}

func (s _select) parseFields(
	selectNode *pnode.Select,
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

		field, err := Field.Parse(source, v)
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
