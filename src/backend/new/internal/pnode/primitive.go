package pnode

import "fmt"

// -------------------------
//      ResultTarget
// -------------------------

// ResultTarget is target for statements like select, insert, update.
// In each means a little different thing,
// for select it is column in row that's being returned,
// for insert and update it is column on which operations are performed
type ResultTarget = *resultTarget

var _ Node = &resultTarget{}

func NewResultTarget(name string, val Node) ResultTarget {
	return &resultTarget{
		node: node{
			tag: TagResultTarget,
		},
		Name: name,
		Val:  val,
	}
}

func NewAStarResultTarget() ResultTarget {
	return &resultTarget{
		node: node{tag: TagResultTarget},
		Name: "",
		Val:  NewAStar(),
	}
}

var _ Node = &resultTarget{}

type resultTarget struct {
	node
	/* Name is in:
	- select - nullable selecting field alias
	- insert - inserting column name
	- update - name of destination column
	*/
	Name string
	/* Val is in:
	- select - expression to get field value
	- insert - not used
	- update - expression to get field value
	*/
	Val Node
}

func (r ResultTarget) Equal(node Node) bool {
	if nodesEqNil(r, node) {
		return true
	} else if !basicNodeEqual(r, node) {
		return false
	}
	raw := node.(ResultTarget)
	return r.Name == raw.Name && r.Val.Equal(raw.Val)
}

// -------------------------
//      ColumnRef
// -------------------------

// ColumnRef reference to column
type ColumnRef = *columnRef

func NewColumnRef(name, tableAlias string) ColumnRef {
	return &columnRef{
		node:       node{tag: TagColumnRef},
		Name:       name,
		TableAlias: tableAlias,
	}
}

var _ Node = &columnRef{}

type columnRef struct {
	node
	Name       string // Name of referenced column
	TableAlias string // TableAlias column's table alias or ""
}

func (c ColumnRef) Equal(node Node) bool {
	if nodesEqNil(c, node) {
		return true
	} else if !basicNodeEqual(c, node) {
		return false
	}

	if c.tag != node.Tag() {
		return false
	}
	cRef := node.(ColumnRef)
	return c.Name == cRef.Name && c.TableAlias == cRef.TableAlias
}

// -------------------------
//      AStar
// -------------------------

type AStar = *aStar

func NewAStar() AStar {
	return &aStar{node{
		tag: TagAStar,
	}}
}

type aStar struct {
	node
}

func (s AStar) Equal(node Node) bool {
	if nodesEqNil(s, node) {
		return true
	}
	return basicNodeEqual(s, node)
}

// -------------------------
//      AConst
// -------------------------

type AConst = *aConst

type aConstType = uint8

const (
	AConstFloat aConstType = iota
	AConstInt
	AConstStr
)

func NewAConstFloat(val float64) AConst {
	return &aConst{
		node:  node{tag: TagAConst},
		Type:  AConstFloat,
		Float: val,
	}
}

func NewAConstInt(val int64) AConst {
	return &aConst{
		node: node{tag: TagAConst},
		Type: AConstInt,
		Int:  val,
	}
}

func NewAConstStr(val string) AConst {
	return &aConst{
		node: node{tag: TagAConst},
		Type: AConstStr,
		Str:  val,
	}
}

type aConst struct {
	node
	Type  aConstType
	Str   string
	Float float64
	Int   int64
}

func (c AConst) Equal(node Node) bool {
	if nodesEqNil(c, node) {
		return true
	} else if !basicNodeEqual(c, node) {
		return false
	}
	raw := node.(AConst)
	if c.Type != raw.Type {
		return false
	}
	switch c.Type {
	case AConstStr:
		return c.Str == raw.Str
	case AConstInt:
		return c.Int == raw.Int
	case AConstFloat:
		return c.Float == raw.Float
	default:
		panic(fmt.Sprintf("not supported AConstType: %d", c.Type))
	}
}

// -------------------------
//      RangeVar
// -------------------------

type RangeVar = *rangeVar

func NewRangeVar(relName string, alias string) RangeVar {
	return &rangeVar{
		node:    node{tag: TagRangeVar},
		RelName: relName,
		Alias:   alias,
	}
}

// RangeVar range variable used in from clauses (basically table)
type rangeVar struct {
	node
	RelName string
	Alias   string
}

func (r RangeVar) Equal(node Node) bool {
	if nodesEqNil(r, node) {
		return true
	} else if !basicNodeEqual(r, node) {
		return false
	}
	raw := node.(RangeVar)
	return r.RelName == raw.RelName && r.Alias == raw.Alias
}

// -------------------------
//      AExpr
// -------------------------

type AExpr = *aExpr

type aExpr struct {
	node
	Kind         AExprKind
	OperatorName string
	Left         Node
	Right        Node
}

func (a AExpr) Equal(node Node) bool {
	if nodesEqNil(a, node) {
		return true
	} else if !basicNodeEqual(a, node) {
		return false
	}
	raw := node.(AExpr)
	return a.Kind == raw.Kind &&
		a.OperatorName == raw.OperatorName &&
		a.Left.Equal(raw.Left) &&
		a.Right.Equal(raw.Right)
}

type AExprKind = uint8

const (
	AExprKindOP AExprKind = iota
)
