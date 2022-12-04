package node

import "HomegrownDB/dbsystem/schema/table"

// -------------------------
//      CommandType
// -------------------------

type CommandType = uint8

const (
	// CommandTypeSelect for select statements
	CommandTypeSelect CommandType = iota
	// CommandTypeInsert for insert statements
	CommandTypeInsert
	// CommandTypeUpdate for update statements
	CommandTypeUpdate
	// CommandTypeDelete for delete statements
	CommandTypeDelete
	// CommandTypeUtils for statements that operate
	// rater on db structure than on db data
	CommandTypeUtils
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
		kind:    RteRelation,
		Id:      rteID,
		TableId: ref.RelationID(),
		Ref:     ref,
	}
}

func NewSelectRTE(id RteID, subquery Query) RangeTableEntry {
	return &rangeTableEntry{
		node:  node{tag: TagRTE},
		kind:  RteSubQuery,
		Id:    id,
		Query: subquery,
	}
}

type RangeTableEntry = *rangeTableEntry

var _ Node = &rangeTableEntry{}

// RangeTableEntry is db table that is used in query or plan
type rangeTableEntry struct {
	node
	kind rteKind

	// kind = RteRelation
	Id       RteID
	LockMode table.TableLockMode
	TableId  table.Id
	Ref      table.RDefinition

	// kind = RteSubQuery
	Query Query
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

type TargetEntry struct {
	Expr              // Expr to treat TargetEntry as Expr node
	ExprToExec *Expr  // ExprToExec expression to evaluate to
	AttribNo   uint16 // AttribNo number of entry
	ColName    string // ColName nullable column name

	TableId table.Id // TableId
	Temp    bool     // Temp if true then entry should be eliminated before tuple is emitted
}
