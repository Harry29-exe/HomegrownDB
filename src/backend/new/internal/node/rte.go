package node

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

// -------------------------
//      RangeTableEntry
// -------------------------

type rteKind = uint8

const (
	RteRelation rteKind = iota
	RteSubQuery
	RteJoin
	RteFunc
	RteTableFunc
	RteValues
	RteCte
	RteNamedTupleStore
	RteResult
)

func NewRelationRTE(rteID RteID, ref table.RDefinition) RangeTableEntry {
	return &rangeTableEntry{
		node:    node{tag: TagRTE},
		Kind:    RteRelation,
		Id:      rteID,
		TableId: ref.RelationID(),
		Ref:     ref,
	}
}

func NewSelectRTE(id RteID, subquery Query) RangeTableEntry {
	return &rangeTableEntry{
		node:     node{tag: TagRTE},
		Kind:     RteSubQuery,
		Id:       id,
		Subquery: subquery,
	}
}

type RangeTableEntry = *rangeTableEntry

var _ Node = &rangeTableEntry{}

// RangeTableEntry is db table that is used in query or plan
type rangeTableEntry struct {
	node
	Kind rteKind

	// Kind = RteRelation
	Id       RteID
	LockMode table.TableLockMode
	TableId  table.Id
	Ref      table.RDefinition

	// Kind = RteSubQuery
	Subquery *query

	//Kind = RteJoin
	JoinType     JoinType
	ResultCols   []Var          // list of columns in result tuples
	LeftColumns  []column.Order // columns
	RightColumns []column.Order

	Alias Alias
}

func (r RangeTableEntry) CreateRef() RangeTableRef {
	return &rangeTableRef{
		node: node{tag: TagRteRef},
		Rte:  r.Id,
	}
}

func (r RangeTableEntry) DEqual() bool {
	//TODO implement me
	panic("implement me")
}

// RteID is id of RangeTableEntry unique for given query/plan
type RteID = uint16

// -------------------------
//      RangeTableRef
// -------------------------

type RangeTableRef = *rangeTableRef

func NewRangeTableRef(rte RangeTableEntry) RangeTableRef {
	return &rangeTableRef{
		node: node{tag: TagRteRef},
		Rte:  rte.Id,
	}
}

var _ Node = &rangeTableRef{}

// rangeTableRef is ref to RTE in query/plan
type rangeTableRef struct {
	node
	Rte RteID
}

func (r RangeTableRef) DEqual() bool {
	//TODO implement me
	panic("implement me")
}

// -------------------------
//      TargetEntry
// -------------------------

func NewTargetEntry(execExpr Expr, attribNo uint16, colName string) TargetEntry {
	return &targetEntry{
		expr:       newExpr(TagTargetEntry),
		ExprToExec: execExpr,
		AttribNo:   attribNo,
		ColName:    colName,
	}
}

type TargetEntry = *targetEntry

var _ Node = &targetEntry{}

type targetEntry struct {
	expr              // Expr to treat TargetEntry as Expr node
	ExprToExec Expr   // ExprToExec expression to evaluate to
	AttribNo   uint16 // AttribNo number of entry
	ColName    string // ColName nullable column alias

	Temp bool // Temp if true then entry should be eliminated before tuple is emitted
}

func (t targetEntry) DEqual() bool {
	//TODO implement me
	panic("implement me")
}
