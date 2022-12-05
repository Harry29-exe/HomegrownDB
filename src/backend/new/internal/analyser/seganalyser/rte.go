package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

// -------------------------
//      RteResult
// -------------------------

func NewSingleRteResult(rte node.RangeTableEntry) RteResult {
	return RteResult{
		RteList: []node.RangeTableEntry{rte},
		rteRoot: rte.CreateRef(),
	}
}

type RteResult struct {
	RteList []node.RangeTableEntry
	rteRoot node.Node
}

// -------------------------
//      RTERangeVar
// -------------------------

var RTERangeVar = rteRangeVar{}

type rteRangeVar struct{}

func (r rteRangeVar) Analyse(rangeVar pnode.RangeVar, ctx query.Ctx) (RteResult, error) {
	def, err := ctx.GetTable(rangeVar.RelName)
	if err != nil {
		return RteResult{}, err
	}

	rte := node.NewRelationRTE(ctx.RteIdCounter.IncrAndGet(), def)
	return NewSingleRteResult(rte), nil
}

// -------------------------
//      RTESelect
// -------------------------

var RTESelect = rteSelect{}

type rteSelect struct{}

func (r rteSelect) Analyse(stmt pnode.SelectStmt, ctx query.Ctx) (RteResult, error) {

}
