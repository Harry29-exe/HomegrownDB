package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var RTE = rte{}

type rte struct{}

func (r rte) Analyse(rtePNode pnode.Node, ctx query.Ctx) (node.RangeTableEntry, error) {
	switch rtePNode.Tag() {
	case pnode.TagRangeVar:
		return r.analyseRangeVar(rtePNode.(pnode.RangeVar), ctx)
	case pnode.TagSelectStmt:
		return r.analyseSelectStmt(rtePNode.(pnode.SelectStmt), ctx)

	default:
		panic("")
	}
}

func (r rte) analyseRangeVar(rangeVar pnode.RangeVar, ctx query.Ctx) (node.RangeTableEntry, error) {
	def, err := ctx.GetTable(rangeVar.RelName)
	if err != nil {
		return nil, err
	}

	return node.NewRelationRTE(ctx.RteIdCounter.IncrAndGet(), def), nil
}

func (r rte) analyseSelectStmt(selectStmt pnode.SelectStmt, ctx query.Ctx) (node.RangeTableEntry, error) {
	q, err := Query.Analyse(selectStmt, ctx)
	if err != nil {
		return nil, err
	}
	node.NewSelectRTE(ctx.RteIdCounter.IncrAndGet(), q)
}
