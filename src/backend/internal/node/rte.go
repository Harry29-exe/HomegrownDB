package node

import (
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
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

func NewRelationRTE(rteID RteID, ref tabdef.TableRDefinition) RangeTableEntry {
	return &rangeTableEntry{
		node:    node{tag: TagRTE},
		Kind:    RteRelation,
		Id:      rteID,
		TableId: ref.OID(),
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
	if len(values) == 0 {
		panic("values argument is empty")
	}
	return &rangeTableEntry{
		node:       node{tag: TagRTE},
		Kind:       RteValues,
		Id:         id,
		ValuesList: values,
		ColAlias:   createGenericAliases(len(values[0])),
	}
}

type RangeTableEntry = *rangeTableEntry

var _ Node = &rangeTableEntry{}

// RangeTableEntry is db tabdef that is used in query or plan
type rangeTableEntry struct {
	node
	Kind rteKind
	Id   RteID

	// Kind = RteRelation
	LockMode relation.LockMode
	TableId  tabdef.Id
	Ref      tabdef.TableRDefinition

	// Kind = RteSubQuery
	Subquery *query

	//Kind = RteJoin
	JoinType     JoinType
	ResultCols   []Var          // list of columns in result tuples
	LeftColumns  []tabdef.Order // columns
	RightColumns []tabdef.Order

	//Kind = RteValues
	ValuesList [][]Expr // list of expression node lists

	//Kind = RteValues, RteCte, RteNamedTupleStore, RteTableFunc
	ColTypes []hgtype.ColumnType

	// general
	Alias    Alias
	ColAlias []Alias
}

func (r RangeTableEntry) CreateRef() RangeTableRef {
	return &rangeTableRef{
		node: node{tag: TagRteRef},
		Rte:  r.Id,
	}
}

func (r RangeTableEntry) CreateVarTargetEntry(col tabdef.Order, attribNo AttribNo, colName string) TargetEntry {
	return NewTargetEntry(NewVar(r.Id, col, r.ColTypes[col]), attribNo, colName)
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
Name: %s,
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
		r.Ref.OID() == r2.Ref.OID()
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

type AttribNo = uint16

func NewTargetEntry(execExpr Expr, attribNo AttribNo, colName string) TargetEntry {
	return &targetEntry{
		expr:       newExpr(TagTargetEntry),
		ExprToExec: execExpr,
		AttribNo:   attribNo,
		ColName:    colName,
	}
}

type TargetEntry = *targetEntry

var _ Expr = &targetEntry{}

type targetEntry struct {
	expr                // Expr to treat TargetEntry as Expr node
	ExprToExec Expr     // ExprToExec expression to evaluate to
	AttribNo   AttribNo // AttribNo number of entry
	ColName    string   // ColName nullable column alias

	Temp bool // Temp if true then entry should be eliminated before tuple is emitted
}

func (t TargetEntry) TypeTag() rawtype.Tag {
	return t.ExprToExec.TypeTag()
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

// -------------------------
//      internal
// -------------------------

func createGenericAliases(colCount int) []Alias {
	aliases := make([]Alias, colCount)
	for i := 0; i < colCount; i++ {
		aliases[i] = NewAlias(fmt.Sprintf("C%d", i))
	}
	return aliases
}
