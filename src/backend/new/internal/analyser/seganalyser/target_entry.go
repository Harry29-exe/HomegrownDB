package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

// -------------------------
//      TargetEntryDelegator
// -------------------------

var TargetEntryDelegator = targetEntryDelegator{}

type targetEntryDelegator struct{}

func (te targetEntryDelegator) Delegate(
	query node.Query,
	resultTarget []pnode.ResultTarget,
	ctx anlsr.Ctx,
) error {
	switch query.Command {
	case node.CommandTypeSelect:
		return TargetEntrySelect.Analyse(query, resultTarget, ctx)
	case node.CommandTypeInsert:
		//todo implement me
		panic("Not implemented")
	default:
		//todo implement me
		panic("Not implemented")
	}
}

// -------------------------
//      TargetEntrySelect
// -------------------------

var TargetEntrySelect = targetEntrySelect{}

type targetEntrySelect struct{}

func (t targetEntrySelect) Analyse(query node.Query, rts []pnode.ResultTarget, ctx anlsr.Ctx) error {

}

// -------------------------
//      TargetEntry
// -------------------------

var TargetEntry = targetEntry{}

type targetEntry struct{}

func (e targetEntry) Analyse(target pnode.ResultTarget, ctx anlsr.Ctx) (node.TargetEntry, error) {
	val := target.Val

	switch val.Tag() {
	case pnode.TagColumnRef:

	}
}

func (e targetEntry) analyseColumnRef(
	target pnode.ResultTarget,
	colRef pnode.ColumnRef,
	ctx anlsr.Ctx,
) (node.TargetEntry, error) {

}
