package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
)

var Select = _select{}

type _select struct{}

func (s _select) Parse(src internal.TokenSource, v tkValidator) (pnode.SelectStmt, error) {
	src.Checkpoint()
	var selectNode pnode.SelectStmt
	var err error

	switch currentTk := src.Current(); currentTk.Code() {
	case token.Select:
		selectNode, err = s.parseFullSelect(src, v)
	case token.Values:
		selectNode, err = s.parseValueStream(src, v)
	}

	return selectNode, err
}

func (s _select) parseFullSelect(src internal.TokenSource, v tkValidator) (pnode.SelectStmt, error) {
	// Select
	selectNode := pnode.NewSelectStmt()
	err := v.NextIs(token.SpaceBreak)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// Fields
	src.Next()
	err = s.parseFields(selectNode, src, v)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// From
	err = v.NextSequence(token.SpaceBreak, token.From, token.SpaceBreak, token.Identifier)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// tables
	err = s.parseTables(selectNode, src, v)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	src.CommitAndInitNode(selectNode)
	return selectNode, nil
}

func (s _select) parseValueStream(src tkSource, v tkValidator) (pnode.SelectStmt, error) {
	selectNode := pnode.NewSelectStmt()
	if err := v.NextIs(token.SpaceBreak); err != nil {
		return selectNode, err
	}

	//todo implement me
	panic("Not implemented")
}

//todo change for ResultTargets
func (s _select) parseFields(
	selectNode pnode.SelectStmt,
	source tkSource,
	v tkValidator,
) error {
	source.Checkpoint()

	parsingToken := source.Current()
	for {
		if parsingToken.Code() != token.Identifier {
			source.Rollback()
			return sqlerr.NewSyntaxError(token.ToString(token.Identifier), parsingToken.Value(), source)
		}

		field, err := ResultTarget.Parse(source, v, TargetEntrySelect)
		if err != nil {
			source.Rollback()
			return err
		}
		selectNode.Targets = append(selectNode.Targets, field)

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
	selectNode pnode.SelectStmt,
	source tkSource,
	v tkValidator,
) error {
	source.Checkpoint()

	for {
		table, err := Table.Parse(source, v)
		if err != nil {
			return err
		}
		selectNode.From = append(selectNode.From, table)

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
