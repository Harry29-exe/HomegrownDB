package node

type PlanStmt = *planStmt

var _ Node = &planStmt{}

type planStmt struct {
	node
	Command CommandType
	Tables  []RangeTableEntry
}

func (p planStmt) dEqual(node Node) bool {
	//TODO implement me
	panic("implement me")
}

func (p planStmt) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}
