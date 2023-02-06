package seganalyser_test

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"testing"
)

func TestSelect_u_name_FROM_users(t *testing.T) {
	// given
	query := "SELECT u.name FROM users u"

	store, usersTable := TestTableStore.StoreWithUsersTable(t)
	expectedQuery := expectedTree_u_name_FROM_users(usersTable)

	//when
	stmt, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(stmt, store)
	assert.ErrIsNil(err, t)

	// then
	NodeAssert.Eq(expectedQuery, queryNode, t)
}

func expectedTree_u_name_FROM_users(users tabdef.Definition) node2.Query {
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
