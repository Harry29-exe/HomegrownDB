package node

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/inputtype"
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

func NewVar(id RteID, colOrder column.Order, t hgtype.TypeTag) Var {
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
	Type     hgtype.TypeTag
}

func (v Var) dEqual(node Node) bool {
	raw := node.(Var)
	return v.RteID == raw.RteID &&
		v.ColOrder == raw.ColOrder &&
		v.Type == raw.Type
}

func (v Var) DPrint(nesting int) string {
	return fmt.Sprintf("%s{RteId: %d, ColOrder: %d, TypeTag: %d}",
		v.dTag(nesting), v.RteID, v.ColOrder, v.Type,
	)
}

// -------------------------
//      Const
// -------------------------

var _ Expr = &_const{}

func NewConst(cType hgtype.Wrapper, val []byte) Const {
	return &_const{
		expr: newExpr(TagConst),
		Type: cType,
		Val:  val,
	}
}

func NewConstInt8(val int64, args hgtype.Args) Const {
	t := hgtype.NewInt8(args)
	serializedVal := inputtype.ConvInt8(val)
	return &_const{
		expr: newExpr(TagConst),
		Type: t,
		Val:  serializedVal,
	}
}

func NewConstStr(val string, args hgtype.Args) (Const, error) {
	t := hgtype.NewStr(args)
	serializedVal, err := inputtype.ConvStr(val)
	if err != nil {
		return nil, err
	}
	return &_const{
		expr: newExpr(TagConst),
		Type: t,
		Val:  serializedVal,
	}, nil
}

type Const = *_const

type _const struct {
	expr
	Type hgtype.Wrapper
	Val  []byte // normalized value
}

func (c Const) dEqual(node Node) bool {
	raw := node.(Const)

	if len(c.Val) != len(raw.Val) {
		return false
	}
	for i := 0; i < len(c.Val); i++ {
		if c.Val[i] != raw.Val[i] {
			return false
		}
	}
	return c.Type.TypeEqual(raw.Type)
}

func (c Const) DPrint(nesting int) string {
	return fmt.Sprintf("@%s{TypeTag: %s, Val: %+v}",
		c.dTag(nesting), c.Type.Tag.ToStr(), c.Val)
}

// -------------------------
//      Func
// -------------------------
