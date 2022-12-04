package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/query"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

type fromAnalyserResult struct {
	From node.FromExpr
	RTEs []node.RangeTableEntry
}

var From = from{}

type from struct{}

func (f from) Analyse(fromRoot pnode.Node, mode node.CommandType, ctx query.Ctx) (fromAnalyserResult, error) {
	switch fromRoot.Tag() {
	case pnode.TagRangeVar:
		return f.analyseRangeVar(fromRoot.(pnode.RangeVar), ctx)
	}
}

func (f from) analyseRangeVar(rangeVar pnode.RangeVar, ctx query.Ctx) (fromAnalyserResult, error) {
	fromExpr := node.NewFromExpr()

	tableDef, err := ctx.GetTable(rangeVar.RelName)
	if err != nil {
		return fromAnalyserResult{}, err
	}
	rangeTblEntry := node.NewRelationRTE(ctx.RteIdCounter.IncrAndGet(), tableDef)
	fromExpr.FromList = append(fromExpr.FromList, node.NewRangeTableRef(rangeTblEntry))

	return fromAnalyserResult{
		From: fromExpr,
		RTEs: []node.RangeTableEntry{rangeTblEntry},
	}, nil
}
