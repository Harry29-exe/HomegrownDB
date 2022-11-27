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
		selectNode, err = StdSelect.parseFullSelect(src, v)
	case token.Values:
		selectNode, err = ValueStreamSelect.parseValueStream(src, v)
	}

	if err != nil {
		src.Rollback()
	} else {
		src.Commit()
	}
	return selectNode, err
}

// -------------------------
//      ValueStreamSelect
// -------------------------

var ValueStreamSelect = valueStreamSelect{Select}

type valueStreamSelect struct{ _select }

func (s valueStreamSelect) parseValueStream(src tkSource, v tkValidator) (pnode.SelectStmt, error) {
	src.Checkpoint()
	selectNode := pnode.NewSelectStmt()
	if err := v.NextIs(token.SpaceBreak); err != nil {
		src.Rollback()
		return selectNode, err
	}

	rows, err := s.parseValueStreamRow(src, v)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	selectNode.Values = rows

	src.CommitAndInitNode(selectNode)
	return selectNode, nil
}

func (s valueStreamSelect) parseValueStreamRow(src tkSource, v tkValidator) ([][]pnode.Node, error) {
	if err := v.NextIs(token.OpeningParenthesis); err != nil {
		return nil, err
	} else if src.Next().Code() == token.SpaceBreak {
		src.Next()
	}
	vals := make([][]pnode.Node, 0, 10)

	val, err := Values.Parse(src, v)
	if err != nil {
		return nil, err
	}
	vals = append(vals, val)
	err = v.SkipCurrentSBAnd().CurrentIs(token.ClosingParenthesis)
	if err != nil {
		return nil, err
	}

	for s.hasNextRow(src, v) {
		val, err = Values.Parse(src, v)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)

		err = v.SkipCurrentSBAnd().CurrentIs(token.ClosingParenthesis)
		if err != nil {
			return nil, err
		}
	}

	return vals, nil
}

func (s valueStreamSelect) hasNextRow(src tkSource, v tkValidator) bool {
	err := v.SkipCurrentSBAnd().
		CurrentIsAnd(token.Comma).
		SkipCurrentSBAnd().
		CurrentIsAnd(token.OpeningParenthesis).
		SkipCurrentSB()

	return err == nil
}

// -------------------------
//      StdSelect
// -------------------------

var StdSelect = stdSelect{Select}

type stdSelect struct {
	_select
}

func (s stdSelect) parseFullSelect(src internal.TokenSource, v tkValidator) (pnode.SelectStmt, error) {
	src.Checkpoint()
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

//todo change for ResultTargets
func (s stdSelect) parseFields(
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

		field, err := ResultTarget.Parse(source, v, ResultTargetSelect)
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

func (s stdSelect) parseTables(
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
