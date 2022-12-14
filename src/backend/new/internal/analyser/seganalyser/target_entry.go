package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/backend/new/internal/sqlerr"
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
	switch query.Command {
	case node.CommandTypeSelect:
		return te.analyseForSelect(resultTarget, query, ctx)
	case node.CommandTypeInsert:
		return te.analyseForInsert(resultTarget, query, ctx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (te targetEntry) analyseForSelect(
	resTarget pnode.ResultTarget,
	query node.Query,
	ctx anlsr.Ctx,
) (node.TargetEntry, error) {
	valExpr, err := ExprDelegator.DelegateAnalyse(resTarget.Val, query, ctx)
	if err != nil {
		return nil, err
	}

	attribNo := uint16(len(query.TargetList))
	entry := node.NewTargetEntry(valExpr, attribNo, resTarget.Name)
	return entry, err
}

func (te targetEntry) analyseForInsert(
	resTarget pnode.ResultTarget,
	query node.Query,
	ctx anlsr.Ctx,
) (node.TargetEntry, error) {
	val := resTarget.Val
	if val.Tag() != pnode.TagColumnRef {
		return nil, sqlerr.NewIllegalPNodeErr(val, "in insert statement only column reference nodes are allowed as target entry")
	}

	colRef := val.(pnode.ColumnRef)

	// rte must be of kind Relation
	rte := QueryHelper.findRteWithId(query.ResultRel, query)
	colDef, ok := rte.Ref.ColumnByName(colRef.Name)
	if !ok {
		return nil, sqlerr.AnlsrErr.NewColumnNotExist(query, colRef.Name, "")
	}

	return node.NewTargetEntry(nil, colDef.Order(), colDef.Name()), nil
}
