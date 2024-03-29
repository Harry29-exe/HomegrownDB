package testinfr

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/planner"
	"HomegrownDB/dbsystem/access/relation"
	"HomegrownDB/lib/datastructs/appsync"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

type RteIdCounter = appsync.SimpleSyncCounter[node2.RteID]

func ParseAndAnalyse(query string, relManager relation.Manager, t *testing.T) node2.Query {
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, relManager)
	assert.ErrIsNil(err, t)

	return queryNode
}

func ParseAnalyseAndPlan(query string, relManager relation.Manager, t *testing.T) node2.PlanedStmt {
	queryTree := ParseAndAnalyse(query, relManager, t)
	planedStmt, err := planner.Plan(queryTree)
	assert.ErrIsNil(err, t)
	return planedStmt
}

func NewConstStr(val string, t *testing.T) node2.Const {
	aConst, err := node2.NewConstStr(val)
	assert.ErrIsNil(err, t)
	return aConst
}
