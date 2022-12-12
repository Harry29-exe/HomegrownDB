package node

import "HomegrownDB/dbsystem/schema/relation"

// -------------------------
//      Plan
// -------------------------

type Plan interface {
	Node
	PlanId() PlanNodeId
}

func newPlan(tag Tag, planNodeId PlanNodeId, query Query) plan {
	return plan{
		node:       node{tag: tag},
		PlanNodeId: planNodeId,
		Query:      query,
	}
}

// plan is abstract node that is composed into
// all nodes that have their executor
type plan struct {
	node

	PlanNodeId PlanNodeId // PlanNodeId unique id of test in given plan
	Query      Query      // Query source of this

	TargetList []TargetEntry // TargetList entries that this plan will produce
	Quality    Expr          // Quality is Expr filter on input data
	Left       Plan         // Left (inner) plan, most nodes uses this plan as it only input
	Right      Plan         // Right (outer) plan, used almost exclusively by joins
	InitNodes  []Plan       // InitNodes are plans that needs to be executed separately from this plan, but this plan is dependent on them (e.g. sub-queries)
}

func (p *plan) dEqual(node Node) bool {
	raw := node.(*plan)
	return p.PlanNodeId == raw.PlanNodeId &&
		cmpNodeArray(p.TargetList, raw.TargetList) &&
		DEqual(p.Quality, raw.Quality) &&
		DEqual(p.Left, raw.Left) &&
		DEqual(p.Right, raw.Right) &&
		cmpNodeArray(p.InitNodes, raw.InitNodes)
}

func (p *plan) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

func (p *plan) PlanId() PlanNodeId {
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

var _ Plan = &scan{}

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

func (s Scan) dEqual(node Node) bool {
	raw := node.(Scan)
	return s.RteId == raw.RteId &&
		DEqual(&s.plan, &raw.plan)
}

// -------------------------
//      SeqScan
// -------------------------

type SeqScan = *seqScan

var _ Plan = &seqScan{}

type seqScan struct {
	scan
}

func (s SeqScan) dEqual(node Node) bool {
	raw := node.(SeqScan)
	return DEqual(&s.scan, &raw.scan)
}
