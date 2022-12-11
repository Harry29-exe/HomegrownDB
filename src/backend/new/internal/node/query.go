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

func NewQuery(commandType CommandType, srcStmt pnode.Node) Query {
	return &query{
		node:    node{tag: TagQuery},
		Command: commandType,
		SrcStmt: srcStmt,
	}
}

var _ Node = &query{}

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

func (q Query) DEqual(node Node) bool {
	if res, ok := nodeEqual(q, node); ok {
		return res
	}
	raw := node.(Query)
	return q.Command == raw.Command &&
		cmpNodeArray(q.TargetList, raw.TargetList) &&
		q.ResultRel == raw.ResultRel &&
		cmpNodeArray(q.RTables, raw.RTables) &&
		q.FromExpr.DEqual(raw.FromExpr)
}
