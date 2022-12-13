package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/datastructs/collection/list"
)

var FromDelegator = fromDelegator{}

type fromDelegator struct{}
type RteList = *list.List[node.RangeTableEntry]

// todo validate result for different operations e.g. select, insert e.t.c
func (f fromDelegator) Analyse(fromRoot []pnode.Node, query node.Query, ctx anlsr.Ctx) error {
	fromExpr := node.NewFromExpr(len(fromRoot))
	rteList := list.CopySliceAsList(query.RTables)

	var err error
	for i, fromNode := range fromRoot {
		fromExpr.FromList[i], err = f.analyseSingle(fromNode, rteList, ctx)
		if err != nil {
			return err
		}
	}

	query.FromExpr = fromExpr
	query.RTables = append(query.RTables, rteList.CurrentSlice()...)
	return nil
}

func (f fromDelegator) analyseSingle(root pnode.Node, rteList RteList, ctx anlsr.Ctx) (node.Node, error) {
	var result RteResult
	var err error

	switch root.Tag() {
	case pnode.TagRangeVar:
		result, err = RteRangeVar.Analyse(root.(pnode.RangeVar), ctx)
	case pnode.TagSelectStmt:
		result, err = RteSubquery.Analyse(root.(pnode.SelectStmt), ctx)

	default:
		//todo implement me
		panic("Not implemented")
	}

	if err != nil {
		return nil, err
	}

	rteList.Add(result.Rte)
	return result.RteRefNode, nil
}
