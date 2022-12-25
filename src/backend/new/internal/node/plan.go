package node

// -------------------------
//      Plan
// -------------------------

type Plan interface {
	Node
	PlanId() PlanNodeId
}

// -------------------------
//      plan
// -------------------------

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
	Left       Plan          // Left (inner) plan, most nodes uses this plan as it only input
	Right      Plan          // Right (outer) plan, used almost exclusively by joins
	InitNodes  []Plan        // InitNodes are plans that needs to be executed separately from this plan, but this plan is dependent on them (e.g. sub-queries)
}

func dPlanEq(p1, p2 *plan) bool {
	return p1.PlanNodeId == p2.PlanNodeId &&
		cmpNodeArray(p1.TargetList, p2.TargetList) &&
		DEqual(p1.Quality, p2.Quality) &&
		DEqual(p1.Left, p2.Left) &&
		DEqual(p1.Right, p2.Right) &&
		cmpNodeArray(p1.InitNodes, p2.InitNodes)
}

func (p *plan) PlanId() PlanNodeId {
	return p.PlanNodeId
}

// -------------------------
//      ModifyTable
// -------------------------

type ModifyTable = *modifyTable

func NewModifyTable(
	planNodeId PlanNodeId,
	operation ModifyTableOp,
	query Query,
) ModifyTable {
	return &modifyTable{
		plan:      newPlan(TagModifyTable, planNodeId, query),
		Operation: operation,
	}
}

type modifyTable struct {
	plan
	Operation       ModifyTableOp
	ResultRelations []RteID
}

func (m ModifyTable) dEqual(node Node) bool {
	raw := node.(ModifyTable)
	return dPlanEq(&m.plan, &raw.plan) &&
		m.Operation == raw.Operation &&
		m.rteIdArrEq(raw)
}

func (m ModifyTable) rteIdArrEq(m2 ModifyTable) bool {
	if len(m.ResultRelations) != len(m2.ResultRelations) {
		return false
	}
	for i := 0; i < len(m.ResultRelations); i++ {
		if m.ResultRelations[i] != m2.ResultRelations[i] {
			return false
		}
	}
	return true
}

func (m ModifyTable) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

type ModifyTableOp uint8

const (
	ModifyTableInsert ModifyTableOp = iota
)

// -------------------------
//      Scan
// -------------------------

var _ Plan = &scan{}

type scan struct {
	plan
	RteId RteID
}

func (s Scan) dEqual(node Node) bool {
	raw := node.(Scan)
	return s.RteId == raw.RteId &&
		dPlanEq(&s.plan, &raw.plan)
}

func (s scan) DPrint(nesting int) string {
	//TODO implement me
	panic("implement me")
}

type SeqScan = *seqScan

// -------------------------
//      SeqScan
// -------------------------

type Scan = *scan

func NewSeqScan(planNodeId PlanNodeId, query Query) SeqScan {
	return &seqScan{
		scan: scan{
			plan: newPlan(TagSeqScan, planNodeId, query),
		},
	}
}

var _ Plan = &seqScan{}

type seqScan struct {
	scan
}

func (s SeqScan) dEqual(node Node) bool {
	raw := node.(SeqScan)
	return DEqual(&s.scan, &raw.scan)
}

// -------------------------
//      ValuesScan
// -------------------------

func NewValueScan(planNodeId PlanNodeId, valuesRTE RangeTableEntry, query Query) ValueScan {
	return &valueScan{
		scan: scan{
			plan:  newPlan(TagValueScan, planNodeId, query),
			RteId: valuesRTE.Id,
		},
		Values: valuesRTE.ValuesList,
	}
}

type ValueScan = *valueScan

type valueScan struct {
	scan
	Values [][]Expr
}

func (v ValueScan) dEqual(node Node) bool {
	raw := node.(ValueScan)
	return DEqual(&v.scan, &raw.scan) &&
		cmpNodeArray2D(v.Values, raw.Values)
}

func (v ValueScan) DPrint(nesting int) string {
	//todo implement me
	panic("Not implemented")
}
