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
	Type() hgtype.Tag
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

func NewVar(id RteID, colOrder column.Order, typeData hgtype.TypeData) Var {
	return &_var{
		expr:     newExpr(TagVar),
		RteID:    id,
		ColOrder: colOrder,
		TypeData: typeData,
	}
}

var _ Expr = &_var{}

type Var = *_var

type _var struct {
	expr
	RteID    RteID
	ColOrder column.Order
	TypeData hgtype.TypeData
}

func (v Var) Type() hgtype.Tag {
	return v.TypeData.Tag
}

func (v Var) dEqual(node Node) bool {
	raw := node.(Var)
	return v.RteID == raw.RteID &&
		v.ColOrder == raw.ColOrder &&
		v.TypeData == raw.TypeData
}

func (v Var) DPrint(nesting int) string {
	return fmt.Sprintf("%s{RteId: %d, ColOrder: %d, TypeData: %+v}",
		v.dTag(nesting), v.RteID, v.ColOrder, v.TypeData,
	)
}

// -------------------------
//      Const
// -------------------------

var _ Expr = &_const{}

func NewConst(cType hgtype.Tag, val []byte) Const {
	return &_const{
		expr:    newExpr(TagConst),
		TypeTag: cType,
		Val:     val,
	}
}

func NewConstInt8(val int64) Const {
	serializedVal := inputtype.ConvInt8(val)
	return &_const{
		expr:    newExpr(TagConst),
		TypeTag: hgtype.TypeInt8,
		Val:     serializedVal,
	}
}

func NewConstStr(val string) (Const, error) {
	serializedVal, err := inputtype.ConvStr(val)
	if err != nil {
		return nil, err
	}
	return &_const{
		expr:    newExpr(TagConst),
		TypeTag: hgtype.TypeStr,
		Val:     serializedVal,
	}, nil
}

type Const = *_const

type _const struct {
	expr
	TypeTag hgtype.Tag
	Val     []byte // normalized value
}

func (c Const) Type() hgtype.Tag {
	return c.TypeTag
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
	return c.TypeTag == c.TypeTag
}

func (c Const) DPrint(nesting int) string {
	return fmt.Sprintf("@%s{Tag: %s, Val: %+v}",
		c.dTag(nesting), c.TypeTag.ToStr(), c.Val)
}

// -------------------------
//      Func
// -------------------------
