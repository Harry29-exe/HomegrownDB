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
	rootState := planner.NewRootState(plan)

	usersRTE, valuesRTE := query.RTables[0], query.RTables[1]
	rootState.AppendRTE(usersRTE, valuesRTE)

	modifyTablePlan := node.NewModifyTable(rootState.NextPlanNodeId(), node.ModifyTableInsert, nil)
	modifyTablePlan.TargetList = query.TargetList
	modifyTablePlan.ResultRelations = []node.RteID{usersRTE.Id}

	valueScan := node.NewValueScan(rootState.NextPlanNodeId(), valuesRTE, nil)
	valueScan.RteId = valuesRTE.Id
	modifyTablePlan.Left = valueScan

	plan.PlanTree = modifyTablePlan
	return plan
}
