package node

import (
	"HomegrownDB/dbsystem/access/relation/table/column"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"fmt"
)

type Expr interface {
	Node
	ExprTag() Tag
	TypeTag() rawtype.Tag
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

func NewVar(id RteID, colOrder column.Order, typeData hgtype.ColType) Var {
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
	TypeData hgtype.ColType
}

func (v Var) TypeTag() rawtype.Tag {
	return v.TypeData.Tag()
}

func (v Var) dEqual(node Node) bool {
	raw := node.(Var)
	return v.RteID == raw.RteID &&
		v.ColOrder == raw.ColOrder &&
		v.TypeData == raw.TypeData
}

func (v Var) DPrint(nesting int) string {
	return fmt.Sprintf("%s{RteId: %d, ColOrder: %d, ColumnType: %+v}",
		v.dTag(nesting), v.RteID, v.ColOrder, v.TypeData,
	)
}

// -------------------------
//      Const
// -------------------------

var _ Expr = &_const{}

func NewConst(cType rawtype.Tag, val []byte) Const {
	return &_const{
		expr: newExpr(TagConst),
		Type: cType,
		Val:  val,
	}
}

func NewConstInt8(val int64) Const {
	serializedVal := intype.ConvInt8(val)
	return &_const{
		expr: newExpr(TagConst),
		Type: rawtype.TypeInt8,
		Val:  serializedVal,
	}
}

func NewConstStr(val string) (Const, error) {
	serializedVal, err := intype.ConvStr(val)
	if err != nil {
		return nil, err
	}
	return &_const{
		expr: newExpr(TagConst),
		Type: rawtype.TypeStr,
		Val:  serializedVal,
	}, nil
}

type Const = *_const

type _const struct {
	expr
	Type rawtype.Tag
	Val  []byte // normalized value
}

func (c Const) TypeTag() rawtype.Tag {
	return c.Type
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
	return c.Type == c.Type
}

func (c Const) DPrint(nesting int) string {
	return fmt.Sprintf("@%s{ColTag: %s, Val: %+v}",
		c.dTag(nesting), c.Type.ToStr(), c.Val)
}

// -------------------------
//      Func
// -------------------------
