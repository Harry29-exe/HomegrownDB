package testinfr

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
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