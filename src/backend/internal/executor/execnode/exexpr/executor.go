package exexpr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/page"
)

type ExNodeInput struct {
	Plan node.Plan

	Internal   page.Tuple // internal result if node is scan node
	LeftInput  page.Tuple // input from left node
	RightInput page.Tuple // input from right node
}

func Execute(expr node.Expr, input ExNodeInput) []byte {
	switch expr.Tag() {
	case node.TagConst:
		return executeConst(expr.(node.Const), input)
	case node.TagVar:
		return executeVar(expr.(node.Var), input)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func executeConst(expr node.Const, input ExNodeInput) []byte {
	return expr.Val
}

func executeVar(expr node.Var, input ExNodeInput) []byte {
	return input.Internal.ColValue(expr.ColOrder)
}
