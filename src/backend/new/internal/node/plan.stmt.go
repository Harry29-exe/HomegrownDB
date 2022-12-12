package node

import "HomegrownDB/common/datastructs/appsync"

type PlanNodeId uint16

func NewPlanedStmt(command CommandType) PlanedStmt {
	return &planedStmt{
		node:            node{tag: TagPlanedStmt},
		Command:         command,
		PlanNodeCounter: appsync.NewSyncCounter[PlanNodeId](0),

		Tables: make([]RangeTableEntry, 0, 10),
	}
}

type PlanedStmt = *planedStmt

var _ Node = &planedStmt{}

type planedStmt struct {
	node

	Command         CommandType
	PlanNodeCounter PlanNodeCounter

	PlanTree Plan
	Tables   []RangeTableEntry
}

func (p PlanedStmt) NextPlanNodeId() PlanNodeId {
	return p.PlanNodeCounter.GetAndIncr()
}

func (p PlanedStmt) dEqual(node Node) bool {
	raw := node.(PlanedStmt)
	return p.Command == raw.Command &&
		DEqual(p.PlanTree, raw.PlanTree) &&
		cmpNodeArray(p.Tables, raw.Tables)
}

func (p PlanedStmt) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

type PlanNodeCounter = appsync.SyncCounter[PlanNodeId]
