package node

type PlanStmt = *planStmt

type planStmt struct {
	node
	Command CommandType
	Tables  []RangeTableEntry
}
