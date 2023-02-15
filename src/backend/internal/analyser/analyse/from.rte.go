package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	node "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
)

// -------------------------
//      RteResult
// -------------------------

func NewSingleRteResult(rte node.RangeTableEntry) RteResult {
	return RteResult{
		Rte:        rte,
		RteRefNode: rte.CreateRef(),
	}
}

type RteResult struct {
	Rte        node.RangeTableEntry
	RteRefNode node.Node // RteRefNode is test like, node.RangeTableRef, Joins etc.
}

// -------------------------
//      RteRangeVar
// -------------------------

var RteRangeVar = rteRangeVar{}

type rteRangeVar struct{}

func (r rteRangeVar) Analyse(rangeVar pnode2.RangeVar, currentCtx anlctx.QueryCtx) (RteResult, error) {
	def, err := currentCtx.GetTable(rangeVar.RelName)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewRelationRTE(currentCtx.RteIdCounter.Next(), def)
	if rangeVar.Alias != "" {
		rte.Alias = node.NewAlias(rangeVar.Alias)
	}

	return NewSingleRteResult(rte), nil
}

// -------------------------
//      RteSubquery
// -------------------------

var RteSubquery = rteSelect{}

type rteSelect struct{}

func (rteSelect) Analyse(stmt pnode2.SelectStmt, currentCtx anlctx.QueryCtx) (RteResult, error) {
	subquery, err := Select.Analyse(stmt, currentCtx)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewSubqueryRTE(currentCtx.RteIdCounter.Next(), subquery)
	return NewSingleRteResult(rte), nil
}
