package node

import "HomegrownDB/backend/new/internal/pnode"

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
//      Query
// -------------------------

func NewQuery() Query {
	//todo implement me
	panic("Not implemented")
}

type Query = *query

type query struct {
	node
	Command CommandType
	SrcStmt pnode.Node // stmt that was used to create this query

	TargetList []TargetEntry

	ResultRel RteID             // Id of result table, for insert, update, delete
	RTables   []RangeTableEntry // Tables used in query
	FromExpr  FromExpr
}
