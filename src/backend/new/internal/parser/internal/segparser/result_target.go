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

	tableToken, err := validator.Current().
		Has(token.Identifier).
		IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return pnode.ResultTarget{}, err
	}

	columnToken, err := validator.
		NextIsAnd(token.Dot).
		Next().IsTextToken().
		DontStartWithDigit().
		AsciiOnly().
		Check()

	if err != nil {
		source.Rollback()
		return pnode.FieldNode{}, err
	}

	fieldNode := pnode.FieldNode{
		TableAlias: tableToken.Value(),
		FieldName:  columnToken.Value(),
		FieldAlias: columnToken.Value(),
	}
	source.CommitAndInitNode(&fieldNode.Node)
	return fieldNode, nil
}

func (f targetEntry) parseSelect(src internal.TokenSource, v validator.Validator) (pnode.ResultTarget, error) {
	src.Checkpoint()

	err := v.CurrentSequence(token.Identifier, token.Dot, token.Identifier)
	if err == nil {
		src.Checkpoint()
		tableAlias, colName := src.GetPtrRelative(-2), src.GetPtrRelative(0)
		colRef := pnode.NewColumnRef(colName.Value(), tableAlias.Value())
		src.CommitAndInitNode(colRef)

		resultTarget := pnode.NewResultTarget("", colRef)
		_ = src.Next()
		src.CommitAndInitNode(resultTarget)

		return resultTarget, nil
	}

	//todo add function support

	err := v.CurrentIs(token.Identifier)
	if err == nil {
		colName := src.Current().Value()

	}
}

func (f targetEntry) parseAlias(resultTarget *pnode.ResultTarget, src internal.TokenSource, v validator.Validator) error {
	//err := v.CurrentSequence(token.SpaceBreak, token.As, )
	return errors.New("column aliases are not supported yet")
}
