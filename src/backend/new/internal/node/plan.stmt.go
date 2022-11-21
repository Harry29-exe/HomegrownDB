package node

type PlanStmt = *planStmt

type planStmt struct {
	Node
	Command CommandType
	Tables  []RangeTableEntry
}
