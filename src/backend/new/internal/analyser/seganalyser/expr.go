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
	currentCtx anlsr.QueryCtx,
) (node.Expr, error) {
	switch pnodeExpr.Tag() {
	case pnode.TagColumnRef:
		return ExprAnalyser.AnalyseColRef(pnodeExpr.(pnode.ColumnRef), currentCtx)
	case pnode.TagAConst:
		return ExprAnalyser.AnalyseConst(pnodeExpr.(pnode.AConst), currentCtx)
	default:
		return nil, errors.New("") //todo better error
	}
}

var ExprAnalyser = exprAnalyser{}

type exprAnalyser struct{}

func (ex exprAnalyser) AnalyseColRef(pnode pnode.ColumnRef, currentCtx anlsr.QueryCtx) (node.Var, error) {
	var rTable node.RangeTableEntry
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

	return node.NewVar(rTable.Id, col.Order(), col.CType()), nil
}

func (ex exprAnalyser) AnalyseConst(aConst pnode.AConst, currentCtx anlsr.QueryCtx) (node.Const, error) {
	switch aConst.Type {
	case pnode.AConstInt:
		return node.NewConstInt8(aConst.Int), nil
	case pnode.AConstStr:
		return node.NewConstStr(aConst.Str)
	case pnode.AConstFloat:
		//todo implement me
		panic("Not implemented")
		//return node.NewConst(hgtype.TypeFloat8, aConst.Float), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
