package pnode

// -------------------------
//      ResultTarget
// -------------------------

// ResultTarget is target for statements like select, insert, update.
// In each means a little different thing,
// for select it is column in row that's being returned,
// for insert and update it is column on which operations are performed
type ResultTarget = *resultTarget

func NewResultTarget(name string, val Node) ResultTarget {
	return &resultTarget{
		node: node{
			tag: TagResultTarget,
		},
		Name: name,
		val:  val,
	}
}

func NewAStarResultTarget() ResultTarget {
	return &resultTarget{
		node: node{tag: TagResultTarget},
		Name: "",
		val:  NewAStar(),
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
	val Node
}

func (r ResultTarget) Equal(node Node) bool {
	if !basicNodeEqual(r, node) {
		return false
	}
	raw := node.(ResultTarget)
	return r.Name == raw.Name && r.val.Equal(raw.val)
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
	if !basicNodeEqual(c, node) {
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
	return basicNodeEqual(s, node)
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

// -------------------------
//      AExpr
// -------------------------

type AExpr = *aExpr

type aExpr struct {
	node
	Kind         AExprKind
	OperatorName string
	Left         *node
	Right        *node
}

type AExprKind = uint8

const (
	AExprKindOP AExprKind = iota
)
