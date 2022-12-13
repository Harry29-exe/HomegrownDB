package planner_test

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/backend/new/internal/planner"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/schema/table"
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

	rte := node.NewRelationRTE(1, usersTab)
	planRoot := node.NewSeqScan(planedStmt.PlanNodeCounter.GetAndIncr(), nil)
	planRoot.RteId = rte.Id
	planRoot.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(rte.Id, tt_user.C2NameOrder, tt_user.C2NameType), 0, ""),
	}

	planedStmt.Tables = []node.RangeTableEntry{rte}
	planedStmt.PlanTree = planRoot
	return planedStmt
}