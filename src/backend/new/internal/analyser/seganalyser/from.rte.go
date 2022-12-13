package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
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
//      RTERangeVar
// -------------------------

var RTERangeVar = rteRangeVar{}

type rteRangeVar struct{}

func (r rteRangeVar) Analyse(rangeVar pnode.RangeVar, ctx anlsr.Ctx) (RteResult, error) {
	def, err := ctx.GetTable(rangeVar.RelName)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewRelationRTE(ctx.RteIdCounter.IncrAndGet(), def)
	if rangeVar.Alias != "" {
		rte.Alias = node.NewAlias(rangeVar.Alias)
	}

	return NewSingleRteResult(rte), nil
}
