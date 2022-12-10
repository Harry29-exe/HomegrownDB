package node

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
)

type Expr interface {
	Node
	ExprTag() Tag
}

func newExpr(exprTag Tag) expr {
	return expr{
		node:    node{tag: TagExpr},
		exprTag: exprTag,
	}
}

type expr struct {
	node
	exprTag Tag
}

func (e expr) ExprTag() Tag {
	return e.exprTag
}

// -------------------------
//      Var
// -------------------------

func NewVar(id RteID, colOrder column.Order, t ctype.Type) Var {
	return &_var{
		expr:     newExpr(TagVar),
		RteID:    0,
		ColOrder: 0,
		Type:     0,
	}
}

var _ Expr = &_var{}

type Var = *_var

type _var struct {
	expr
	RteID    RteID
	ColOrder column.Order
	Type     ctype.Type
}

func (_ _var) DEqual() bool {
	//TODO implement me
	panic("implement me")
}

// -------------------------
//      Const
// -------------------------

type Const = *_const

type _const struct {
	expr
	Type ctype.Type
	Val  any
}

// -------------------------
//      Func
// -------------------------
