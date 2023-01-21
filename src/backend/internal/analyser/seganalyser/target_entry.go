package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	"HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
	sqlerr2 "HomegrownDB/backend/internal/sqlerr"
)

// -------------------------
//      TargetEntry
// -------------------------

var TargetEntry = targetEntry{}

type targetEntry struct{}

func (te targetEntry) AnalyseForSelect(
	resTarget pnode2.ResultTarget,
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
	resTarget pnode2.ResultTarget,
	currentCtx anlsr.QueryCtx,
) (node.TargetEntry, error) {
	val := resTarget.Val
	if val.Tag() != pnode2.TagColumnRef {
		return nil, sqlerr2.NewIllegalPNodeErr(val, "in insert statement only column reference nodes are allowed as target entry")
	}

	colRef := val.(pnode2.ColumnRef)
	query := currentCtx.Query

	// rte must be of kind OwnerID
	rte := query.GetRTE(query.ResultRel)
	colDef, ok := rte.Ref.ColumnByName(colRef.Name)
	if !ok {
		return nil, sqlerr2.AnlsrErr.NewColumnNotExist(query, colRef.Name, "")
	}

	return node.NewTargetEntry(nil, node.AttribNo(colDef.Order()), colDef.Name()), nil
}
