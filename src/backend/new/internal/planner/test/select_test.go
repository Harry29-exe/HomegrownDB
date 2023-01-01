package planner_test

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/backend/new/internal/planner"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/relation/table"
	"testing"
)

func TestSelectPlanner_SimpleSelect(t *testing.T) {
	//given
	query := "SELECT u.name FROM users u"
	store, usersTab := TestTableStore.StoreWithUsersTable(t)
	expectedPlannedStmt := expectedPlan_SimpleSelect(usersTab, t)

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

func expectedPlan_SimpleSelect(usersTab table.Definition, t *testing.T) node.PlanedStmt {
	planedStmt := node.NewPlanedStmt(node.CommandTypeSelect)
	rootState := planner.NewRootState(planedStmt)

	rte := node.NewRelationRTE(0, usersTab)
	rte.Alias = node.NewAlias("u")
	planRoot := node.NewSeqScan(rootState.NextPlanNodeId(), nil)
	planRoot.RteId = rte.Id
	planRoot.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(rte.Id, tt_user.C2NameOrder, tt_user.C2NameType), 0, ""),
	}

	planedStmt.Tables = []node.RangeTableEntry{rte}
	planedStmt.PlanTree = planRoot
	return planedStmt
}
