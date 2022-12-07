package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	"HomegrownDB/common/datastructs/collection/list"
)

var FromDelegator = fromDelegator{}

type fromDelegator struct{}

func (f fromDelegator) Analyse(query node.Query, fromRoot []pnode.Node, ctx anlsr.Ctx) error {
	switch query.Command {
	case node.CommandTypeSelect:
		return SelectFromExpr.Analyse(query, fromRoot, ctx)
	case node.CommandTypeInsert:
		return InsertFromExpr.Analyse(query, fromRoot, ctx)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

// -------------------------
//      SelectFromExpr
// -------------------------

var SelectFromExpr = selectFromExpr{}

type selectFromExpr struct{}

func (f selectFromExpr) Analyse(query node.Query, fromRoot []pnode.Node, ctx anlsr.Ctx) error {
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
	query.RTables = rteList.CurrentSlice()
	return nil
}

func (f selectFromExpr) analyseSingle(root pnode.Node, rteList list.List[node.RangeTableEntry], ctx anlsr.Ctx) (node.Node, error) {
	var result RteResult
	var err error

	switch root.Tag() {
	case pnode.TagRangeVar:
		result, err = RTERangeVar.Analyse(root.(pnode.RangeVar), ctx)
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

// -------------------------
//      InsertFromExpr
// -------------------------

var InsertFromExpr = insertFromExpr{}

type insertFromExpr struct{}

func (f insertFromExpr) Analyse(query node.Query, fromRoot []pnode.Node, ctx anlsr.Ctx) error {
	//todo implement me
	panic("Not implemented")
}
