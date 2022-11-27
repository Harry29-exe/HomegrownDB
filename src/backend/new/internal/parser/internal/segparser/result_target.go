package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/sqlerr"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
	"errors"
)

type resultTargetMode = uint8

const (
	ResultTargetSelect resultTargetMode = iota
	ResultTargetInsert
	ResultTargetUpdate
)

// -------------------------
//      ResultTargets
// -------------------------

var ResultTargets = resultTargets{}

type resultTargets struct{}

func (t resultTargets) Parse(src tkSource, v tkValidator, mode resultTargetMode) ([]pnode.ResultTarget, error) {
	src.Checkpoint()
	targets := make([]pnode.ResultTarget, 0, 10)

	if mode == ResultTargetInsert {
		err := v.CurrentIsAnd(token.OpeningParenthesis).
			SkipCurrentSB()
		if err != nil {
			src.Rollback()
			return nil, err
		}
	}

	target, err := ResultTarget.Parse(src, v, mode)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	targets = append(targets, target)
	for t.hasNext(src, v) {
		target, err = ResultTarget.Parse(src, v, mode)
		if err != nil {
			src.Rollback()
			return nil, err
		}
		targets = append(targets, target)
	}

	if mode == ResultTargetInsert {
		err = v.SkipCurrentSBAnd().
			CurrentIs(token.ClosingParenthesis)
		if err != nil {
			src.Rollback()
			return nil, err
		}
	}

	src.Commit()
	return targets, nil
}

func (t resultTargets) hasNext(src tkSource, v tkValidator) bool {
	err := v.SkipTokens().
		TypeExactly(token.Comma, 1).
		TypeMax(token.SpaceBreak, 2).
		SkipFromCurrent()

	return err == nil
}

// -------------------------
//      ResultTarget
// -------------------------

var ResultTarget = resultTargetParser{}

type resultTargetParser struct{}

// Parse todo add support for field without table alias
func (t resultTargetParser) Parse(src tkSource, v tkValidator, mode resultTargetMode) (pnode.ResultTarget, error) {
	switch mode {
	case ResultTargetSelect:
		return t.parseSelect(src, v)
	case ResultTargetInsert:
		return t.parseInsert(src, v)
	case ResultTargetUpdate:
		//todo implement me
		panic("Not implemented")
	}

	panic("not supported mode")
}

func (t resultTargetParser) parseSelect(src tkSource, v tkValidator) (pnode.ResultTarget, error) {
	src.Checkpoint()

	err := v.CurrentSequence(token.Identifier, token.Dot, token.Identifier)
	if err == nil {
		src.Checkpoint()
		tableAlias, colName := src.GetPtrRelative(-2), src.GetPtrRelative(0)
		colRef := pnode.NewColumnRef(colName.Value(), tableAlias.Value())
		src.Next()
		src.CommitAndInitNode(colRef)

		resultTarget := pnode.NewResultTarget("", colRef)
		src.CommitAndInitNode(resultTarget)
		return resultTarget, nil
	}

	//todo add function support

	if src.Current().Code() == token.Identifier {
		src.Checkpoint()
		colName := src.Current().Value()
		colRef := pnode.NewColumnRef(colName, "")
		src.Next()
		src.CommitAndInitNode(colRef)

		resultTarget := pnode.NewResultTarget("", colRef)
		src.CommitAndInitNode(resultTarget)
		return resultTarget, nil
	}

	src.Rollback()
	return nil, errors.New("could not parse field") //todo better err
}

func (t resultTargetParser) parseInsert(src tkSource, v tkValidator) (pnode.ResultTarget, error) {
	src.Checkpoint()
	current := src.Current()
	if current.Code() != token.Identifier {
		return nil, sqlerr.NewTokenSyntaxError(token.Identifier, current.Code(), src)
	}
	src.Checkpoint()
	src.Next()
	cRef := pnode.NewColumnRef(current.Value(), "")
	src.CommitAndInitNode(cRef)

	rt := pnode.NewResultTarget("", cRef)
	src.CommitAndInitNode(rt)

	return rt, nil
}

func (t resultTargetParser) parseAlias(resultTarget *pnode.ResultTarget, src internal.TokenSource, v tkValidator) error {
	//err := v.CurrentSequence(token.SpaceBreak, token.As, )
	return errors.New("column aliases are not supported yet")
}
