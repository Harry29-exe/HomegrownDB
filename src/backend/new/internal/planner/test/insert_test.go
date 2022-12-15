package planner

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/backend/new/internal/planner"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

func TestInsertPlanner_SimpleInsert(t *testing.T) {
	//given
	query := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	store, usersTab := TestTableStore.StoreWithUsersTable(t)
	expectedPlan := InsertTests.expectedPlan_SimpleInsert(usersTab, t)

	//when
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	plan, err := planner.Plan(queryNode)
	assert.ErrIsNil(err, t)

	//then
	NodeAssert.Eq(expectedPlan, plan, t)
}

var InsertTests = insertTests{}

type insertTests struct{}

func (insertTests) expectedPlan_SimpleInsert(usersTab table.Definition, t *testing.T) node.PlanedStmt {
	nodeIdCounter := appsync.NewSimpleCounter[node.PlanNodeId](0)
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	plannedStmt := node.NewPlanedStmt(node.CommandTypeInsert)
	usersRTE := node.NewRelationRTE(rteIdCounter.Next(), usersTab)
	plannedStmt.Tables = []node.RangeTableEntry{usersRTE}

	modifyTablePlan := node.NewModifyTable(nodeIdCounter.Next(), node.ModifyTableInsert, nil)
	modifyTablePlan.TargetList = []node.TargetEntry{
		node.NewTargetEntry(nil, tt_user.C0IdOrder, "id"),
		node.NewTargetEntry(nil, tt_user.C2NameOrder, "name"),
	}
	modifyTablePlan.ResultRelations = []node.RteID{usersRTE.Id}

	valueScan := node.NewValueScan(nodeIdCounter.Next(), nil)
	valueScan.Values = [][]node.Expr{
		{node.NewConst(ctype.TypeInt8, 1)},
		{node.NewConst(ctype.TypeStr, "bob")},
	}

	modifyTablePlan.Left = valueScan
	plannedStmt.PlanTree = modifyTablePlan

	return plannedStmt
}
