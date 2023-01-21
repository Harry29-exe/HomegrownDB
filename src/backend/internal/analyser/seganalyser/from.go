package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	node2 "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
	"HomegrownDB/common/datastructs/collection/list"
)

var FromDelegator = fromDelegator{}

type fromDelegator struct{}
type RteList = *list.List[node2.RangeTableEntry]

// todo validate result for different operations e.g. select, insert e.t.c
func (f fromDelegator) Analyse(fromRoot []pnode2.Node, currentCtx anlsr.QueryCtx) error {
	query := currentCtx.Query
	fromExpr := node2.NewFromExpr(len(fromRoot))

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

func (f fromDelegator) analyseSingle(root pnode2.Node, currentCtx anlsr.QueryCtx) (node2.Node, error) {
	var result RteResult
	var err error

	switch root.Tag() {
	case pnode2.TagRangeVar:
		result, err = RteRangeVar.Analyse(root.(pnode2.RangeVar), currentCtx)
	case pnode2.TagSelectStmt:
		result, err = RteSubquery.Analyse(root.(pnode2.SelectStmt), currentCtx)

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
