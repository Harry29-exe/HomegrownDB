package node

import (
	"HomegrownDB/backend/new/internal/pnode"
	"fmt"
)

// -------------------------
//      CommandType
// -------------------------

type CommandType uint8

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

func (c CommandType) ToString() string {
	return [...]string{
		"CommandTypeSelect",
		"CommandTypeUpdate",
		"CommandTypeDelete",
		"CommandTypeUtils",
	}[c]
}

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
	SrcStmt pnode.Node // stmt that was used to create this Query

	TargetList []TargetEntry

	ResultRel RteID             // Id of result table, for insert, update, delete
	RTables   []RangeTableEntry // Tables used in Query
	FromExpr  FromExpr
}

func (q Query) dEqual(node Node) bool {

	raw := node.(Query)
	return q.Command == raw.Command &&
		cmpNodeArray(q.TargetList, raw.TargetList) &&
		q.ResultRel == raw.ResultRel &&
		cmpNodeArray(q.RTables, raw.RTables) &&
		DEqual(q.FromExpr, raw.FromExpr)
}

func (q query) DPrint(nesting int) string {
	n1 := nesting + 1
	return fmt.Sprintf(
		`%s{
Command: 	%s, 
TargetList: %s,
ResultRel: 	%d,
RTables: 	%s
FromExpr: 	%s`,
		q.dTag(nesting),
		q.Command.ToString(),
		dPrintArr(n1, q.TargetList),
		q.ResultRel,
		dPrintArr(n1, q.RTables),
		q.FromExpr.DPrint(n1),
	)
}

func (q Query) GetRTE(id RteID) RangeTableEntry {
	for i := 0; i < len(q.RTables); i++ {
		if q.RTables[i].Id == id {
			return q.RTables[i]
		}
	}
	return nil
}

func (q Query) AppendRTE(rte RangeTableEntry) {
	q.RTables = append(q.RTables, rte)
}
