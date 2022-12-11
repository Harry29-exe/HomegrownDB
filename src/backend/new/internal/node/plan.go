package node

import "HomegrownDB/dbsystem/schema/relation"

// -------------------------
//      Plan
// -------------------------

type Plan interface {
	PlanId() PlanNodeId
}

func newPlan(tag Tag, planNodeId PlanNodeId, query Query) plan {
	return plan{
		node:       node{tag: tag},
		PlanNodeId: planNodeId,
		Query:      query,
	}
}

var _ Node = &plan{}

// plan is abstract node that is composed into
// all nodes that have their executor
type plan struct {
	node

	PlanNodeId PlanNodeId // PlanNodeId unique id of node in given plan
	Query      Query      // Query source of this

	TargetList []TargetEntry // TargetList entries that this plan will produce
	Quality    Expr          // Quality is Expr filter on input data
	Left       *plan         // Left (inner) plan, most nodes uses this plan as it only input
	Right      *plan         // Right (outer) plan, used almost exclusively by joins
	InitNodes  []*plan       // InitNodes are plans that needs to be executed separately from this plan, but this plan is dependent on them (e.g. sub-queries)
}

func (p plan) dEqual(node Node) bool {
	//TODO implement me
	panic("implement me")
}

func (p plan) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

func (p plan) PlanId() PlanNodeId {
	return p.PlanNodeId
}

// -------------------------
//      ModifyTable
// -------------------------

type ModifyTable struct {
	plan
	Operation    ModifyTableOp
	RootRelation relation.ID
}

type ModifyTableOp = uint8

const (
	ModifyTableInsert ModifyTableOp = iota
)

// -------------------------
//      Scan
// -------------------------

type Scan = *scan

type scan struct {
	plan
	RteId RteID
}

func NewSeqScan(planNodeId PlanNodeId, query Query) SeqScan {
	return &seqScan{
		scan: scan{
			plan: newPlan(TagSeqScan, planNodeId, query),
		},
	}
}

// -------------------------
//      SeqScan
// -------------------------

type SeqScan = *seqScan

var _ Node = &seqScan{}

type seqScan struct {
	scan
}
