package analyse

import (
	"HomegrownDB/backend/internal/analyser/anlctx"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/lib/datastructs/collection/list"
)

var FromDelegator = fromDelegator{}

type fromDelegator struct{}
type RteList = *list.List[node.RangeTableEntry]

// todo validate result for different operations e.g. select, insert e.t.c
func (f fromDelegator) Analyse(fromRoot []pnode.Node, currentCtx anlctx.QueryCtx) error {
	query := currentCtx.Query
	fromExpr := node.NewFromExpr(len(fromRoot))

	var err error
	for i, fromNode := range fromRoot {
		fromExpr.FromList[i], err = f.analyseSingle(fromNode, currentCtx)
		if err != nil {
			return err
		}
	}

	query.FromExpr = fromExpr
	return nil
}

func (f fromDelegator) analyseSingle(root pnode.Node, currentCtx anlctx.QueryCtx) (node.Node, error) {
	var result RteResult
	var err error

	switch root.Tag() {
	case pnode.TagRangeVar:
		result, err = RteRangeVar.Analyse(root.(pnode.RangeVar), currentCtx)
	case pnode.TagSelectStmt:
		result, err = RteSubquery.Analyse(root.(pnode.SelectStmt), currentCtx)

	default:
		//todo implement me
		panic("Not implemented")
	}

	if err != nil {
		return nil, err
	}

	currentCtx.Query.RTables = append(currentCtx.Query.RTables, result.Rte)
	return result.RteRefNode, nil
}
