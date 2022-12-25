package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/backend/new/internal/sqlerr"
)

// -------------------------
//      TargetEntry
// -------------------------

var TargetEntry = targetEntry{}

type targetEntry struct{}

func (te targetEntry) AnalyseForSelect(
	resTarget pnode.ResultTarget,
	currentCtx anlsr.QueryCtx,
) (node.TargetEntry, error) {
	valExpr, err := ExprDelegator.DelegateAnalyse(resTarget.Val, currentCtx)
	if err != nil {
		return nil, err
	}

	attribNo := node.AttribNo(len(currentCtx.Query.TargetList))
	entry := node.NewTargetEntry(valExpr, attribNo, resTarget.Name)
	return entry, err
}

func (te targetEntry) AnalyseForInsert(
	resTarget pnode.ResultTarget,
	currentCtx anlsr.QueryCtx,
) (node.TargetEntry, error) {
	val := resTarget.Val
	if val.Tag() != pnode.TagColumnRef {
		return nil, sqlerr.NewIllegalPNodeErr(val, "in insert statement only column reference nodes are allowed as target entry")
	}

	colRef := val.(pnode.ColumnRef)
	query := currentCtx.Query

	// rte must be of kind Relation
	rte := query.GetRTE(query.ResultRel)
	colDef, ok := rte.Ref.ColumnByName(colRef.Name)
	if !ok {
		return nil, sqlerr.AnlsrErr.NewColumnNotExist(query, colRef.Name, "")
	}

	return node.NewTargetEntry(nil, node.AttribNo(colDef.Order()), colDef.Name()), nil
}
