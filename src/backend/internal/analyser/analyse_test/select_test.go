package analyse_test

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/tests/assert"
	"HomegrownDB/lib/tests/tutils/testtable/tt_user"
	"testing"
)

func TestSelect_u_name_FROM_users(t *testing.T) {
	// given
	query := "SELECT u.name FROM users u"
	testdb := hgtest.CreateAndLoadDBWith(nil, t).
		WithUsersTable().
		Build()
	store := testdb.DB.AccessModule().RelationManager()
	users := testdb.TableByName(tt_user.TableName)
	expectedQuery := expectedTree_u_name_FROM_users(users)

	//when
	stmt, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(stmt, store)
	assert.ErrIsNil(err, t)

	// then
	NodeAssert.Eq(expectedQuery, queryNode, t)
}

func expectedTree_u_name_FROM_users(users reldef.TableDefinition) node2.Query {
	expectedQuery := node2.NewQuery(node2.CommandTypeSelect, nil)
	rte := node2.NewRelationRTE(0, users)
	rte.Alias = node2.NewAlias("u")
	expectedQuery.RTables = []node2.RangeTableEntry{rte}
	expectedQuery.TargetList = []node2.TargetEntry{
		node2.NewTargetEntry(
			node2.NewVar(rte.Id, tt_user.C2NameOrder, tt_user.C2NameType),
			0,
			"",
		),
	}
	fromExpr := node2.NewFromExpr(1)
	fromExpr.FromList[0] = rte.CreateRef()
	expectedQuery.FromExpr = fromExpr

	return expectedQuery
}

func TestName(t *testing.T) {

}
