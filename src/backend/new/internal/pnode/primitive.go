package node

import "HomegrownDB/backend/new/internal/node"

type ResultTarget struct {
	PNode
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
	val *node.Node
}

// RangeVar range variable used in from clauses (basically table)
type RangeVar struct {
	PNode
	RelName string
	Alias   string
}

type AExpr struct {
	PNode
	Kind         AExprKind
	OperatorName string
	Left         *PNode
	Right        *PNode
}

type AExprKind = uint8

const (
	AExprKindOP AExprKind = iota
)
