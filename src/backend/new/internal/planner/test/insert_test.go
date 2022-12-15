package planner

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/planner"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestInsertPlanner_SimpleInsert(t *testing.T) {
	//given
	inputQuery := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	store, _ := TestTableStore.StoreWithUsersTable(t)
	query := ParseAndAnalyse(inputQuery, store, t)
	expectedPlan := SimpleInsert.expectedPlan(query, t)

	//when
	actualPlan, err := planner.Plan(query)

	//then
	assert.ErrIsNil(err, t)
	NodeAssert.Eq(expectedPlan, actualPlan, t)
}

var SimpleInsert = simpleInsert{}

type simpleInsert struct{}

func (i simpleInsert) expectedPlan(query node.Query, t *testing.T) node.PlanedStmt {
	plan := node.NewPlanedStmt(node.CommandTypeInsert)

	usersRTE, valuesRTE, subqueryRTE := i.unpackRTEs(query.RTables, t)
	plan.AppendRTEs(usersRTE, valuesRTE, subqueryRTE)

	modifyTablePlan := node.NewModifyTable(plan.NextPlanNodeId(), node.ModifyTableInsert, nil)
	modifyTablePlan.TargetList = query.TargetList
	modifyTablePlan.ResultRelations = []node.RteID{usersRTE.Id}

	valueScan := node.NewValueScan(plan.NextPlanNodeId(), valuesRTE.ValuesList, nil)
	modifyTablePlan.Left = valueScan

	plan.PlanTree = modifyTablePlan
	return plan
}

func (i simpleInsert) unpackRTEs(
	insertQueryRTEs []node.RangeTableEntry,
	t *testing.T,
) (
	users node.RangeTableEntry,
	values node.RangeTableEntry,
	subquery node.RangeTableEntry,
) {
	assert.Eq(len(insertQueryRTEs), 2, t)
	for _, rte := range insertQueryRTEs {
		switch rte.Kind {
		case node.RteRelation:
			users = rte
		case node.RteSubQuery:
			subquery = rte
		}
	}
	subqueryRtes := subquery.Subquery.RTables
	assert.Eq(len(subqueryRtes), 1, t)
	values = subqueryRtes[0]
	return
}

func (i simpleInsert) findSubqueryRTE(rtes []node.RangeTableEntry, t *testing.T) node.RangeTableEntry {
	for _, rte := range rtes {
		if rte.Kind != node.RteSubQuery {
			return rte
		}
	}
	t.Error("Expected to find subquery rte")
	return nil
}
