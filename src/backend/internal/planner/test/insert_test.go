package planner

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/planner"
	. "HomegrownDB/backend/internal/testinfr"
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

func (i simpleInsert) expectedPlan(query node2.Query, t *testing.T) node2.PlanedStmt {
	plan := node2.NewPlanedStmt(node2.CommandTypeInsert)
	rootState := planner.NewRootState(plan)

	usersRTE, valuesRTE := query.RTables[0], query.RTables[1]
	rootState.AppendRTE(usersRTE, valuesRTE)

	modifyTablePlan := node2.NewModifyTable(rootState.NextPlanNodeId(), node2.ModifyTableInsert, nil)
	modifyTablePlan.TargetList = nil
	modifyTablePlan.ResultRelations = []node2.RteID{usersRTE.Id}

	valueScan := node2.NewValueScan(rootState.NextPlanNodeId(), valuesRTE, nil)
	valueScan.TargetList = query.TargetList
	valueScan.RteId = valuesRTE.Id
	modifyTablePlan.Left = valueScan

	plan.PlanTree = modifyTablePlan
	return plan
}
