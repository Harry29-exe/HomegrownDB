package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
)

var ExprDelegator = exprDelegator{}

type exprDelegator struct{}

func (ex exprDelegator) DelegateAnalyse(node pnode.Node, query node.Query, ctx anlsr.Ctx) {

}

var ExprAnalyser = exprAnalyser{}

type exprAnalyser struct {}

func (ex exprAnalyser) AnalyseColRef(node pnode.ColumnRef, query node.Query, ctx anlsr.Ctx) (node.Var, error) {
	if alias := node.TableAlias; alias != "" {

	}
}

func (ex exprAnalyser) findRteWithAlias(alias string, query node.Query) node.RangeTableEntry {
	for _, table := range query.RTables {
		if table.Subquery {

		}
	}
}


