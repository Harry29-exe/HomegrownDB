package node

import "HomegrownDB/dbsystem/schema/relation"

type PlanNodeId = uint16
type Plan = *plan

var _ Node = &plan{}

type plan struct {
	node

	planNodeId PlanNodeId    // planNodeId unique id of node in given plan
	query      Query         // query source of this
	targetList []TargetEntry // targetList entries that this plan will produce
	quality    Expr          // quality is Expr filter on input data
	left       *plan         // left (inner) plan, most nodes uses this plan as it only input
	right      *plan         // right (outer) plan, used almost exclusively by joins
	initNodes  []*plan       // initNodes are plans that needs to be executed separately from this plan, but this plan is dependent on them (e.g. sub-queries)

}

func (p plan) DEqual(node Node) bool {
	//TODO implement me
	panic("implement me")
}

type ModifyTable struct {
	plan
	Operation    ModifyTableOp
	RootRelation relation.ID
}

type ModifyTableOp = uint8

const (
	ModifyTableInsert ModifyTableOp = iota
)
