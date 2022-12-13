package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	. "HomegrownDB/backend/new/internal/sqlerr"
	"errors"
)

var ExprDelegator = exprDelegator{}

type exprDelegator struct{}

func (ex exprDelegator) DelegateAnalyse(
	pnodeExpr pnode.Node,
	query node.Query,
	ctx anlsr.Ctx,
) (node.Expr, error) {
	switch pnodeExpr.Tag() {
	case pnode.TagColumnRef:
		return ExprAnalyser.AnalyseColRef(pnodeExpr.(pnode.ColumnRef), query, ctx)
	default:
		return nil, errors.New("") //todo better error
	}
}

var ExprAnalyser = exprAnalyser{}

type exprAnalyser struct{}

func (ex exprAnalyser) AnalyseColRef(pnode pnode.ColumnRef, query node.Query, ctx anlsr.Ctx) (node.Var, error) {
	var rTable node.RangeTableEntry

	if alias := pnode.TableAlias; alias != "" {
		rTable = ex.findRteByAlias(alias, query)
	} else {
		rTable = ex.findRteWithField(pnode.Name, query)
	}

	if rTable == nil {
		return nil, AnlsrErr.NewColumnNotExist(query, pnode.Name, pnode.TableAlias)
	}

	col, ok := rTable.Ref.ColumnByName(pnode.Name)
	if !ok {
		return nil, AnlsrErr.NewColumnNotExist(query, pnode.Name, pnode.TableAlias)
	}

	return node.NewVar(rTable.Id, col.Order(), col.Type()), nil
}

func (ex exprAnalyser) findRteByAlias(alias string, query node.Query) node.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Alias != nil && rTable.Alias.AliasName == alias {
			return rTable
		}
	}
	return nil
}

func (ex exprAnalyser) findRteWithField(fieldName string, query node.Query) node.RangeTableEntry {
	for _, rTable := range query.RTables {
		if rTable.Kind != node.RteRelation {
			continue
		}

		for _, col := range rTable.Ref.Columns() {
			colName := col.Name()
			if colName == fieldName {
				return rTable
			}
		}
	}
	return nil
}