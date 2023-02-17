package planner_test

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/planner"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

var CreateRelation = createRelationUtils{}

type createRelationUtils struct{}

func TestCreateRelationPlanner_SimpleTable(t *testing.T) {
	// given
	query := `CREATE TABLE users (
		username VARCHAR(255),
		surname VARCHAR(255)		
	)`
	testDB := hgtest.CreateAndLoadDBWith(nil, t).Build()
	store := testDB.DB.AccessModule().RelationManager()

	queryTree := ParseAndAnalyse(query, store, t)
	expectedPlan := CreateRelation.simpleTablePlan(queryTree)

	// when
	plan, err := planner.Plan(queryTree)

	// then
	assert.ErrIsNil(err, t)
	NodeAssert.Eq(expectedPlan, plan, t)
}

func (createRelationUtils) simpleTablePlan(queryTree node.Query) node.PlanedStmt {
	plannedStmt := node.NewPlanedStmt(node.CommandTypeUtils)
	state := planner.NewRootState(plannedStmt)
	table := queryTree.UtilsStmt.(node.CreateRelation).FutureTable
	plannedStmt.PlanTree = node.NewCreateRelationPlan(table, state.NextPlanNodeId(), queryTree)

	return plannedStmt
}
