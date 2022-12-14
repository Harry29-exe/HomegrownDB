package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser/anlsr"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/pnode"
	. "HomegrownDB/backend/new/internal/sqlerr"
	"HomegrownDB/dbsystem/ctype"
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
	case pnode.TagAConst:
		return ExprAnalyser.AnalyseConst(pnodeExpr.(pnode.AConst), query, ctx)
	default:
		return nil, errors.New("") //todo better error
	}
}

var ExprAnalyser = exprAnalyser{}

type exprAnalyser struct{}

func (ex exprAnalyser) AnalyseColRef(pnode pnode.ColumnRef, query node.Query, ctx anlsr.Ctx) (node.Var, error) {
	var rTable node.RangeTableEntry

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

	return node.NewVar(rTable.Id, col.Order(), col.Type()), nil
}

func (ex exprAnalyser) AnalyseConst(aConst pnode.AConst, query node.Query, ctx anlsr.Ctx) (node.Const, error) {
	switch aConst.Type {
	case pnode.AConstInt:
		return node.NewConst(ctype.TypeInt8, aConst.Int), nil
	case pnode.AConstStr:
		return node.NewConst(ctype.TypeStr, aConst.Str), nil
	case pnode.AConstFloat:
		return node.NewConst(ctype.TypeFloat8, aConst.Float), nil
	default:
		//todo implement me
		panic("Not implemented")
	}
}
