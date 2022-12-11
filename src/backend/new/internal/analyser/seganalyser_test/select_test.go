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

func TestSelect_u_name_FROM_users(t *testing.T) {
	// given
	query := "SELECT u.name FROM users u"

	store, usersTable := StoreWithUsersTable(t)
	expectedQuery := createExpectedTree_u_name_FROM_users(usersTable)

	//when
	stmt, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(stmt, store)
	assert.ErrIsNil(err, t)

	// then
	Assert.Node(expectedQuery, queryNode, t)
}

func createExpectedTree_u_name_FROM_users(users table.Definition) node.Query {
	expectedQuery := node.NewQuery(node.CommandTypeSelect, nil)
	rte := node.NewRelationRTE(1, users)
	expectedQuery.RTables = []node.RangeTableEntry{rte}
	expectedQuery.TargetList = []node.TargetEntry{
		node.NewTargetEntry(
			node.NewVar(rte.Id, tt_user.C2NameOrder, tt_user.C2NameType),
			0,
			"",
		),
	}
	fromExpr := node.NewFromExpr(1)
	fromExpr.FromList[0] = rte.CreateRef()
	expectedQuery.FromExpr = fromExpr

	return expectedQuery
}

func TestName(t *testing.T) {

}
