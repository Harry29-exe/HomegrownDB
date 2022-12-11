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
	Id   RteID

	// Kind = RteRelation
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

	// general
	Alias Alias
}

func (r RangeTableEntry) CreateRef() RangeTableRef {
	return &rangeTableRef{
		node: node{tag: TagRteRef},
		Rte:  r.Id,
	}
}

func (r RangeTableEntry) DEqual(node Node) bool {
	if res, ok := nodeEqual(r, node); ok {
		return res
	}
	raw := node.(RangeTableEntry)

	if r.Kind != raw.Kind {
		return false
	} else if !r.dRelationEqual(raw) {
		return false
	}

	switch r.Kind {
	case RteRelation:
		return r.dRelationEqual(raw)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (r RangeTableEntry) dRelationEqual(r2 RangeTableEntry) bool {
	return r.LockMode == r2.LockMode &&
		r.TableId == r2.TableId &&
		r.Ref.TableId() == r2.Ref.TableId()
}

func (r RangeTableEntry) dGenericFieldEqual(r2 RangeTableEntry) bool {
	return r.Id == r2.Id &&
		r.Alias.DEqual(r2.Alias)
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

func (r RangeTableRef) DEqual(node Node) bool {
	if res, ok := nodeEqual(r, node); ok {
		return res
	}
	raw := node.(RangeTableRef)
	return r.Rte == raw.Rte
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

func (t TargetEntry) DEqual(node Node) bool {
	if res, ok := nodeEqual(t, node); ok {
		return res
	}
	raw := node.(TargetEntry)
	return t.AttribNo == raw.AttribNo &&
		t.ColName == raw.ColName &&
		t.Temp == raw.Temp &&
		t.ExprToExec.DEqual(raw.ExprToExec)
}
