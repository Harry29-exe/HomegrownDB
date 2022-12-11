package node

import (
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/column"
	"fmt"
)

type Expr interface {
	Node
	ExprTag() Tag
}

func newExpr(exprTag Tag) expr {
	return expr{
		node: node{tag: exprTag},
	}
}

type expr struct {
	node
}

func (e expr) ExprTag() Tag {
	return TagExpr
}

// -------------------------
//      Var
// -------------------------

func NewVar(id RteID, colOrder column.Order, t ctype.Type) Var {
	return &_var{
		expr:     newExpr(TagVar),
		RteID:    id,
		ColOrder: colOrder,
		Type:     t,
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

func (v Var) dEqual(node Node) bool {
	raw := node.(Var)
	return v.RteID == raw.RteID &&
		v.ColOrder == raw.ColOrder &&
		v.Type == raw.Type
}

func (v Var) DPrint(nesting int) string {
	return fmt.Sprintf("%s{RteId: %d, ColOrder: %d, Type: %d}",
		v.dTag(nesting), v.RteID, v.ColOrder, v.Type,
	)
}

// -------------------------
//      Const
// -------------------------

var _ Expr = &_const{}

type Const = *_const

type _const struct {
	expr
	Type ctype.Type
	Val  any
}

func (c Const) dEqual(node Node) bool {
	raw := node.(Const)
	return c.Type == raw.Type &&
		c.Val == raw.Type
}

func (c Const) DPrint(nesting int) string {
	return fmt.Sprintf("@%s{Type: %d, Val: %+v}",
		c.dTag(nesting), c.Type, c.Val)
}

// -------------------------
//      Func
// -------------------------
