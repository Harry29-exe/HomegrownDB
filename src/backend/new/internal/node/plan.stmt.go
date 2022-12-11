package node

type PlanStmt = *planStmt

var _ Node = &planStmt{}

type planStmt struct {
	node
	Command CommandType
	Tables  []RangeTableEntry
}

func (p planStmt) DEqual(node Node) bool {
	//TODO implement me
	panic("implement me")
}
