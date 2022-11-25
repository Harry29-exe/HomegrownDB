package segparser

import (
	"HomegrownDB/backend/new/internal/parser/internal"
	"HomegrownDB/backend/new/internal/parser/internal/tokenizer/token"
	"HomegrownDB/backend/new/internal/parser/internal/validator"
	"HomegrownDB/backend/new/internal/pnode"
	"errors"
)

var TargetEntry = targetEntry{}

type targetEntryMode = uint8

const (
	TargetEntrySelect targetEntryMode = iota
	TargetEntryInset
	TargetEntryUpdate
)

type targetEntry struct{}

// Parse todo add support for field without table alias
func (f targetEntry) Parse(source internal.TokenSource, validator validator.Validator, mode targetEntryMode) (pnode.ResultTarget, error) {
	source.Checkpoint()

	switch mode {
	case TargetEntrySelect:
		return f.parseSelect(source, validator)
	case TargetEntryInset:
	case TargetEntryUpdate:
		//todo implement me
		panic("Not implemented")
	}

	panic("not supported mode")
}

func (f targetEntry) parseSelect(src internal.TokenSource, v validator.Validator) (pnode.ResultTarget, error) {
	src.Checkpoint()

	err := v.CurrentSequence(token.Identifier, token.Dot, token.Identifier)
	if err == nil {
		src.Checkpoint()
		tableAlias, colName := src.GetPtrRelative(-2), src.GetPtrRelative(0)
		colRef := pnode.NewColumnRef(colName.Value(), tableAlias.Value())
		_ = src.Next()
		src.CommitAndInitNode(&colRef)

		resultTarget := pnode.NewResultTarget("", &colRef)
		src.CommitAndInitNode(&resultTarget)

		return resultTarget, nil
	}

	//todo add function support

	err = v.CurrentIs(token.Identifier)
	if err == nil {
		src.Checkpoint()
		colName := src.Current().Value()
		colRef := pnode.NewColumnRef(colName, "")
		_ = src.Next()
		src.CommitAndInitNode(&colRef)

		resultTarget := pnode.NewResultTarget("", &colRef)
		src.CommitAndInitNode(&resultTarget)

		return resultTarget, nil
	}

	return pnode.ResultTarget{}, errors.New("could not parse field") //todo better err
}

func (f targetEntry) parseAlias(resultTarget *pnode.ResultTarget, src internal.TokenSource, v validator.Validator) error {
	//err := v.CurrentSequence(token.SpaceBreak, token.As, )
	return errors.New("column aliases are not supported yet")
}
