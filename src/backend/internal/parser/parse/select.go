package parse

import (
	"HomegrownDB/backend/internal/parser/tokenizer"
	token2 "HomegrownDB/backend/internal/parser/tokenizer/token"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
	"fmt"
)

var Select = _select{}

type _select struct{}

func (s _select) Parse(src tokenizer.TokenSource, v tkValidator) (pnode2.SelectStmt, error) {
	src.Checkpoint()
	var selectNode pnode2.SelectStmt
	var err error

	switch currentTk := src.Current(); currentTk.Code() {
	case token2.Select:
		selectNode, err = StdSelect.parseFullSelect(src, v)
	case token2.Values:
		selectNode, err = ValueStreamSelect.parseValueStream(src, v)
	default:
		expected := fmt.Sprintf("%s or %s", token2.ToString(token2.Select), token2.ToString(token2.Values))
		err = sqlerr.NewSyntaxError(expected, token2.ToString(currentTk.Code()), src)
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

func (s valueStreamSelect) parseValueStream(src tkSource, v tkValidator) (pnode2.SelectStmt, error) {
	src.Checkpoint()
	selectNode := pnode2.NewSelectStmt()
	if err := v.NextIs(token2.SpaceBreak); err != nil {
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

func (s valueStreamSelect) parseValueStreamRow(src tkSource, v tkValidator) ([][]pnode2.Node, error) {
	if err := v.NextIs(token2.OpeningParenthesis); err != nil {
		return nil, err
	} else if src.Next().Code() == token2.SpaceBreak {
		src.Next()
	}
	vals := make([][]pnode2.Node, 0, 10)

	val, err := Values.Parse(src, v)
	if err != nil {
		return nil, err
	}
	vals = append(vals, val)
	err = v.SkipCurrentSBAnd().CurrentIs(token2.ClosingParenthesis)
	if err != nil {
		return nil, err
	}

	for s.hasNextRow(src, v) {
		val, err = Values.Parse(src, v)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)

		err = v.SkipCurrentSBAnd().CurrentIs(token2.ClosingParenthesis)
		if err != nil {
			return nil, err
		}
	}

	return vals, nil
}

func (s valueStreamSelect) hasNextRow(src tkSource, v tkValidator) bool {
	err := v.SkipCurrentSBAnd().
		CurrentIsAnd(token2.Comma).
		SkipCurrentSBAnd().
		CurrentIsAnd(token2.OpeningParenthesis).
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

func (s stdSelect) parseFullSelect(src tokenizer.TokenSource, v tkValidator) (pnode2.SelectStmt, error) {
	src.Checkpoint()
	// Select
	selectNode := pnode2.NewSelectStmt()
	err := v.NextIs(token2.SpaceBreak)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// Fields
	src.Next()
	selectNode.Targets, err = ResultTargets.Parse(src, v, ResultTargetSelect)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// From
	err = v.CurrentSequence(token2.SpaceBreak, token2.From, token2.SpaceBreak, token2.Identifier)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	// tables
	rangeVar, err := RangeVar.Parse(src, v)
	selectNode.From = []pnode2.Node{rangeVar}
	//err = s.parseTables(selectNode, src, v)
	if err != nil {
		src.Rollback()
		return selectNode, err
	}

	src.CommitAndInitNode(selectNode)
	return selectNode, nil
}
