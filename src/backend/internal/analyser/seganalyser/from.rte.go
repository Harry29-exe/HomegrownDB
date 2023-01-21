package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	node2 "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
)

// -------------------------
//      RteResult
// -------------------------

func NewSingleRteResult(rte node2.RangeTableEntry) RteResult {
	return RteResult{
		Rte:        rte,
		RteRefNode: rte.CreateRef(),
	}
}

type RteResult struct {
	Rte        node2.RangeTableEntry
	RteRefNode node2.Node // RteRefNode is test like, node.RangeTableRef, Joins etc.
}

// -------------------------
//      RteRangeVar
// -------------------------

var RteRangeVar = rteRangeVar{}

type rteRangeVar struct{}

func (r rteRangeVar) Analyse(rangeVar pnode2.RangeVar, currentCtx anlsr.QueryCtx) (RteResult, error) {
	def, err := currentCtx.GetTable(rangeVar.RelName)
	if err != nil {
		return RteResult{}, err
	}

	rte := node2.NewRelationRTE(currentCtx.RteIdCounter.Next(), def)
	if rangeVar.Alias != "" {
		rte.Alias = node2.NewAlias(rangeVar.Alias)
	}

	return NewSingleRteResult(rte), nil
}

// -------------------------
//      RteSubquery
// -------------------------

var RteSubquery = rteSelect{}

type rteSelect struct{}

func (rteSelect) Analyse(stmt pnode2.SelectStmt, currentCtx anlsr.QueryCtx) (RteResult, error) {
	subquery, err := Select.Analyse(stmt, currentCtx)
	if err != nil {
		return RteResult{}, err
	}

	rte := node2.NewSubqueryRTE(currentCtx.RteIdCounter.Next(), subquery)
	return NewSingleRteResult(rte), nil
}
