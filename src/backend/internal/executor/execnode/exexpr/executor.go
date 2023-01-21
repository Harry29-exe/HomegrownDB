package exexpr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
)

type ExNodeInput struct {
	Plan node2.Plan

	Internal   dpage.Tuple // internal result if node is scan node
	LeftInput  dpage.Tuple // input from left node
	RightInput dpage.Tuple // input from right node
}

func Execute(expr node2.Expr, input ExNodeInput) []byte {
	switch expr.Tag() {
	case node2.TagConst:
		return executeConst(expr.(node2.Const), input)
	case node2.TagVar:
		return executeVar(expr.(node2.Var), input)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func executeConst(expr node2.Const, input ExNodeInput) []byte {
	return expr.Val
}

func executeVar(expr node2.Var, input ExNodeInput) []byte {
	return input.Internal.ColValue(expr.ColOrder)
}
