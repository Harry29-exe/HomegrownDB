package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/pnode"
	"errors"
)

type resultTargetMode = uint8

const (
	TargetEntrySelect resultTargetMode = iota
	TargetEntryInset
	TargetEntryUpdate
)

var ResultTargets = resultTargets{}

type resultTargets struct{}

func (t resultTargets) Parse(src tkSource, v tkValidator) ([]pnode.ResultTarget, error) {
	//todo implement me
	panic("Not implemented")
}

var ResultTarget = resultTargetParser{}

type resultTargetParser struct{}

// Parse todo add support for field without table alias
func (t resultTargetParser) Parse(source tkSource, validator tkValidator, mode resultTargetMode) (pnode.ResultTarget, error) {
	source.Checkpoint()

	switch mode {
	case TargetEntrySelect:
		return t.parseSelect(source, validator)
	case TargetEntryInset:
	case TargetEntryUpdate:
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
		_ = src.Next()
		src.CommitAndInitNode(colRef)

		resultTarget := pnode.NewResultTarget("", colRef)
		src.CommitAndInitNode(resultTarget)

		return resultTarget, nil
	}

	//todo add function support

	err = v.CurrentIs(token.Identifier)
	if err == nil {
		src.Checkpoint()
		colName := src.Current().Value()
		colRef := pnode.NewColumnRef(colName, "")
		_ = src.Next()
		src.CommitAndInitNode(colRef)

		resultTarget := pnode.NewResultTarget("", colRef)
		src.CommitAndInitNode(resultTarget)

		return resultTarget, nil
	}

	return nil, errors.New("could not parse field") //todo better err
}

func (t resultTargetParser) parseAlias(resultTarget *pnode.ResultTarget, src internal.TokenSource, v tkValidator) error {
	//err := v.CurrentSequence(token.SpaceBreak, token.As, )
	return errors.New("column aliases are not supported yet")
}
