package testinfr

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/backend/new/internal/planner"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

type RteIdCounter = appsync.SimpleSyncCounter[node.RteID]

func ParseAndAnalyse(query string, store table.Store, t *testing.T) node.Query {
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	return queryNode
}

func ParseAnalyseAndPlan(query string, store table.Store, t *testing.T) node.PlanedStmt {
	queryTree := ParseAndAnalyse(query, store, t)
	planedStmt, err := planner.Plan(queryTree)
	assert.ErrIsNil(err, t)
	return planedStmt
}

func NewConstStr(val string, t *testing.T) node.Const {
	aConst, err := node.NewConstStr(val)
	assert.ErrIsNil(err, t)
	return aConst
}
