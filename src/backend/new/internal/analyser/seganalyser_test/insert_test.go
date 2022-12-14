package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/ctype"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

func TestInsertAnalyse_SimplePositive1(t *testing.T) {
	//given
	query := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	store, users := TestTableStore.StoreWithUsersTable(t)
	expectedNode := InsertTests.expectedSimplePositive1(users)

	//when
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	//then
	NodeAssert.Eq(expectedNode, queryNode, t)
}

func (insertTest) expectedSimplePositive1(users table.Definition) node.Query {
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	query := node.NewQuery(node.CommandTypeInsert, nil)
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(nil, tt_user.C0IdOrder, "id"),
		node.NewTargetEntry(nil, tt_user.C2NameOrder, "name"),
	}

	resultRel := node.NewRelationRTE(rteIdCounter.Next(), users)
	query.RTables = []node.RangeTableEntry{resultRel}
	query.ResultRel = resultRel.Id

	subQuery := node.NewQuery(node.CommandTypeSelect, nil)
	valuesRte := node.NewValuesRTE(rteIdCounter.Next(), [][]node.Node{
		{node.NewConst(ctype.TypeInt8, int64(1)), node.NewConst(ctype.TypeStr, "bob")},
	})
	subQuery.RTables = []node.RangeTableEntry{valuesRte}
	subQuery.FromExpr = node.NewFromExpr2(nil, valuesRte.CreateRef())

	subQueryRte := node.NewSubqueryRTE(rteIdCounter.Next(), subQuery)
	query.RTables = append(query.RTables, subQueryRte)
	query.FromExpr = node.NewFromExpr2(nil, subQueryRte.CreateRef())

	return query
}

var InsertTests = insertTest{}

type insertTest struct{}
