package node

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
)

// -------------------------
//      RangeTableEntry
// -------------------------

type rteKind uint8

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

func (k rteKind) ToString() string {
	return [...]string{
		"RteRelation",
		"RteSubQuery",
		"RteJoin",
		"RteFunc",
		"RteTableFunc",
		"RteValues",
		"RteCte",
		"RteNamedTupleStore",
		"RteResult",
	}[k]
}

func NewRelationRTE(rteID RteID, ref table.RDefinition) RangeTableEntry {
	return &rangeTableEntry{
		node:    node{tag: TagRTE},
		Kind:    RteRelation,
		Id:      rteID,
		TableId: ref.RelationID(),
		Ref:     ref,
	}
}

func NewSubqueryRTE(id RteID, subquery Query) RangeTableEntry {
	return &rangeTableEntry{
		node:     node{tag: TagRTE},
		Kind:     RteSubQuery,
		Id:       id,
		Subquery: subquery,
	}
}

func NewValuesRTE(id RteID, values [][]Expr) RangeTableEntry {
	return &rangeTableEntry{
		node:       node{tag: TagRTE},
		Kind:       RteValues,
		Id:         id,
		ValuesList: values,
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

	//Kind = RteValues
	ValuesList [][]Expr // list of expression node lists

	//Kind = RteValues, RteCte, RteNamedTupleStore, RteTableFunc
	ColTypes []hgtype.Type

	// general
	Alias Alias
}

func (r RangeTableEntry) CreateRef() RangeTableRef {
	return &rangeTableRef{
		node: node{tag: TagRteRef},
		Rte:  r.Id,
	}
}

func (r RangeTableEntry) dEqual(node Node) bool {
	raw := node.(RangeTableEntry)

	if r.Kind != raw.Kind {
		return false
	} else if !r.dGenericFieldEqual(raw) {
		return false
	}

	switch r.Kind {
	case RteRelation:
		return r.dRelationEqual(raw)
	case RteSubQuery:
		return r.dSubqueryEqual(raw)
	case RteValues:
		return r.dValuesEqual(raw)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func (r rangeTableEntry) DPrint(nesting int) string {
	n1 := nesting + 1
	insertStr := fmt.Sprintf(
		`{
Kind: %s,
Id: %d,
LockMode: %d,
TableId: %d,
Ref: %+v,
Subquery: %s,
JoinType: %s,
ResultCols: %s,
LeftColumns: %+v,
RightColumns: %+v,
Alias: %s,
`,
		r.Kind.ToString(),
		r.Id,
		r.LockMode,
		r.TableId,
		r.Ref,
		r.Subquery.DPrint(n1),
		r.JoinType.ToString(),
		dPrintArr(n1, r.ResultCols),
		r.LeftColumns,
		r.RightColumns,
		r.Alias.DPrint(n1),
	)
	return r.dFormat(insertStr, nesting)
}

func (r RangeTableEntry) dRelationEqual(r2 RangeTableEntry) bool {
	return r.LockMode == r2.LockMode &&
		r.TableId == r2.TableId &&
		r.Ref.TableId() == r2.Ref.TableId()
}

func (r RangeTableEntry) dSubqueryEqual(r2 RangeTableEntry) bool {
	return DEqual(r.Subquery, r2.Subquery)
}

func (r RangeTableEntry) dGenericFieldEqual(r2 RangeTableEntry) bool {
	return r.Id == r2.Id &&
		DEqual(r.Alias, r2.Alias)
}

func (r RangeTableEntry) dValuesEqual(r2 RangeTableEntry) bool {
	return cmpNodeArray2D(r.ValuesList, r2.ValuesList)
}

// RteID is id of RangeTableEntry unique for given Query/plan
type RteID uint16

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

func (r RangeTableRef) dEqual(node Node) bool {
	raw := node.(RangeTableRef)
	return r.Rte == raw.Rte
}

func (r rangeTableRef) DPrint(nesting int) string {
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

func (t TargetEntry) dEqual(node Node) bool {
	raw := node.(TargetEntry)
	return t.AttribNo == raw.AttribNo &&
		t.ColName == raw.ColName &&
		t.Temp == raw.Temp &&
		DEqual(t.ExprToExec, raw.ExprToExec)
}

func (t targetEntry) DPrint(nesting int) string {
	return fmt.Sprintf(
		`
%s{
ExprToExec: %s,
AttribNo: 	%d,
ColName:	%s,
Temp: 		%t
}`,
		t.dTag(nesting),
		t.ExprToExec.DPrint(nesting+1),
		t.AttribNo,
		t.ColName,
		t.Temp)
}
