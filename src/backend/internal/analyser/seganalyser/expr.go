package seganalyser

import (
	"HomegrownDB/backend/internal/analyser/anlsr"
	node2 "HomegrownDB/backend/internal/node"
	pnode2 "HomegrownDB/backend/internal/pnode"
	. "HomegrownDB/backend/internal/sqlerr"
	"errors"
)

var ExprDelegator = exprDelegator{}

type exprDelegator struct{}

func (ex exprDelegator) DelegateAnalyse(
	pnodeExpr pnode2.Node,
	currentCtx anlsr.QueryCtx,
) (node2.Expr, error) {
	switch pnodeExpr.Tag() {
	case pnode2.TagColumnRef:
		return ExprAnalyser.AnalyseColRef(pnodeExpr.(pnode2.ColumnRef), currentCtx)
	case pnode2.TagAConst:
		return ExprAnalyser.AnalyseConst(pnodeExpr.(pnode2.AConst), currentCtx)
	default:
		return nil, errors.New("") //todo better error
	}
}

var ExprAnalyser = exprAnalyser{}

type exprAnalyser struct{}

func (ex exprAnalyser) AnalyseColRef(pnode pnode2.ColumnRef, currentCtx anlsr.QueryCtx) (node2.Var, error) {
	var rTable node2.RangeTableEntry
	var query = currentCtx.Query

	if alias := pnode.TableAlias; alias != "" {
		rTable = QueryHelper.findRteByAlias(alias, query)
	} else {
		rTable = QueryHelper.findRteWithField(pnode.Name, query)
	}

	if rTable == nil {
		return nil, AnlsrErr.NewColumnNotExist(query, pnode.Name, pnode.TableAlias)
	}

	col, ok := rTable.Ref.ColumnByName(pnode.Name)
	if !ok {
		return nil, AnlsrErr.NewColumnNotExist(query, pnode.Name, pnode.TableAlias)
	}

	return node2.NewVar(rTable.Id, col.Order(), col.CType()), nil
}

func (ex exprAnalyser) AnalyseConst(aConst pnode2.AConst, currentCtx anlsr.QueryCtx) (node2.Const, error) {
	switch aConst.Type {
	case pnode2.AConstInt:
		return node2.NewConstInt8(aConst.Int), nil
	case pnode2.AConstStr:
		return node2.NewConstStr(aConst.Str)
	case pnode2.AConstFloat:
		//todo implement me
		panic("Not implemented")
		//return node.NewConst(hgtype.TypeFloat8, aConst.Float), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
