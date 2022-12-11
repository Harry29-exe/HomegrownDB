package seganalyser_test

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

func TestSelect_SimpleQuery(t *testing.T) {
	// given
	query := "SELECT u.name FROM users u"
	expectedQuery := node.NewQuery(node.CommandTypeSelect, nil)

	stmt, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	store, err := table.NewTestTableStore(tt_user.Def(t))
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(stmt, store)
	assert.ErrIsNil(err, t)

	_, _ = expectedQuery, queryNode
}
