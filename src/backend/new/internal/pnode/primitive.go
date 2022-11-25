package pnode

func NewResultTarget(name string, val Node) ResultTarget {
	return ResultTarget{
		node: node{
			tag: TagResultTarget,
		},
		Name: name,
		val:  val,
	}
}

type ResultTarget struct {
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

func NewColumnRef(name, tableAlias string) ColumnRef {
	return ColumnRef{
		node:       node{tag: TagColumnRef},
		Name:       name,
		TableAlias: tableAlias,
	}
}

type ColumnRef struct {
	node
	Name       string // Name of referenced column
	TableAlias string // TableAlias column's table alias or ""
}

func NewRangeVar(relName string, alias string) RangeVar {
	return RangeVar{
		node:    node{tag: TagRangeVar},
		RelName: relName,
		Alias:   alias,
	}
}

// RangeVar range variable used in from clauses (basically table)
type RangeVar struct {
	node
	RelName string
	Alias   string
}

type AExpr struct {
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
