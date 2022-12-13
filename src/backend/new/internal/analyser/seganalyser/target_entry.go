package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

// -------------------------
//      TargetEntries
// -------------------------

var TargetEntries = targetEntries{}

type targetEntries struct{}

func (te targetEntries) Analyse(
	resEntries []pnode.ResultTarget,
	query node.Query,
	ctx anlsr.Ctx,
) ([]node.TargetEntry, error) {

	entries := make([]node.TargetEntry, len(resEntries))

	for i, resTarget := range resEntries {
		entry, err := TargetEntry.Analyse(resTarget, query, ctx)
		if err != nil {
			return nil, err
		}
		entries[i] = entry
	}

	return entries, nil
}

// -------------------------
//      TargetEntry
// -------------------------

var TargetEntry = targetEntry{}

type targetEntry struct{}

func (te targetEntry) Analyse(
	resultTarget pnode.ResultTarget,
	query node.Query,
	ctx anlsr.Ctx,
) (node.TargetEntry, error) {
	valExpr, err := ExprDelegator.DelegateAnalyse(resultTarget.Val, query, ctx)
	if err != nil {
		return nil, err
	}

	attribNo := uint16(len(query.TargetList))
	entry := node.NewTargetEntry(valExpr, attribNo, resultTarget.Name)
	return entry, err
}
