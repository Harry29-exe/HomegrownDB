package segparser

import (
	"HomegrownDB/backend/internal/parser/tokenizer"
	"HomegrownDB/backend/internal/parser/tokenizer/token"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/backend/internal/sqlerr"
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

	targetValue, err := Value.Parse(src, v)
	if err != nil {
		src.Rollback()
		return nil, err
	}
	resTarget := pnode.NewResultTarget("", targetValue)
	err = t.parseAlias(resTarget, src, v)

	return resTarget, err
}

func (t resultTargetParser) parseInsert(src tkSource, v tkValidator) (pnode.ResultTarget, error) {
	src.Checkpoint()
	current := src.Current()
	if current.Code() != token.Identifier {
		src.Rollback()
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

func (t resultTargetParser) parseAlias(resultTarget pnode.ResultTarget, src tokenizer.TokenSource, v tkValidator) error {
	err := v.CurrentSequence(token.SpaceBreak, token.As, token.SpaceBreak)
	if err != nil {
		return nil
	}

	alias := src.Next()
	if alias.Code() != token.Identifier {
		return sqlerr.NewTokenSyntaxError(token.Identifier, alias.Code(), src)
	}
	resultTarget.Name = alias.Value()
	return nil
}
