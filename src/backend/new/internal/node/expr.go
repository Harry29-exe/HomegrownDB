package node

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
)

type Expr = *expr

type expr struct {
	node
}

// -------------------------
//      Var
// -------------------------

func NewVar(id RteID, colOrder column.Order, t ctype.Type) Var {
	return &_var{
		expr:     expr{node{tag: TagVar}},
		RteID:    0,
		ColOrder: 0,
		Type:     0,
	}
}

type Var = *_var

type _var struct {
	expr
	RteID    RteID
	ColOrder column.Order
	Type     ctype.Type
}

// -------------------------
//      Const
// -------------------------

type Const = *_const

type _const struct {
	expr
	Type ctype.Type
	Val any
}

// -------------------------
//      Func
// -------------------------
