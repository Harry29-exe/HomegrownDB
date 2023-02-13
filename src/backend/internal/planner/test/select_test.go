package planner_test

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planner"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/tests/assert"
	"HomegrownDB/lib/tests/tutils/testtable/tt_user"
	"testing"
)

func TestSelectPlanner_SimpleSelect(t *testing.T) {
	//given
	query := "SELECT u.name FROM users u"
	testDB := hgtest.CreateAndLoadDBWith(nil, t).WithUsersTable().Build()
	store := testDB.DB.AccessModule().RelationManager()
	users := testDB.TableByName(tt_user.TableName)
	expectedPlannedStmt := expectedPlan_SimpleSelect(users, t)

	//when
	rawStmt, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(rawStmt, store)
	assert.ErrIsNil(err, t)

	planedStmt, err := planner.Plan(queryNode)
	assert.ErrIsNil(err, t)

	// then
	NodeAssert.Eq(expectedPlannedStmt, planedStmt, t)
}

func expectedPlan_SimpleSelect(usersTab tabdef.Definition, t *testing.T) node2.PlanedStmt {
	planedStmt := node2.NewPlanedStmt(node2.CommandTypeSelect)
	rootState := planner.NewRootState(planedStmt)

	rte := node2.NewRelationRTE(0, usersTab)
	rte.Alias = node2.NewAlias("u")
	planRoot := node2.NewSeqScan(rootState.NextPlanNodeId(), nil)
	planRoot.RteId = rte.Id
	planRoot.TargetList = []node2.TargetEntry{
		node2.NewTargetEntry(node2.NewVar(rte.Id, tt_user.C2NameOrder, tt_user.C2NameType), 0, ""),
	}

	planedStmt.Tables = []node2.RangeTableEntry{rte}
	planedStmt.PlanTree = planRoot
	return planedStmt
}
